// Header text

package sample

import (
	"strconv"
	"sync"
	"time"

	"github.com/takuoki/gsheets"
	"github.com/takuoki/sheetdb"
)

const (
	_Bar_sheetName        = "bars"
	_Bar_column_UserID    = 0 // A
	_Bar_column_Datetime  = 1 // B
	_Bar_column_Value     = 2 // C
	_Bar_column_Note      = 3 // D
	_Bar_column_UpdatedAt = 4 // E
	_Bar_column_DeletedAt = 5 // F
)

var (
	_Bar_mutex    = sync.RWMutex{}
	_Bar_cache    = map[int]map[sheetdb.Datetime]*Bar{} // map[userID][datetime]*Bar
	_Bar_rowNoMap = map[int]map[sheetdb.Datetime]int{}  // map[userID][datetime]rowNo
	_Bar_maxRowNo int
)

func init() {
	dbClient.AddModel("Bar", _Bar_load)
}

func _Bar_load(data *gsheets.Sheet) error {

	_Bar_mutex.Lock()
	defer _Bar_mutex.Unlock()

	_Bar_cache = map[int]map[sheetdb.Datetime]*Bar{}
	_Bar_rowNoMap = map[int]map[sheetdb.Datetime]int{}
	_Bar_maxRowNo = 0

	for i, r := range data.Rows() {
		if i == 0 {
			continue
		}
		if r.Value(_Bar_column_DeletedAt) != "" {
			_Bar_maxRowNo++
			continue
		}
		if r.Value(_Bar_column_UserID) == "" {
			break
		}

		userID, err := _Bar_parseUserID(r.Value(_Bar_column_UserID))
		if err != nil {
			return err
		}
		datetime, err := _Bar_parseDatetime(r.Value(_Bar_column_Datetime))
		if err != nil {
			return err
		}
		value, err := _Bar_parseValue(r.Value(_Bar_column_Value))
		if err != nil {
			return err
		}
		note := r.Value(_Bar_column_Note)
		if err := _Bar_validateNote(note); err != nil {
			return err
		}

		bar := Bar{
			UserID:   userID,
			Datetime: datetime,
			Value:    value,
			Note:     note,
		}

		_Bar_maxRowNo++
		if _, ok := _Bar_cache[bar.UserID]; !ok {
			_Bar_cache[bar.UserID] = map[sheetdb.Datetime]*Bar{}
		}
		_Bar_cache[bar.UserID][bar.Datetime] = &bar
		_Bar_rowNoMap[bar.UserID][bar.Datetime] = _Bar_maxRowNo
	}

	return nil
}

func GetBar(userID int, datetime sheetdb.Datetime) (*Bar, error) {
	_Bar_mutex.RLock()
	defer _Bar_mutex.RUnlock()
	if v, ok := _Bar_cache[userID][datetime]; ok {
		return v, nil
	}
	return nil, &sheetdb.NotFoundError{Model: "Bar"}
}

type BarQuery struct {
	filter func(bar *Bar) bool
	sort   func(bars []*Bar)
}

type BarQueryOption func(query *BarQuery) *BarQuery

func BarFilter(filterFunc func(bar *Bar) bool) func(query *BarQuery) *BarQuery {
	return func(query *BarQuery) *BarQuery {
		if query != nil {
			query.filter = filterFunc
		}
		return query
	}
}

func BarSort(sortFunc func(bars []*Bar)) func(query *BarQuery) *BarQuery {
	return func(query *BarQuery) *BarQuery {
		if query != nil {
			query.sort = sortFunc
		}
		return query
	}
}

func GetBars(userID int, opts ...BarQueryOption) ([]*Bar, error) {
	barQuery := &BarQuery{}
	for _, opt := range opts {
		barQuery = opt(barQuery)
	}
	_Bar_mutex.RLock()
	defer _Bar_mutex.RUnlock()
	bars := []*Bar{}
	if barQuery.filter != nil {
		for _, v := range _Bar_cache[userID] {
			if barQuery.filter(v) {
				bars = append(bars, v)
			}
		}
	} else {
		for _, v := range _Bar_cache[userID] {
			bars = append(bars, v)
		}
	}
	if barQuery.sort != nil {
		barQuery.sort(bars)
	}
	return bars, nil
}

func (m *User) AddBar(datetime sheetdb.Datetime, value float32, note string) (*Bar, error) {
	_Bar_mutex.Lock()
	defer _Bar_mutex.Unlock()
	if err := _Bar_validateNote(note); err != nil {
		return nil, err
	}
	bar := &Bar{
		UserID:   m.UserID,
		Datetime: datetime,
		Value:    value,
		Note:     note,
	}
	if err := bar._asyncUpdate(); err != nil {
		return nil, err
	}
	_Bar_maxRowNo++
	if _, ok := _Bar_cache[bar.UserID]; !ok {
		_Bar_cache[bar.UserID] = map[sheetdb.Datetime]*Bar{}
	}
	_Bar_cache[bar.UserID][bar.Datetime] = bar
	_Bar_rowNoMap[bar.UserID][bar.Datetime] = _Bar_maxRowNo
	return bar, nil
}

func AddBar(userID int, datetime sheetdb.Datetime, value float32, note string) (*Bar, error) {
	m, err := GetUser(userID)
	if err != nil {
		return nil, err
	}
	return m.AddBar(datetime, value, note)
}

func (m *User) UpdateBar(datetime sheetdb.Datetime, value float32, note string) (*Bar, error) {
	_Bar_mutex.Lock()
	defer _Bar_mutex.Unlock()
	bar, ok := _Bar_cache[m.UserID][datetime]
	if !ok {
		return nil, &sheetdb.NotFoundError{Model: "Bar"}
	}
	if err := _Bar_validateNote(note); err != nil {
		return nil, err
	}
	barCopy := *bar
	barCopy.Value = value
	barCopy.Note = note
	if err := (&barCopy)._asyncUpdate(); err != nil {
		return nil, err
	}
	bar = &barCopy
	return bar, nil
}

func UpdateBar(userID int, datetime sheetdb.Datetime, value float32, note string) (*Bar, error) {
	m, err := GetUser(userID)
	if err != nil {
		return nil, err
	}
	return m.UpdateBar(datetime, value, note)
}

func (m *User) DeleteBar(datetime sheetdb.Datetime) error {
	_Bar_mutex.Lock()
	defer _Bar_mutex.Unlock()
	bar, ok := _Bar_cache[m.UserID][datetime]
	if !ok {
		return &sheetdb.NotFoundError{Model: "Bar"}
	}
	if err := bar._asyncDelete(); err != nil {
		return err
	}
	delete(_Bar_cache[m.UserID], datetime)
	return nil
}

func DeleteBar(userID int, datetime sheetdb.Datetime) error {
	m, err := GetUser(userID)
	if err != nil {
		return err
	}
	return m.DeleteBar(datetime)
}

func _Bar_validateNote(note string) error {
	return nil
}

func _Bar_parseUserID(userID string) (int, error) {
	v, err := strconv.Atoi(userID)
	if err != nil {
		return 0, &sheetdb.InvalidValueError{FieldName: "UserID", Err: err}
	}
	return v, nil
}

func _Bar_parseDatetime(datetime string) (sheetdb.Datetime, error) {
	v, err := sheetdb.NewDatetime(datetime)
	if err != nil {
		return v, &sheetdb.InvalidValueError{FieldName: "Datetime", Err: err}
	}
	return v, nil
}

func _Bar_parseValue(value string) (float32, error) {
	v, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0, &sheetdb.InvalidValueError{FieldName: "Value", Err: err}
	}
	return float32(v), nil
}

func (m *Bar) _asyncUpdate() error {
	data := []gsheets.UpdateValue{
		{
			SheetName: _Bar_sheetName,
			RowNo:     _Bar_rowNoMap[m.UserID][m.Datetime],
			Values: []interface{}{
				m.UserID,
				m.Datetime.String(),
				m.Value,
				m.Note,
				time.Now(),
				"",
			},
		},
	}
	return dbClient.AsyncUpdate(data)
}

func (m *Bar) _asyncDelete() error {
	now := time.Now()
	data := []gsheets.UpdateValue{
		{
			SheetName: _Bar_sheetName,
			RowNo:     _Bar_rowNoMap[m.UserID][m.Datetime],
			Values: []interface{}{
				m.UserID,
				m.Datetime.String(),
				m.Value,
				m.Note,
				now,
				now,
			},
		},
	}
	return dbClient.AsyncUpdate(data)
}
