package sample

import (
	"context"
	"testing"

	"github.com/takuoki/gsheets"
	"github.com/takuoki/sheetdb"
)

var (
	loaded                       bool
	loadedUserCache              = map[int]*User{} // map[userID]*User
	loadedUserRowNoMap           = map[int]int{}   // map[userID]rowNo
	loadedUserMaxRowNo           = 0
	loadedUserEmailUniqueMap     = map[string]*User{}
	loadedFooCache               = map[int]map[int]*Foo{} // map[userID][fooID]*Foo
	loadedFooRowNoMap            = map[int]map[int]int{}  // map[userID][fooID]rowNo
	loadedFooMaxRowNo            = 0
	loadedFooChildCache          = map[int]map[int]map[int]*FooChild{} // map[userID][fooID][childID]*FooChild
	loadedFooChildRowNoMap       = map[int]map[int]map[int]int{}       // map[userID][fooID][childID]rowNo
	loadedFooChildMaxRowNo       = 0
	loadedFooChildValueUniqueMap = map[string]*FooChild{}
	loadedBarCache               = map[int]map[sheetdb.Datetime]*Bar{} // map[userID][datetime]*Bar
	loadedBarRowNoMap            = map[int]map[sheetdb.Datetime]int{}  // map[userID][datetime]rowNo
	loadedBarMaxRowNo            = 0
	loadedTypeTestCache          = map[int]*TypeTest{} // map[id]*TypeTest
	loadedTypeTestRowNoMap       = map[int]int{}       // map[id]rowNo
	loadedTypeTestMaxRowNo       = 0
)

func LoadUser(t *testing.T, data [][]interface{}) error {
	t.Helper()
	// add header row
	data = append([][]interface{}{{}}, data...)
	return _User_load(gsheets.NewSheet(t, data))
}

func Reload(t *testing.T) {
	t.Helper()
	if !loaded {
		if err := LoadData(context.Background()); err != nil {
			t.Fatalf("Unable to load data from spreadsheet: %v", err)
		}
		loadedUserCache = copyUserCache(_User_cache)
		loadedUserRowNoMap = copyRowNoMap(_User_rowNoMap)
		loadedUserMaxRowNo = _User_maxRowNo
		loadedUserEmailUniqueMap = copyUserEmailMap(_User_Email_uniqueMap)
		loadedFooCache = copyFooCache(_Foo_cache)
		loadedFooRowNoMap = copyFooRowNoMap(_Foo_rowNoMap)
		loadedFooMaxRowNo = _Foo_maxRowNo
		loadedFooChildCache = copyFooChildCache(_FooChild_cache)
		loadedFooChildRowNoMap = copyFooChildRowNoMap(_FooChild_rowNoMap)
		loadedFooChildMaxRowNo = _FooChild_maxRowNo
		loadedFooChildValueUniqueMap = copyFooChildValueMap(_FooChild_Value_uniqueMap)
		loadedBarCache = copyBarCache(_Bar_cache)
		loadedBarRowNoMap = copyBarRowNoMap(_Bar_rowNoMap)
		loadedBarMaxRowNo = _Bar_maxRowNo
		loadedTypeTestCache = copyTypeTestCache(_TypeTest_cache)
		loadedTypeTestRowNoMap = copyRowNoMap(_TypeTest_rowNoMap)
		loadedTypeTestMaxRowNo = _TypeTest_maxRowNo
		loaded = true
	} else {
		_User_cache = copyUserCache(loadedUserCache)
		_User_rowNoMap = copyRowNoMap(loadedUserRowNoMap)
		_User_maxRowNo = loadedUserMaxRowNo
		_User_Email_uniqueMap = copyUserEmailMap(loadedUserEmailUniqueMap)
		_Foo_cache = copyFooCache(loadedFooCache)
		_Foo_rowNoMap = copyFooRowNoMap(loadedFooRowNoMap)
		_Foo_maxRowNo = loadedFooMaxRowNo
		_FooChild_cache = copyFooChildCache(loadedFooChildCache)
		_FooChild_rowNoMap = copyFooChildRowNoMap(loadedFooChildRowNoMap)
		_FooChild_maxRowNo = loadedFooChildMaxRowNo
		_FooChild_Value_uniqueMap = copyFooChildValueMap(loadedFooChildValueUniqueMap)
		_Bar_cache = copyBarCache(loadedBarCache)
		_Bar_rowNoMap = copyBarRowNoMap(loadedBarRowNoMap)
		_Bar_maxRowNo = loadedBarMaxRowNo
		_TypeTest_cache = copyTypeTestCache(loadedTypeTestCache)
		_TypeTest_rowNoMap = copyRowNoMap(loadedTypeTestRowNoMap)
		_TypeTest_maxRowNo = loadedTypeTestMaxRowNo
	}
}

func copyRowNoMap(m map[int]int) map[int]int {
	r := map[int]int{}
	for k, v := range m {
		r[k] = v
	}
	return r
}

func copyUserCache(m map[int]*User) map[int]*User {
	r := map[int]*User{}
	for k, v := range m {
		u := *v
		r[k] = &u
	}
	return r
}

func copyUserEmailMap(m map[string]*User) map[string]*User {
	r := map[string]*User{}
	for k, v := range m {
		u := *v
		r[k] = &u
	}
	return r
}

func copyFooCache(m map[int]map[int]*Foo) map[int]map[int]*Foo {
	r := map[int]map[int]*Foo{}
	for k, v := range m {
		r[k] = map[int]*Foo{}
		for k2, v2 := range v {
			f := *v2
			r[k][k2] = &f
		}
	}
	return r
}

func copyFooRowNoMap(m map[int]map[int]int) map[int]map[int]int {
	r := map[int]map[int]int{}
	for k, v := range m {
		r[k] = map[int]int{}
		for k2, v2 := range v {
			r[k][k2] = v2
		}
	}
	return r
}

func copyFooChildCache(m map[int]map[int]map[int]*FooChild) map[int]map[int]map[int]*FooChild {
	r := map[int]map[int]map[int]*FooChild{}
	for k, v := range m {
		r[k] = map[int]map[int]*FooChild{}
		for k2, v2 := range v {
			r[k][k2] = map[int]*FooChild{}
			for k3, v3 := range v2 {
				f := *v3
				r[k][k2][k3] = &f
			}
		}
	}
	return r
}

func copyFooChildRowNoMap(m map[int]map[int]map[int]int) map[int]map[int]map[int]int {
	r := map[int]map[int]map[int]int{}
	for k, v := range m {
		r[k] = map[int]map[int]int{}
		for k2, v2 := range v {
			r[k][k2] = map[int]int{}
			for k3, v3 := range v2 {
				r[k][k2][k3] = v3
			}
		}
	}
	return r
}

func copyFooChildValueMap(m map[string]*FooChild) map[string]*FooChild {
	r := map[string]*FooChild{}
	for k, v := range m {
		f := *v
		r[k] = &f
	}
	return r
}

func copyBarCache(m map[int]map[sheetdb.Datetime]*Bar) map[int]map[sheetdb.Datetime]*Bar {
	r := map[int]map[sheetdb.Datetime]*Bar{}
	for k, v := range m {
		r[k] = map[sheetdb.Datetime]*Bar{}
		for k2, v2 := range v {
			b := *v2
			r[k][k2] = &b
		}
	}
	return r
}

func copyBarRowNoMap(m map[int]map[sheetdb.Datetime]int) map[int]map[sheetdb.Datetime]int {
	r := map[int]map[sheetdb.Datetime]int{}
	for k, v := range m {
		r[k] = map[sheetdb.Datetime]int{}
		for k2, v2 := range v {
			r[k][k2] = v2
		}
	}
	return r
}

func copyTypeTestCache(m map[int]*TypeTest) map[int]*TypeTest {
	r := map[int]*TypeTest{}
	for k, v := range m {
		t := *v
		r[k] = &t
	}
	return r
}
