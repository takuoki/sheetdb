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
	datetime19590525, _ = sheetdb.NewDate("1959-05-25")
)

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
