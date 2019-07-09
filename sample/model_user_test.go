package sample_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/takuoki/sheetdb"
	"github.com/takuoki/sheetdb/sample"
)

func TestGetUser(t *testing.T) {
	cases := map[string]struct {
		id          int
		name, email string
		sex         sample.Sex
		birthday    *sheetdb.Date
		notFound    bool
	}{
		"success": {
			id:       10001,
			name:     "Jorge B. Farley",
			email:    "jorge.b.farley@sample.com",
			sex:      sample.Male,
			birthday: &date19590525,
		},
		"empty-birthday": {
			id:       10005,
			name:     "Judith C. Thrash",
			email:    "judith.c.thrash@sample.com",
			sex:      sample.Female,
			birthday: nil,
		},
		"deleted": {
			id:       10006,
			notFound: true,
		},
		"not-found": {
			id:       10007,
			notFound: true,
		},
	}
	sample.Reload(t)
	for casename, c := range cases {
		t.Run(casename, func(t *testing.T) {
			user, err := sample.GetUser(c.id)
			if !c.notFound {
				if err != nil {
					t.Errorf("User must be found (case: %s)", casename)
					return
				}
				if user.Name != c.name {
					t.Errorf("Name does not match expected (case: %s, expected=%s, actual=%s)", casename, c.name, user.Name)
				}
				if user.Email != c.email {
					t.Errorf("Email does not match expected (case: %s, expected=%s, actual=%s)", casename, c.email, user.Email)
				}
				if user.Sex != c.sex {
					t.Errorf("Sex does not match expected (case: %s, expected=%s, actual=%s)", casename, c.sex, user.Sex)
				}
				if !reflect.DeepEqual(user.Birthday, c.birthday) {
					t.Errorf("Birthday does not match expected (case: %s, expected=%+v, actual=%+v)", casename, c.birthday, user.Birthday)
				}
			} else {
				if err == nil {
					t.Errorf("Error must occur (case: %s)", casename)
					return
				}
				if e, ok := err.(*sheetdb.NotFoundError); !ok {
					t.Errorf("Error must be sheetd.NotFoundError (case: %s, actual=%T)", casename, err)
				} else if e.Model != "User" {
					t.Errorf("Error model must be 'User' (case: %s, actual=%s)", casename, e.Model)
				}
			}
		})
	}
}

func TestGetUseryEmail(t *testing.T) {
	cases := map[string]struct {
		id          int
		name, email string
		notFound    bool
	}{
		"success": {
			id:    10001,
			name:  "Jorge B. Farley",
			email: "jorge.b.farley@sample.com",
		},
		"deleted": {
			email:    "mark.f.oswald@sample.com",
			notFound: true,
		},
		"not-found": {
			email:    "betty.m.sinclair@sample.com",
			notFound: true,
		},
	}
	sample.Reload(t)
	for casename, c := range cases {
		t.Run(casename, func(t *testing.T) {
			user, err := sample.GetUserByEmail(c.email)
			if !c.notFound {
				if err != nil {
					t.Errorf("User must be found (case: %s)", casename)
					return
				}
				if user.UserID != c.id {
					t.Errorf("UserID does not match expected (case: %s, expected=%d, actual=%d)", casename, c.id, user.UserID)
				}
				if user.Name != c.name {
					t.Errorf("Name does not match expected (case: %s, expected=%s, actual=%s)", casename, c.name, user.Name)
				}
			} else {
				if err == nil {
					t.Errorf("Error must occur (case: %s)", casename)
					return
				}
				if e, ok := err.(*sheetdb.NotFoundError); !ok {
					t.Errorf("Error must be sheetd.NotFoundError (case: %s, actual=%T)", casename, err)
				} else if e.Model != "User" {
					t.Errorf("Error model must be 'User' (case: %s, actual=%s)", casename, e.Model)
				}
			}
		})
	}
}

func TestGetUsers(t *testing.T) {
	cases := map[string]struct {
		filterFunc  func(user *sample.User) bool
		sortFunc    func(users []*sample.User)
		expectedIDs []int
	}{
		"no-filter-and-sort": {
			expectedIDs: []int{10001, 10002, 10003, 10004, 10005},
		},
		"filter-and-sort": {
			filterFunc: func(user *sample.User) bool {
				return user.Sex == sample.Male
			},
			sortFunc: func(users []*sample.User) {
				sort.Slice(users, func(i, j int) bool {
					return users[i].UserID > users[j].UserID
				})
			},
			expectedIDs: []int{10004, 10002, 10001},
		},
		"no-result": {
			filterFunc: func(user *sample.User) bool {
				return false
			},
			expectedIDs: nil,
		},
	}
	sample.Reload(t)
	for casename, c := range cases {
		t.Run(casename, func(t *testing.T) {
			users, err := sample.GetUsers(sample.UserFilter(c.filterFunc), sample.UserSort(c.sortFunc))
			if err != nil {
				t.Errorf("Error must not occur (case: %s, err=%v)", casename, err)
				return
			}
			var userIDs []int
			for _, user := range users {
				userIDs = append(userIDs, user.UserID)
			}
			if len(users) != len(c.expectedIDs) {
				t.Errorf("The number of users does not match expected (case: %s, expected=%v, actual=%v)", casename, c.expectedIDs, userIDs)
			} else if c.sortFunc != nil && !reflect.DeepEqual(userIDs, c.expectedIDs) {
				t.Errorf("The order of users does not match expected (case: %s, expected=%v, actual=%v)", casename, c.expectedIDs, userIDs)
			}
		})
	}
}

func TestAddUser(t *testing.T) {
	cases := map[string]struct {
		name, email  string
		sex          sample.Sex
		birthday     *sheetdb.Date
		expectedUser sample.User
		err          error
	}{
		"success": {
			name:     "Betty M. Sinclair",
			email:    "betty.m.sinclair@sample.com",
			sex:      sample.Female,
			birthday: &date19590525,
			expectedUser: sample.User{
				UserID:   10007,
				Name:     "Betty M. Sinclair",
				Email:    "betty.m.sinclair@sample.com",
				Sex:      sample.Female,
				Birthday: &date19590525,
			},
		},
		"name-empty": {
			name:     "",
			email:    "betty.m.sinclair@sample.com",
			sex:      sample.Female,
			birthday: &date19590525,
			err:      &sheetdb.EmptyStringError{FieldName: "Name"},
		},
		"email-empty": {
			name:     "Betty M. Sinclair",
			email:    "",
			sex:      sample.Female,
			birthday: &date19590525,
			err:      &sheetdb.EmptyStringError{FieldName: "Email"},
		},
		"email-duplicate": {
			name:     "Betty M. Sinclair",
			email:    "kathy.m.fisher@sample.com",
			sex:      sample.Female,
			birthday: &date19590525,
			err:      &sheetdb.DuplicationError{FieldName: "Email"},
		},
	}
	for casename, c := range cases {
		t.Run(casename, func(t *testing.T) {
			sample.Reload(t)
			user, err := sample.AddUser(c.name, c.email, c.sex, c.birthday)
			if c.err == nil {
				if err != nil {
					t.Errorf("Error must not occur in AddUser (case: %s, err=%v)", casename, err)
					return
				}
				if !reflect.DeepEqual(user, &c.expectedUser) {
					t.Errorf("User that AddUser returns does not match expected (case: %s, expected=%+v, actual=%+v)", casename, &c.expectedUser, user)
					return
				}
				user, err := sample.GetUser(c.expectedUser.UserID)
				if err != nil {
					t.Errorf("Error must not occur in GetUser (case: %s, err=%v)", casename, err)
					return
				}
				if !reflect.DeepEqual(user, &c.expectedUser) {
					t.Errorf("User that GetUser returns does not match expected (case: %s, expected=%+v, actual=%+v)", casename, &c.expectedUser, user)
					return
				}
				user, err = sample.GetUserByEmail(c.expectedUser.Email)
				if err != nil {
					t.Errorf("Error must not occur in GetUserByEmail (case: %s, err=%v)", casename, err)
					return
				}
				if !reflect.DeepEqual(user, &c.expectedUser) {
					t.Errorf("User that GetUserByEmail returns does not match expected (case: %s, expected=%+v, actual=%+v)", casename, &c.expectedUser, user)
					return
				}
			} else {
				if err == nil {
					t.Errorf("Error must occur (case: %s)", casename)
					return
				}
				if !reflect.DeepEqual(err, c.err) {
					t.Errorf("Error does not match expected (case: %s, expected=%+v, actual=%+v)", casename, c.err, err)
					return
				}
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	cases := map[string]struct {
		id           int
		name, email  string
		sex          sample.Sex
		birthday     *sheetdb.Date
		expectedUser sample.User
		err          error
	}{
		"update-all": {
			id:       10004,
			name:     "Betty M. Sinclair",
			email:    "betty.m.sinclair@sample.com",
			sex:      sample.Female,
			birthday: &date19590525,
			expectedUser: sample.User{
				UserID:   10004,
				Name:     "Betty M. Sinclair",
				Email:    "betty.m.sinclair@sample.com",
				Sex:      sample.Female,
				Birthday: &date19590525,
			},
		},
		"update-name": {
			id:       10004,
			name:     "Betty M. Sinclair",
			email:    "matthew.j.mclane@sample.com",
			sex:      sample.Male,
			birthday: &date19950914,
			expectedUser: sample.User{
				UserID:   10004,
				Name:     "Betty M. Sinclair",
				Email:    "matthew.j.mclane@sample.com",
				Sex:      sample.Male,
				Birthday: &date19950914,
			},
		},
		"not-found": {
			id:  10007,
			err: &sheetdb.NotFoundError{Model: "User"},
		},
		"name-empty": {
			id:       10004,
			name:     "",
			email:    "betty.m.sinclair@sample.com",
			sex:      sample.Female,
			birthday: &date19590525,
			err:      &sheetdb.EmptyStringError{FieldName: "Name"},
		},
		"email-empty": {
			id:       10004,
			name:     "Betty M. Sinclair",
			email:    "",
			sex:      sample.Female,
			birthday: &date19590525,
			err:      &sheetdb.EmptyStringError{FieldName: "Email"},
		},
		"email-duplicate": {
			id:       10004,
			name:     "Betty M. Sinclair",
			email:    "kathy.m.fisher@sample.com",
			sex:      sample.Female,
			birthday: &date19590525,
			err:      &sheetdb.DuplicationError{FieldName: "Email"},
		},
	}
	for casename, c := range cases {
		t.Run(casename, func(t *testing.T) {
			sample.Reload(t)
			oldUser, _ := sample.GetUser(c.id)
			user, err := sample.UpdateUser(c.id, c.name, c.email, c.sex, c.birthday)
			if c.err == nil {
				if err != nil {
					t.Errorf("Error must not occur in UpdateUser (case: %s, err=%v)", casename, err)
					return
				}
				if !reflect.DeepEqual(user, &c.expectedUser) {
					t.Errorf("User that UpdateUser returns does not match expected (case: %s, expected=%+v, actual=%+v)", casename, &c.expectedUser, user)
					return
				}
				user, err := sample.GetUser(c.id)
				if err != nil {
					t.Errorf("Error must not occur in GetUser (case: %s, err=%v)", casename, err)
					return
				}
				if !reflect.DeepEqual(user, &c.expectedUser) {
					t.Errorf("User that GetUser returns does not match expected (case: %s, expected=%+v, actual=%+v)", casename, &c.expectedUser, user)
					return
				}
				user, err = sample.GetUserByEmail(c.email)
				if err != nil {
					t.Errorf("Error must not occur in GetUserByEmail (case: %s, err=%v)", casename, err)
					return
				}
				if !reflect.DeepEqual(user, &c.expectedUser) {
					t.Errorf("User that GetUserByEmail returns does not match expected (case: %s, expected=%+v, actual=%+v)", casename, &c.expectedUser, user)
					return
				}
				if c.email != oldUser.Email {
					if _, err := sample.GetUserByEmail(oldUser.Email); !reflect.DeepEqual(err, &sheetdb.NotFoundError{Model: "User"}) {
						t.Errorf("Error must occur when GetUserByEmail is called by old email (case: %s, err=%v)", casename, err)
						return
					}
				}
			} else {
				if err == nil {
					t.Errorf("Error must occur (case: %s)", casename)
					return
				}
				if !reflect.DeepEqual(err, c.err) {
					t.Errorf("Error does not match expected (case: %s, expected=%+v, actual=%+v)", casename, c.err, err)
					return
				}
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	cases := map[string]struct {
		id             int
		email          string
		fooIDs         []int
		fooChildIDs    [][2]int
		fooChildValues []string
		barIDs         []sheetdb.Datetime
		err            error
	}{
		"delete-single": {
			id:    10005,
			email: "judith.c.thrash@sample.com",
		},
		"delete-with-children": {
			id:             10002,
			email:          "guillermo.l.shanks@sample.com",
			fooIDs:         []int{1, 2, 3},
			fooChildIDs:    [][2]int{{1, 1}, {2, 1}},
			fooChildValues: []string{"b", "h"},
			barIDs:         []sheetdb.Datetime{datetime20190707000000},
		},
		"not-found": {
			id:  10007,
			err: &sheetdb.NotFoundError{Model: "User"},
		},
	}
	for casename, c := range cases {
		t.Run(casename, func(t *testing.T) {
			sample.Reload(t)
			// pre-check
			if c.err == nil {
				if _, err := sample.GetUserByEmail(c.email); err != nil {
					t.Errorf("[Pre-check] Error must not occur in GetUserByEmail (case: %s, email=%s, err=%v)", casename, c.email, err)
					return
				}
				for _, fooID := range c.fooIDs {
					if _, err := sample.GetFoo(c.id, fooID); err != nil {
						t.Errorf("[Pre-check] Error must not occur in GetFoo (case: %s, fooID=%d, err=%v)", casename, fooID, err)
						return
					}
				}
				for _, id := range c.fooChildIDs {
					if _, err := sample.GetFooChild(c.id, id[0], id[1]); err != nil {
						t.Errorf("[Pre-check] Error must not occur in GetFooChild (case: %s, fooID=%d, fooChildID=%d, err=%v)", casename, id[0], id[1], err)
						return
					}
				}
				for _, v := range c.fooChildValues {
					if _, err := sample.GetFooChildByValue(v); err != nil {
						t.Errorf("[Pre-check] Error must not occur in GetFooChildByValue (case: %s, value=%s, err=%v)", casename, v, err)
						return
					}
				}
				for _, barID := range c.barIDs {
					if _, err := sample.GetBar(c.id, barID); err != nil {
						t.Errorf("[Pre-check] Error must not occur in GetBar (case: %s, barID=%s, err=%v)", casename, barID, err)
						return
					}
				}
			}
			err := sample.DeleteUser(c.id)
			if c.err == nil {
				if err != nil {
					t.Errorf("Error must not occur in DeleteUser (case: %s, err=%v)", casename, err)
					return
				}
				if _, err := sample.GetUser(c.id); !reflect.DeepEqual(err, &sheetdb.NotFoundError{Model: "User"}) {
					t.Errorf("Error in GetUser does not match expected (case: %s, err=%v)", casename, err)
					return
				}
				if _, err := sample.GetUserByEmail(c.email); !reflect.DeepEqual(err, &sheetdb.NotFoundError{Model: "User"}) {
					t.Errorf("Error in GetUserByEmail does not match expected (case: %s, email=%s, err=%v)", casename, c.email, err)
					return
				}
				for _, fooID := range c.fooIDs {
					if _, err := sample.GetFoo(c.id, fooID); !reflect.DeepEqual(err, &sheetdb.NotFoundError{Model: "User"}) {
						t.Errorf("Error in GetFoo does not match expected (case: %s, fooID=%d, err=%v)", casename, fooID, err)
						return
					}
				}
				for _, id := range c.fooChildIDs {
					if _, err := sample.GetFooChild(c.id, id[0], id[1]); !reflect.DeepEqual(err, &sheetdb.NotFoundError{Model: "User"}) {
						t.Errorf("Error in GetFooChild does not match expected (case: %s, fooID=%d, fooChildID=%d, err=%v)", casename, id[0], id[1], err)
						return
					}
				}
				for _, v := range c.fooChildValues {
					if _, err := sample.GetFooChildByValue(v); !reflect.DeepEqual(err, &sheetdb.NotFoundError{Model: "FooChild"}) {
						t.Errorf("Error in GetFooChildByValue does not match expected (case: %s, value=%s, err=%v)", casename, v, err)
						return
					}
				}
				for _, barID := range c.barIDs {
					if _, err := sample.GetBar(c.id, barID); !reflect.DeepEqual(err, &sheetdb.NotFoundError{Model: "User"}) {
						t.Errorf("Error in GetBar does not match expected (case: %s, barID=%s, err=%v)", casename, barID, err)
						return
					}
				}
			} else {
				if err == nil {
					t.Errorf("Error must occur (case: %s)", casename)
					return
				}
				if !reflect.DeepEqual(err, c.err) {
					t.Errorf("Error does not match expected (case: %s, expected=%+v, actual=%+v)", casename, c.err, err)
					return
				}
			}
		})
	}
}
