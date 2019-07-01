package sample

import (
	"strconv"
	"sync"
	"time"

	"github.com/takuoki/gsheets"
	"github.com/takuoki/sheetdb"
)

const (
	_Foo_sheetName        = "foos"
	_Foo_column_UserID    = 0 // A
	_Foo_column_FooID     = 1 // B
	_Foo_column_Value     = 2 // C
	_Foo_column_Note      = 3 // D
	_Foo_column_UpdatedAt = 4 // E
	_Foo_column_DeletedAt = 5 // F
)

var (
	_Foo_mutex    = sync.RWMutex{}
	_Foo_cache    = map[int]map[int]*Foo{}
	_Foo_rowNoMap = map[int]map[int]int{}
	_Foo_maxRowNo int
)

func init() {
	dbClient.AddModel("Foo", _Foo_load)
}

func _Foo_load(data *gsheets.Sheet) error {

	_Foo_mutex.Lock()
	defer _Foo_mutex.Unlock()

	_Foo_cache = map[int]map[int]*Foo{}
	_Foo_rowNoMap = map[int]map[int]int{}
	_Foo_maxRowNo = 0

	for i, r := range data.Rows() {
		if i == 0 {
			continue
		}
		if r.Value(_Foo_column_DeletedAt) != "" {
			_Foo_maxRowNo++
			continue
		}
		if r.Value(_Foo_column_UserID) == "" {
			break
		}

		userID, err := _Foo_parseUserID(r.Value(_Foo_column_UserID))
		if err != nil {
			return err
		}
		fooID, err := _Foo_parseFooID(r.Value(_Foo_column_FooID))
		if err != nil {
			return err
		}
		value, err := _Foo_parseValue(r.Value(_Foo_column_Value))
		if err != nil {
			return err
		}
		note := r.Value(_Foo_column_Note)
		if err := _Foo_validateNote(note); err != nil {
			return err
		}

		foo := Foo{
			UserID: userID,
			FooID:  fooID,
			Value:  value,
			Note:   note,
		}

		_Foo_maxRowNo++
		if _, ok := _Foo_cache[foo.UserID]; !ok {
			_Foo_cache[foo.UserID] = map[int]*Foo{}
		}
		_Foo_cache[userID][fooID] = &foo
		_User_rowNoMap[userID] = _User_maxRowNo
	}

	return nil
}

func GetFoo(userID int, fooID int) (*Foo, error) {
	_Foo_mutex.RLock()
	defer _Foo_mutex.RUnlock()
	if v, ok := _Foo_cache[userID][fooID]; ok {
		return v, nil
	}
	return nil, &sheetdb.NotFoundError{Model: "Foo"}
}

type FooQuery struct {
	filter func(foo *Foo) bool
	sort   func(foos []*Foo)
}

type FooQueryOption func(query *FooQuery) *FooQuery

func FooFilter(filterFunc func(foo *Foo) bool) func(query *FooQuery) *FooQuery {
	return func(query *FooQuery) *FooQuery {
		if query != nil {
			query.filter = filterFunc
		}
		return query
	}
}

func FooSort(sortFunc func(foos []*Foo)) func(query *FooQuery) *FooQuery {
	return func(query *FooQuery) *FooQuery {
		if query != nil {
			query.sort = sortFunc
		}
		return query
	}
}

func GetFoos(userID int, opts ...FooQueryOption) ([]*Foo, error) {
	fooQuery := &FooQuery{}
	for _, opt := range opts {
		fooQuery = opt(fooQuery)
	}
	_Foo_mutex.RLock()
	defer _Foo_mutex.RUnlock()
	foos := []*Foo{}
	if fooQuery.filter != nil {
		for _, v := range _Foo_cache[userID] {
			if fooQuery.filter(v) {
				foos = append(foos, v)
			}
		}
	} else {
		for _, v := range _Foo_cache[userID] {
			foos = append(foos, v)
		}
	}
	if fooQuery.sort != nil {
		fooQuery.sort(foos)
	}
	return foos, nil
}

func (m *User) AddFoo(value float32, note string) (*Foo, error) {
	_Foo_mutex.Lock()
	defer _Foo_mutex.Unlock()
	if err := _Foo_validateNote(note); err != nil {
		return nil, err
	}
	foo := &Foo{
		UserID: m.UserID,
		FooID:  _Foo_maxRowNo + 10001,
		Value:  value,
		Note:   note,
	}
	if err := foo._asyncUpdate(); err != nil {
		return nil, err
	}
	_Foo_maxRowNo++
	if _, ok := _Foo_cache[foo.UserID]; !ok {
		_Foo_cache[foo.UserID] = map[int]*Foo{}
	}
	_Foo_cache[foo.UserID][foo.FooID] = foo
	_Foo_rowNoMap[foo.UserID][foo.FooID] = _Foo_maxRowNo
	return foo, nil
}

func AddFoo(userID int, value float32, note string) (*Foo, error) {
	m, err := GetUser(userID)
	if err != nil {
		return nil, err
	}
	return m.AddFoo(value, note)
}

func (m *User) UpdateFoo(fooID int, value float32, note string) (*Foo, error) {
	_Foo_mutex.Lock()
	defer _Foo_mutex.Unlock()
	if err := _Foo_validateNote(note); err != nil {
		return nil, err
	}
	foo, ok := _Foo_cache[m.UserID][fooID]
	if !ok {
		return nil, &sheetdb.NotFoundError{Model: "Foo"}
	}
	fooCopy := *foo
	fooCopy.Value = value
	fooCopy.Note = note
	if err := (&fooCopy)._asyncUpdate(); err != nil {
		return nil, err
	}
	foo = &fooCopy
	return foo, nil
}

func UpdateFoo(userID int, fooID int, value float32, note string) (*Foo, error) {
	m, err := GetUser(userID)
	if err != nil {
		return nil, err
	}
	return m.UpdateFoo(fooID, value, note)
}

func (m *User) DeleteFoo(fooID int) error {
	_Foo_mutex.Lock()
	defer _Foo_mutex.Unlock()
	foo, ok := _Foo_cache[m.UserID][fooID]
	if !ok {
		return &sheetdb.NotFoundError{Model: "Foo"}
	}
	if err := foo._asyncDelete(); err != nil {
		return err
	}
	delete(_Foo_cache[m.UserID], fooID)
	return nil
}

func DeleteFoo(userID int, fooID int) error {
	m, err := GetUser(userID)
	if err != nil {
		return err
	}
	return m.DeleteFoo(fooID)
}

func _Foo_validateNote(note string) error {
	return nil
}

func _Foo_parseUserID(userID string) (int, error) {
	v, err := strconv.Atoi(userID)
	if err != nil {
		return 0, &sheetdb.InvalidValueError{FieldName: "UserID", Err: err}
	}
	return v, nil
}

func _Foo_parseFooID(fooID string) (int, error) {
	v, err := strconv.Atoi(fooID)
	if err != nil {
		return 0, &sheetdb.InvalidValueError{FieldName: "FooID", Err: err}
	}
	return v, nil
}

func _Foo_parseValue(value string) (float32, error) {
	v, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0, &sheetdb.InvalidValueError{FieldName: "Value", Err: err}
	}
	return float32(v), nil
}

func (m *Foo) _asyncUpdate() error {
	data := []gsheets.UpdateValue{
		{
			SheetName: _Foo_sheetName,
			RowNo:     _Foo_rowNoMap[m.UserID][m.FooID],
			Values: []interface{}{
				m.UserID,
				m.FooID,
				m.Value,
				m.Note,
				time.Now(),
				"",
			},
		},
	}
	return dbClient.AsyncUpdate(data)
}

func (m *Foo) _asyncDelete() error {
	now := time.Now()
	data := []gsheets.UpdateValue{
		{
			SheetName: _Foo_sheetName,
			RowNo:     _Foo_rowNoMap[m.UserID][m.FooID],
			Values: []interface{}{
				m.UserID,
				m.FooID,
				m.Value,
				m.Note,
				now,
				now,
			},
		},
	}
	return dbClient.AsyncUpdate(data)
}
