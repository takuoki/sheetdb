package sheetdb_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/takuoki/sheetdb"
)

const (
	tStr    string = "2019-07-07T00:00:00+01:00"
	utcStr  string = "2019-07-07T00:00:00Z"
	fullfmt string = "2006-01-02T15:04:05.000000000Z07:00"
)

func TestNewDate(t *testing.T) {

	cases := map[string]struct {
		dateStr string
		isErr   bool
	}{
		"success":                   {dateStr: "2001-02-03"},
		"failure-empty":             {dateStr: "", isErr: true},
		"failure-invalid-delimitor": {dateStr: "2001/02/03", isErr: true},
		"failure-invalid-digit":     {dateStr: "2001-2-3", isErr: true},
	}

	for casename, c := range cases {
		t.Run(casename, func(t *testing.T) {
			d1, err := sheetdb.NewDate(c.dateStr)
			if !c.isErr {
				if err != nil {
					t.Fatalf("Error must not occur (case: %s, err: %v)", casename, err)
				}
				d2, _ := sheetdb.NewDate(c.dateStr)
				if d1 != d2 {
					t.Fatalf("Values generated from the same string should be equal (case: %s)", casename)
				}
				if d1.String() != c.dateStr {
					t.Fatalf("The string returned by String method does not match the string used at generation (case: %s, actual=%s)", casename, d1.String())
				}
			} else {
				if err == nil {
					t.Fatalf("Error must occur (case: %s)", casename)
				}
			}
		})
	}
}

func TestNewDatetime(t *testing.T) {

	_, offset := time.Now().Local().Zone()
	var local string
	if offset == 0 {
		local = "Z"
	} else if offset%3600 == 0 {
		if offset > 0 {
			local = fmt.Sprintf("+%02d:00", offset/3600)
		} else {
			local = fmt.Sprintf("%02d:00", offset*-1/3600)
		}
	} else {
		t.Fatalf("Unable to decode the local offset (offset=%d)", offset)
	}

	cases := map[string]struct {
		datetimeStr string
		isErr       bool
	}{
		"success-UTC":               {datetimeStr: "2001-02-03T01:02:03Z"},
		"success-local":             {datetimeStr: fmt.Sprintf("2001-02-03T01:02:03%s", local)},
		"success-other":             {datetimeStr: "2001-02-03T01:02:03+07:00"},
		"failure-empty":             {datetimeStr: "", isErr: true},
		"failure-invalid-delimitor": {datetimeStr: "2001/02/03T01:02:03Z", isErr: true},
		"failure-invalid-digit":     {datetimeStr: "2001-2-3T01:02:03Z", isErr: true},
	}

	for casename, c := range cases {
		t.Run(casename, func(t *testing.T) {
			d1, err := sheetdb.NewDatetime(c.datetimeStr)
			if !c.isErr {
				if err != nil {
					t.Fatalf("Error must not occur (case: %s, err: %v)", casename, err)
				}
				d2, _ := sheetdb.NewDatetime(c.datetimeStr)
				if d1 != d2 {
					t.Fatalf("Values generated from the same string should be equal (case: %s)", casename)
				}
				parsedD1, _ := time.Parse(time.RFC3339, c.datetimeStr)
				if d1.String() != parsedD1.UTC().Format(time.RFC3339) {
					t.Fatalf("The string returned by the String method does not match the string converted the string used at generation to UTC (case: %s, actual=%s)", casename, d1.String())
				}
			} else {
				if err == nil {
					t.Fatalf("Error must occur (case: %s)", casename)
				}
			}
		})
	}
}
