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
	_User_sheetName        = "users"
	_User_column_UserID    = 0 // A
	_User_column_Name      = 1 // B
	_User_column_Email     = 2 // C
	_User_column_Sex       = 3 // D
	_User_column_Birthday  = 4 // E
	_User_column_UpdatedAt = 5 // F
	_User_column_DeletedAt = 6 // G
)

var (
	_User_mutex           = sync.RWMutex{}
	_User_cache           = map[int]*User{} // map[userID]*User
	_User_rowNoMap        = map[int]int{}   // map[userID]rowNo
	_User_maxRowNo        int
	_User_Email_uniqueSet = map[string]struct{}{}
)

func init() {
	dbClient.AddModel("User", _User_load)
}

func _User_load(data *gsheets.Sheet) error {

	_User_mutex.Lock()
	defer _User_mutex.Unlock()

	_User_cache = map[int]*User{}
	_User_rowNoMap = map[int]int{}
	_User_maxRowNo = 0

	for i, r := range data.Rows() {
		if i == 0 {
			continue
		}
		if r.Value(_User_column_DeletedAt) != "" {
			_User_maxRowNo++
			continue
		}
		if r.Value(_User_column_UserID) == "" {
			break
		}

		userID, err := _User_parseUserID(r.Value(_User_column_UserID))
		if err != nil {
			return err
		}
		name := r.Value(_User_column_Name)
		if err := _User_validateName(name); err != nil {
			return err
		}
		email := r.Value(_User_column_Email)
		if err := _User_validateEmail(email, nil); err != nil {
			return err
		}
		sex, err := _User_parseSex(r.Value(_User_column_Sex))
		if err != nil {
			return err
		}
		birthday, err := _User_parseBirthday(r.Value(_User_column_Birthday))
		if err != nil {
			return err
		}

		user := User{
			UserID:   userID,
			Name:     name,
			Email:    email,
			Sex:      sex,
			Birthday: birthday,
		}

		_User_Email_uniqueSet[user.Email] = struct{}{}
		_User_maxRowNo++
		_User_cache[user.UserID] = &user
		_User_rowNoMap[user.UserID] = _User_maxRowNo
	}

	return nil
}

func GetUser(userID int) (*User, error) {
	_User_mutex.RLock()
	defer _User_mutex.RUnlock()
	if v, ok := _User_cache[userID]; ok {
		return v, nil
	}
	return nil, &sheetdb.NotFoundError{Model: "User"}
}

type UserQuery struct {
	filter func(user *User) bool
	sort   func(users []*User)
}

type UserQueryOption func(query *UserQuery) *UserQuery

func UserFilter(filterFunc func(user *User) bool) func(query *UserQuery) *UserQuery {
	return func(query *UserQuery) *UserQuery {
		if query != nil {
			query.filter = filterFunc
		}
		return query
	}
}

func UserSort(sortFunc func(users []*User)) func(query *UserQuery) *UserQuery {
	return func(query *UserQuery) *UserQuery {
		if query != nil {
			query.sort = sortFunc
		}
		return query
	}
}

func GetUsers(opts ...UserQueryOption) ([]*User, error) {
	userQuery := &UserQuery{}
	for _, opt := range opts {
		userQuery = opt(userQuery)
	}
	_User_mutex.RLock()
	defer _User_mutex.RUnlock()
	users := []*User{}
	if userQuery.filter != nil {
		for _, v := range _User_cache {
			if userQuery.filter(v) {
				users = append(users, v)
			}
		}
	} else {
		for _, v := range _User_cache {
			users = append(users, v)
		}
	}
	if userQuery.sort != nil {
		userQuery.sort(users)
	}
	return users, nil
}

func AddUser(name string, email string, sex Sex, birthday *sheetdb.Date) (*User, error) {
	_User_mutex.Lock()
	defer _User_mutex.Unlock()
	if err := _User_validateName(name); err != nil {
		return nil, err
	}
	if err := _User_validateEmail(email, nil); err != nil {
		return nil, err
	}
	user := &User{
		UserID:   _User_maxRowNo + 10001,
		Name:     name,
		Email:    email,
		Sex:      sex,
		Birthday: birthday,
	}
	if err := user._asyncUpdate(); err != nil {
		return nil, err
	}
	_User_Email_uniqueSet[user.Email] = struct{}{}
	_User_maxRowNo++
	_User_cache[user.UserID] = user
	_User_rowNoMap[user.UserID] = _User_maxRowNo
	return user, nil
}

func UpdateUser(userID int, name string, email string, sex Sex, birthday *sheetdb.Date) (*User, error) {
	_User_mutex.Lock()
	defer _User_mutex.Unlock()
	user, ok := _User_cache[userID]
	if !ok {
		return nil, &sheetdb.NotFoundError{Model: "User"}
	}
	if err := _User_validateName(name); err != nil {
		return nil, err
	}
	if err := _User_validateEmail(email, &user.Email); err != nil {
		return nil, err
	}
	userCopy := *user
	userCopy.Name = name
	userCopy.Email = email
	userCopy.Sex = sex
	userCopy.Birthday = birthday
	if err := (&userCopy)._asyncUpdate(); err != nil {
		return nil, err
	}
	if userCopy.Email != user.Email {
		_User_Email_uniqueSet[userCopy.Email] = struct{}{}
		delete(_User_Email_uniqueSet, user.Email)
	}
	user = &userCopy
	return user, nil
}

func DeleteUser(userID int) error {
	_User_mutex.Lock()
	defer _User_mutex.Unlock()
	_Foo_mutex.Lock()
	defer _Foo_mutex.Unlock()
	_Bar_mutex.Lock()
	defer _Bar_mutex.Unlock()
	user, ok := _User_cache[userID]
	if !ok {
		return &sheetdb.NotFoundError{Model: "User"}
	}
	var foos []*Foo
	for _, v := range _Foo_cache[userID] {
		foos = append(foos, v)
	}
	var bars []*Bar
	for _, v := range _Bar_cache[userID] {
		bars = append(bars, v)
	}
	if err := user._asyncDelete(foos, bars); err != nil {
		return err
	}
	delete(_User_Email_uniqueSet, user.Email)
	delete(_User_cache, userID)
	delete(_Foo_cache, userID)
	delete(_Bar_cache, userID)
	return nil
}

func _User_validateName(name string) error {
	if name == "" {
		return &sheetdb.EmptyStringError{FieldName: "Name"}
	}
	return nil
}

func _User_validateEmail(email string, oldEmail *string) error {
	if email == "" {
		return &sheetdb.EmptyStringError{FieldName: "Email"}
	}
	if oldEmail == nil || *oldEmail != email {
		if _, ok := _User_Email_uniqueSet[email]; ok {
			return &sheetdb.DuplicationError{FieldName: "Email"}
		}
	}
	return nil
}

func _User_parseUserID(userID string) (int, error) {
	v, err := strconv.Atoi(userID)
	if err != nil {
		return 0, &sheetdb.InvalidValueError{FieldName: "UserID", Err: err}
	}
	return v, nil
}

func _User_parseSex(sex string) (Sex, error) {
	v, err := NewSex(sex)
	if err != nil {
		return v, &sheetdb.InvalidValueError{FieldName: "Sex", Err: err}
	}
	return v, nil
}

func _User_parseBirthday(birthday string) (*sheetdb.Date, error) {
	var val *sheetdb.Date
	if birthday != "" {
		v, err := sheetdb.NewDate(birthday)
		if err != nil {
			return nil, &sheetdb.InvalidValueError{FieldName: "Birthday", Err: err}
		}
		val = &v
	}
	return val, nil
}

func (m *User) _asyncUpdate() error {
	data := []gsheets.UpdateValue{
		{
			SheetName: _User_sheetName,
			RowNo:     _User_rowNoMap[m.UserID],
			Values: []interface{}{
				m.UserID,
				m.Name,
				m.Email,
				m.Sex.String(),
				m.Birthday.String(),
				time.Now(),
				"",
			},
		},
	}
	return dbClient.AsyncUpdate(data)
}

func (m *User) _asyncDelete(foos []*Foo, bars []*Bar) error {
	now := time.Now()
	data := []gsheets.UpdateValue{
		{
			SheetName: _User_sheetName,
			RowNo:     _User_rowNoMap[m.UserID],
			Values: []interface{}{
				m.UserID,
				m.Name,
				m.Email,
				m.Sex.String(),
				m.Birthday.String(),
				now,
				now,
			},
		},
	}
	for _, v := range foos {
		data = append(data, gsheets.UpdateValue{
			SheetName: _Foo_sheetName,
			RowNo:     _Foo_rowNoMap[v.UserID][v.FooID],
			Values: []interface{}{
				v.UserID,
				v.FooID,
				v.Value,
				v.Note,
				now,
				now,
			},
		})
	}
	for _, v := range bars {
		data = append(data, gsheets.UpdateValue{
			SheetName: _Bar_sheetName,
			RowNo:     _Bar_rowNoMap[v.UserID][v.Datetime],
			Values: []interface{}{
				v.UserID,
				v.Datetime.String(),
				v.Value,
				v.Note,
				now,
				now,
			},
		})
	}
	return dbClient.AsyncUpdate(data)
}
