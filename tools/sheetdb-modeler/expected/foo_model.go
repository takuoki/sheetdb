package sample

import (
	"sync"

	"github.com/takuoki/gsheets"
)

const (
	_Foo_sheetName = "foos"
)

var (
	_Foo_mutex    = sync.RWMutex{}
	_Foo_cache    = map[int]map[int]*Foo{} // map[userID][fooID]*Foo
	_Foo_rowNoMap = map[int]map[int]int{}  // map[userID][fooID]rowNo
	_Foo_maxRowNo int
)

func init() {
	dbClient.AddModel("foo", loadFoo)
}

func loadFoo(data *gsheets.Sheet) error {

	return nil
}
