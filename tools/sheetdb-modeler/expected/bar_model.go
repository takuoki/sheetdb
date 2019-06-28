package sample

import (
	"sync"

	"github.com/takuoki/gsheets"
	"github.com/takuoki/sheetdb"
)

const (
	_Bar_sheetName = "bars"
)

var (
	_Bar_mutex    = sync.RWMutex{}
	_Bar_cache    = map[int]map[sheetdb.Datetime]*Bar{} // map[userID][datetime]*Bar
	_Bar_rowNoMap = map[int]map[sheetdb.Datetime]int{}  // map[userID][fooID]rowNo
	_Bar_maxRowNo int
)

func init() {
	dbClient.AddModel("bar", loadBar)
}

func loadBar(data *gsheets.Sheet) error {

	return nil
}
