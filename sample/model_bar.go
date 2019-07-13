// Code generated by "sheetdb-modeler"; DO NOT EDIT.
// Create a Spreadsheet (sheet name: "bars") as data storage.
// The spreadsheet header is as follows:
//   user_id | datetime | value | note | updated_at | deleted_at
// Please copy and paste this header on the first line of the sheet.

package sample

import (
	"strconv"
	"sync"
	"time"

	"github.com/takuoki/gsheets"
	"github.com/takuoki/sheetdb"
)

const (
	// Sheet definition
	_Bar_sheetName        = "bars"
	_Bar_column_UserID    = 0 // A
	_Bar_column_Datetime  = 1 // B
	_Bar_column_Value     = 2 // C
	_Bar_column_Note      = 3 // D
	_Bar_column_UpdatedAt = 4 // E
	_Bar_column_DeletedAt = 5 // F

	// Parent children relation for compile check
	_Bar_modelSetName_default = 0
	_Bar_parent_User          = 0
	_Bar_numOfChildren        = 0
	_Bar_numOfDirectChildren  = 0
)

var (
	_Bar_mutex    = sync.RWMutex{}
	_Bar_cache    = map[int]map[sheetdb.Datetime]*Bar{} // map[userID][datetime]*Bar
	_Bar_rowNoMap = map[int]map[sheetdb.Datetime]int{}  // map[userID][datetime]rowNo
	_Bar_maxRowNo = 0
)

func _() {
	// An "undeclared name" compiler error signifies that parent-children option conflicts between models.
	// Make sure that the parent-children options are correct for all relevant models and try again.
	_ = _User_child_Bar
}

func init() {
	sheetdb.RegisterModel("default", "Bar", _Bar_sheetName, _Bar_load)
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

		if _, ok := _Bar_cache[userID][datetime]; ok {
			return &sheetdb.DuplicationError{FieldName: "Datetime"}
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
			_Bar_rowNoMap[bar.UserID] = map[sheetdb.Datetime]int{}
		}
		_Bar_cache[bar.UserID][bar.Datetime] = &bar
		_Bar_rowNoMap[bar.UserID][bar.Datetime] = _Bar_maxRowNo
	}

	return nil
}

// GetBar returns a bar by Datetime.
// If it can not be found, this method returns *sheetdb.NotFoundError.
func (m *User) GetBar(datetime sheetdb.Datetime) (*Bar, error) {
	_Bar_mutex.RLock()
	defer _Bar_mutex.RUnlock()
	if v, ok := _Bar_cache[m.UserID][datetime]; ok {
		return v, nil
	}
	return nil, &sheetdb.NotFoundError{Model: "Bar"}
}

// GetBar returns a bar by primary keys.
// If it can not be found, this function returns *sheetdb.NotFoundError.
func GetBar(userID int, datetime sheetdb.Datetime) (*Bar, error) {
	m, err := GetUser(userID)
	if err != nil {
		return nil, err
	}
	return m.GetBar(datetime)
}

// BarQuery is used for selecting bars.
type BarQuery struct {
	filter func(bar *Bar) bool
	sort   func(bars []*Bar)
}

// BarQueryOption is an option to change the behavior of BarQuery.
type BarQueryOption func(query *BarQuery) *BarQuery

// BarFilter is an option to change the filtering behavior of BarQuery.
func BarFilter(filterFunc func(bar *Bar) bool) func(query *BarQuery) *BarQuery {
	return func(query *BarQuery) *BarQuery {
		if query != nil {
			query.filter = filterFunc
		}
		return query
	}
}

// BarSort is an option to change the sorting behavior of BarQuery.
func BarSort(sortFunc func(bars []*Bar)) func(query *BarQuery) *BarQuery {
	return func(query *BarQuery) *BarQuery {
		if query != nil {
			query.sort = sortFunc
		}
		return query
	}
}

// GetBars returns all bars that user has.
// If any options are specified, the result according to the specified option is returned.
// If there are no bar to return, this method returns an nil array.
// If the sort option is not specified, the order of bars is random.
func (m *User) GetBars(opts ...BarQueryOption) ([]*Bar, error) {
	barQuery := &BarQuery{}
	for _, opt := range opts {
		barQuery = opt(barQuery)
	}
	_Bar_mutex.RLock()
	defer _Bar_mutex.RUnlock()
	var bars []*Bar
	if barQuery.filter != nil {
		for _, v := range _Bar_cache[m.UserID] {
			if barQuery.filter(v) {
				bars = append(bars, v)
			}
		}
	} else {
		for _, v := range _Bar_cache[m.UserID] {
			bars = append(bars, v)
		}
	}
	if barQuery.sort != nil {
		barQuery.sort(bars)
	}
	return bars, nil
}

// GetBars returns all bars that user has.
// If any options are specified, the result according to the specified option is returned.
// If there are no bar to return, this function returns an nil array.
// If the sort option is not specified, the order of bars is random.
func GetBars(userID int, opts ...BarQueryOption) ([]*Bar, error) {
	m, err := GetUser(userID)
	if err != nil {
		return nil, err
	}
	return m.GetBars(opts...)
}

// AddBar adds new bar to user.
// If argument 'datetime' already exists in this user, this method returns *sheetdb.DuplicationError.
// If any fields are invalid, this method returns error.
func (m *User) AddBar(datetime sheetdb.Datetime, value float32, note string) (*Bar, error) {
	_Bar_mutex.Lock()
	defer _Bar_mutex.Unlock()
	if _, ok := _Bar_cache[m.UserID][datetime]; ok {
		return nil, &sheetdb.DuplicationError{FieldName: "Datetime"}
	}
	if err := _Bar_validateNote(note); err != nil {
		return nil, err
	}
	bar := &Bar{
		UserID:   m.UserID,
		Datetime: datetime,
		Value:    value,
		Note:     note,
	}
	if err := bar._asyncAdd(_Bar_maxRowNo + 1); err != nil {
		return nil, err
	}
	_Bar_maxRowNo++
	if _, ok := _Bar_cache[bar.UserID]; !ok {
		_Bar_cache[bar.UserID] = map[sheetdb.Datetime]*Bar{}
		_Bar_rowNoMap[bar.UserID] = map[sheetdb.Datetime]int{}
	}
	_Bar_cache[bar.UserID][bar.Datetime] = bar
	_Bar_rowNoMap[bar.UserID][bar.Datetime] = _Bar_maxRowNo
	return bar, nil
}

// AddBar adds new bar to user.
// If primary keys already exist, this function returns *sheetdb.DuplicationError.
// If any fields are invalid, this function returns error.
func AddBar(userID int, datetime sheetdb.Datetime, value float32, note string) (*Bar, error) {
	m, err := GetUser(userID)
	if err != nil {
		return nil, err
	}
	return m.AddBar(datetime, value, note)
}

// UpdateBar updates bar.
// If it can not be found, this method returns *sheetdb.NotFoundError.
// If any fields are invalid, this method returns error.
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
	*bar = barCopy
	return bar, nil
}

// UpdateBar updates bar.
// If it can not be found, this function returns *sheetdb.NotFoundError.
// If any fields are invalid, this function returns error.
func UpdateBar(userID int, datetime sheetdb.Datetime, value float32, note string) (*Bar, error) {
	m, err := GetUser(userID)
	if err != nil {
		return nil, err
	}
	return m.UpdateBar(datetime, value, note)
}

// DeleteBar deletes bar from user.
// If it can not be found, this method returns *sheetdb.NotFoundError.
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

// DeleteBar deletes bar from user.
// If it can not be found, this function returns *sheetdb.NotFoundError.
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
		return sheetdb.Datetime{}, &sheetdb.InvalidValueError{FieldName: "Datetime", Err: err}
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

func (m *Bar) _asyncAdd(rowNo int) error {
	data := []gsheets.UpdateValue{
		{
			SheetName: _Bar_sheetName,
			RowNo:     rowNo,
			Values: []interface{}{
				m.UserID,
				m.Datetime.String(),
				m.Value,
				m.Note,
				time.Now().UTC(),
				"",
			},
		},
	}
	return dbClient.AsyncUpdate(data)
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
				time.Now().UTC(),
				"",
			},
		},
	}
	return dbClient.AsyncUpdate(data)
}

func (m *Bar) _asyncDelete() error {
	now := time.Now().UTC()
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
