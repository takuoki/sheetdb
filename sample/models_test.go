package sample_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/takuoki/sheetdb"
	"github.com/takuoki/sheetdb/sample"
)

var (
	date19590525, _           = sheetdb.NewDate("1959-05-25")
	date19950914, _           = sheetdb.NewDate("1995-09-14")
	datetime20190707000000, _ = sheetdb.NewDatetime("2019-07-07T00:00:00.000+09:00")
)

func _() {
	// Run tests only for code generated in test mode.
	_ = sample.TestMode_User
	_ = sample.TestMode_Foo
	_ = sample.TestMode_FooChild
	_ = sample.TestMode_Bar
	_ = sample.TestMode_TypeTest
}

func TestMain(m *testing.M) {

	sheetdb.SetLogLevel(sheetdb.LevelNoLogging)

	// https://docs.google.com/spreadsheets/d/1dIxSIUM1vqehzt7gRz6Qi-UUlxeyl657ma88bfySs3E
	spreadsheetID := "1dIxSIUM1vqehzt7gRz6Qi-UUlxeyl657ma88bfySs3E"

	if err := sample.Initialize(context.Background(),
		os.Getenv("GOOGLE_API_CREDENTIALS"),
		os.Getenv("GOOGLE_API_TOKEN"),
		spreadsheetID); err != nil {
		fmt.Printf("Unable to initialize: %v\n", err)
		return
	}
	os.Exit(m.Run())
}
