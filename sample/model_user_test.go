package sample_test

import (
	"context"
	"reflect"
	"sort"
	"testing"

	"github.com/takuoki/sheetdb"
	"github.com/takuoki/sheetdb/sample"
)

func TestGetUser(t *testing.T) {
	if err := sample.LoadData(context.Background()); err != nil {
		t.Fatalf("Unable to load data from spreadsheet: %v", err)
	}
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
			birthday: &datetime19590525,
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
	for casename, c := range cases {
		user, err := sample.GetUser(c.id)
		if !c.notFound {
			if err != nil {
				t.Errorf("User must be found (case: %s)", casename)
				continue
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
				continue
			}
			if e, ok := err.(*sheetdb.NotFoundError); !ok {
				t.Errorf("Error must be sheetd.NotFoundError (case: %s, actual=%T)", casename, err)
			} else if e.Model != "User" {
				t.Errorf("Error model must be 'User' (case: %s, actual=%s)", casename, e.Model)
			}
		}
	}
}

func TestGetUseryEmail(t *testing.T) {
	if err := sample.LoadData(context.Background()); err != nil {
		t.Fatalf("Unable to load data from spreadsheet: %v", err)
	}
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
	for casename, c := range cases {
		user, err := sample.GetUserByEmail(c.email)
		if !c.notFound {
			if err != nil {
				t.Errorf("User must be found (case: %s)", casename)
				continue
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
				continue
			}
			if e, ok := err.(*sheetdb.NotFoundError); !ok {
				t.Errorf("Error must be sheetd.NotFoundError (case: %s, actual=%T)", casename, err)
			} else if e.Model != "User" {
				t.Errorf("Error model must be 'User' (case: %s, actual=%s)", casename, e.Model)
			}
		}
	}
}

func TestGetUsers(t *testing.T) {
	if err := sample.LoadData(context.Background()); err != nil {
		t.Fatalf("Unable to load data from spreadsheet: %v", err)
	}
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
	for casename, c := range cases {
		users, err := sample.GetUsers(sample.UserFilter(c.filterFunc), sample.UserSort(c.sortFunc))
		if err != nil {
			t.Errorf("Error must not occur (case: %s, err=%v)", casename, err)
			continue
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
	}
}
