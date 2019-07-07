// Code generated by "sheetdb-modeler"; DO NOT EDIT.
// Create a Spreadsheet (sheet name: "typeTests") as data storage.
// The spreadsheet header is as follows:
//   id | string_value | bool_value | int_value | int_8_value | int_16_value | int_32_value | int_64_value | uint_value | uint_8_value | uint_16_value | uint_32_value | uint_64_value | float_32_value | float_64_value | date_value | datetime_value | p_bool_value | p_int_value | p_int_8_value | p_int_16_value | p_int_32_value | p_int_64_value | p_uint_value | p_uint_8_value | p_uint_16_value | p_uint_32_value | p_uint_64_value | p_float_32_value | p_float_64_value | p_date_value | p_datetime_value | updated_at | deleted_at
// Please copy and paste this header on the first line of the sheet.

package sample

import (
	"strconv"
	"sync"
	"time"

	"github.com/takuoki/gsheets"
	"github.com/takuoki/sheetdb"
)

const (
	// Sheet definition
	_TypeTest_sheetName             = "typeTests"
	_TypeTest_column_ID             = 0  // A
	_TypeTest_column_StringValue    = 1  // B
	_TypeTest_column_BoolValue      = 2  // C
	_TypeTest_column_IntValue       = 3  // D
	_TypeTest_column_Int8Value      = 4  // E
	_TypeTest_column_Int16Value     = 5  // F
	_TypeTest_column_Int32Value     = 6  // G
	_TypeTest_column_Int64Value     = 7  // H
	_TypeTest_column_UintValue      = 8  // I
	_TypeTest_column_Uint8Value     = 9  // J
	_TypeTest_column_Uint16Value    = 10 // K
	_TypeTest_column_Uint32Value    = 11 // L
	_TypeTest_column_Uint64Value    = 12 // M
	_TypeTest_column_Float32Value   = 13 // N
	_TypeTest_column_Float64Value   = 14 // O
	_TypeTest_column_DateValue      = 15 // P
	_TypeTest_column_DatetimeValue  = 16 // Q
	_TypeTest_column_PBoolValue     = 17 // R
	_TypeTest_column_PIntValue      = 18 // S
	_TypeTest_column_PInt8Value     = 19 // T
	_TypeTest_column_PInt16Value    = 20 // U
	_TypeTest_column_PInt32Value    = 21 // V
	_TypeTest_column_PInt64Value    = 22 // W
	_TypeTest_column_PUintValue     = 23 // X
	_TypeTest_column_PUint8Value    = 24 // Y
	_TypeTest_column_PUint16Value   = 25 // Z
	_TypeTest_column_PUint32Value   = 26 // AA
	_TypeTest_column_PUint64Value   = 27 // AB
	_TypeTest_column_PFloat32Value  = 28 // AC
	_TypeTest_column_PFloat64Value  = 29 // AD
	_TypeTest_column_PDateValue     = 30 // AE
	_TypeTest_column_PDatetimeValue = 31 // AF
	_TypeTest_column_UpdatedAt      = 32 // AG
	_TypeTest_column_DeletedAt      = 33 // AH
)

var (
	_TypeTest_mutex    = sync.RWMutex{}
	_TypeTest_cache    = map[int]*TypeTest{} // map[id]*TypeTest
	_TypeTest_rowNoMap = map[int]int{}       // map[id]rowNo
	_TypeTest_maxRowNo = 0
)

func init() {
	sheetdb.RegisterModel("default", "TypeTest", _TypeTest_sheetName, _TypeTest_load)
}

func _TypeTest_load(data *gsheets.Sheet) error {

	_TypeTest_mutex.Lock()
	defer _TypeTest_mutex.Unlock()

	_TypeTest_cache = map[int]*TypeTest{}
	_TypeTest_rowNoMap = map[int]int{}
	_TypeTest_maxRowNo = 0

	for i, r := range data.Rows() {
		if i == 0 {
			continue
		}
		if r.Value(_TypeTest_column_DeletedAt) != "" {
			_TypeTest_maxRowNo++
			continue
		}
		if r.Value(_TypeTest_column_ID) == "" {
			break
		}

		id, err := _TypeTest_parseID(r.Value(_TypeTest_column_ID))
		if err != nil {
			return err
		}
		stringValue := r.Value(_TypeTest_column_StringValue)
		if err := _TypeTest_validateStringValue(stringValue); err != nil {
			return err
		}
		boolValue, err := _TypeTest_parseBoolValue(r.Value(_TypeTest_column_BoolValue))
		if err != nil {
			return err
		}
		intValue, err := _TypeTest_parseIntValue(r.Value(_TypeTest_column_IntValue))
		if err != nil {
			return err
		}
		int8Value, err := _TypeTest_parseInt8Value(r.Value(_TypeTest_column_Int8Value))
		if err != nil {
			return err
		}
		int16Value, err := _TypeTest_parseInt16Value(r.Value(_TypeTest_column_Int16Value))
		if err != nil {
			return err
		}
		int32Value, err := _TypeTest_parseInt32Value(r.Value(_TypeTest_column_Int32Value))
		if err != nil {
			return err
		}
		int64Value, err := _TypeTest_parseInt64Value(r.Value(_TypeTest_column_Int64Value))
		if err != nil {
			return err
		}
		uintValue, err := _TypeTest_parseUintValue(r.Value(_TypeTest_column_UintValue))
		if err != nil {
			return err
		}
		uint8Value, err := _TypeTest_parseUint8Value(r.Value(_TypeTest_column_Uint8Value))
		if err != nil {
			return err
		}
		uint16Value, err := _TypeTest_parseUint16Value(r.Value(_TypeTest_column_Uint16Value))
		if err != nil {
			return err
		}
		uint32Value, err := _TypeTest_parseUint32Value(r.Value(_TypeTest_column_Uint32Value))
		if err != nil {
			return err
		}
		uint64Value, err := _TypeTest_parseUint64Value(r.Value(_TypeTest_column_Uint64Value))
		if err != nil {
			return err
		}
		float32Value, err := _TypeTest_parseFloat32Value(r.Value(_TypeTest_column_Float32Value))
		if err != nil {
			return err
		}
		float64Value, err := _TypeTest_parseFloat64Value(r.Value(_TypeTest_column_Float64Value))
		if err != nil {
			return err
		}
		dateValue, err := _TypeTest_parseDateValue(r.Value(_TypeTest_column_DateValue))
		if err != nil {
			return err
		}
		datetimeValue, err := _TypeTest_parseDatetimeValue(r.Value(_TypeTest_column_DatetimeValue))
		if err != nil {
			return err
		}
		pBoolValue, err := _TypeTest_parsePBoolValue(r.Value(_TypeTest_column_PBoolValue))
		if err != nil {
			return err
		}
		pIntValue, err := _TypeTest_parsePIntValue(r.Value(_TypeTest_column_PIntValue))
		if err != nil {
			return err
		}
		pInt8Value, err := _TypeTest_parsePInt8Value(r.Value(_TypeTest_column_PInt8Value))
		if err != nil {
			return err
		}
		pInt16Value, err := _TypeTest_parsePInt16Value(r.Value(_TypeTest_column_PInt16Value))
		if err != nil {
			return err
		}
		pInt32Value, err := _TypeTest_parsePInt32Value(r.Value(_TypeTest_column_PInt32Value))
		if err != nil {
			return err
		}
		pInt64Value, err := _TypeTest_parsePInt64Value(r.Value(_TypeTest_column_PInt64Value))
		if err != nil {
			return err
		}
		pUIntValue, err := _TypeTest_parsePUintValue(r.Value(_TypeTest_column_PUintValue))
		if err != nil {
			return err
		}
		pUInt8Value, err := _TypeTest_parsePUint8Value(r.Value(_TypeTest_column_PUint8Value))
		if err != nil {
			return err
		}
		pUInt16Value, err := _TypeTest_parsePUint16Value(r.Value(_TypeTest_column_PUint16Value))
		if err != nil {
			return err
		}
		pUInt32Value, err := _TypeTest_parsePUint32Value(r.Value(_TypeTest_column_PUint32Value))
		if err != nil {
			return err
		}
		pUInt64Value, err := _TypeTest_parsePUint64Value(r.Value(_TypeTest_column_PUint64Value))
		if err != nil {
			return err
		}
		pFloat32Value, err := _TypeTest_parsePFloat32Value(r.Value(_TypeTest_column_PFloat32Value))
		if err != nil {
			return err
		}
		pFloat64Value, err := _TypeTest_parsePFloat64Value(r.Value(_TypeTest_column_PFloat64Value))
		if err != nil {
			return err
		}
		pDateValue, err := _TypeTest_parsePDateValue(r.Value(_TypeTest_column_PDateValue))
		if err != nil {
			return err
		}
		pDatetimeValue, err := _TypeTest_parsePDatetimeValue(r.Value(_TypeTest_column_PDatetimeValue))
		if err != nil {
			return err
		}

		typeTest := TypeTest{
			ID:             id,
			StringValue:    stringValue,
			BoolValue:      boolValue,
			IntValue:       intValue,
			Int8Value:      int8Value,
			Int16Value:     int16Value,
			Int32Value:     int32Value,
			Int64Value:     int64Value,
			UintValue:      uintValue,
			Uint8Value:     uint8Value,
			Uint16Value:    uint16Value,
			Uint32Value:    uint32Value,
			Uint64Value:    uint64Value,
			Float32Value:   float32Value,
			Float64Value:   float64Value,
			DateValue:      dateValue,
			DatetimeValue:  datetimeValue,
			PBoolValue:     pBoolValue,
			PIntValue:      pIntValue,
			PInt8Value:     pInt8Value,
			PInt16Value:    pInt16Value,
			PInt32Value:    pInt32Value,
			PInt64Value:    pInt64Value,
			PUintValue:     pUIntValue,
			PUint8Value:    pUInt8Value,
			PUint16Value:   pUInt16Value,
			PUint32Value:   pUInt32Value,
			PUint64Value:   pUInt64Value,
			PFloat32Value:  pFloat32Value,
			PFloat64Value:  pFloat64Value,
			PDateValue:     pDateValue,
			PDatetimeValue: pDatetimeValue,
		}

		_TypeTest_maxRowNo++
		_TypeTest_cache[typeTest.ID] = &typeTest
		_TypeTest_rowNoMap[typeTest.ID] = _TypeTest_maxRowNo
	}

	return nil
}

// GetTypeTest returns a typeTest by ID.
// If it can not be found, this function returns sheetdb.NotFoundError.
func GetTypeTest(id int) (*TypeTest, error) {
	_TypeTest_mutex.RLock()
	defer _TypeTest_mutex.RUnlock()
	if v, ok := _TypeTest_cache[id]; ok {
		return v, nil
	}
	return nil, &sheetdb.NotFoundError{Model: "TypeTest"}
}

// TypeTestQuery is used for selecting typeTests.
type TypeTestQuery struct {
	filter func(typeTest *TypeTest) bool
	sort   func(typeTests []*TypeTest)
}

// TypeTestQueryOption is an option to change the behavior of TypeTestQuery.
type TypeTestQueryOption func(query *TypeTestQuery) *TypeTestQuery

// TypeTestFilter is an option to change the filtering behavior of TypeTestQuery.
func TypeTestFilter(filterFunc func(typeTest *TypeTest) bool) func(query *TypeTestQuery) *TypeTestQuery {
	return func(query *TypeTestQuery) *TypeTestQuery {
		if query != nil {
			query.filter = filterFunc
		}
		return query
	}
}

// TypeTestSort is an option to change the sorting behavior of TypeTestQuery.
func TypeTestSort(sortFunc func(typeTests []*TypeTest)) func(query *TypeTestQuery) *TypeTestQuery {
	return func(query *TypeTestQuery) *TypeTestQuery {
		if query != nil {
			query.sort = sortFunc
		}
		return query
	}
}

// GetTypeTests returns all typeTests.
// If any options are specified, the result according to the specified option is returned.
// If there are no typeTest to return, this function returns an nil array.
func GetTypeTests(opts ...TypeTestQueryOption) ([]*TypeTest, error) {
	typeTestQuery := &TypeTestQuery{}
	for _, opt := range opts {
		typeTestQuery = opt(typeTestQuery)
	}
	_TypeTest_mutex.RLock()
	defer _TypeTest_mutex.RUnlock()
	var typeTests []*TypeTest
	if typeTestQuery.filter != nil {
		for _, v := range _TypeTest_cache {
			if typeTestQuery.filter(v) {
				typeTests = append(typeTests, v)
			}
		}
	} else {
		for _, v := range _TypeTest_cache {
			typeTests = append(typeTests, v)
		}
	}
	if typeTestQuery.sort != nil {
		typeTestQuery.sort(typeTests)
	}
	return typeTests, nil
}

// AddTypeTest adds new typeTest.
// ID is generated automatically.
// If any fields are invalid, this function returns error.
func AddTypeTest(stringValue string, boolValue bool, intValue int, int8Value int8, int16Value int16, int32Value int32, int64Value int64, uintValue uint, uint8Value uint8, uint16Value uint16, uint32Value uint32, uint64Value uint64, float32Value float32, float64Value float64, dateValue sheetdb.Date, datetimeValue sheetdb.Datetime, pBoolValue *bool, pIntValue *int, pInt8Value *int8, pInt16Value *int16, pInt32Value *int32, pInt64Value *int64, pUIntValue *uint, pUInt8Value *uint8, pUInt16Value *uint16, pUInt32Value *uint32, pUInt64Value *uint64, pFloat32Value *float32, pFloat64Value *float64, pDateValue *sheetdb.Date, pDatetimeValue *sheetdb.Datetime) (*TypeTest, error) {
	_TypeTest_mutex.Lock()
	defer _TypeTest_mutex.Unlock()
	if err := _TypeTest_validateStringValue(stringValue); err != nil {
		return nil, err
	}
	typeTest := &TypeTest{
		ID:             _TypeTest_maxRowNo + 1,
		StringValue:    stringValue,
		BoolValue:      boolValue,
		IntValue:       intValue,
		Int8Value:      int8Value,
		Int16Value:     int16Value,
		Int32Value:     int32Value,
		Int64Value:     int64Value,
		UintValue:      uintValue,
		Uint8Value:     uint8Value,
		Uint16Value:    uint16Value,
		Uint32Value:    uint32Value,
		Uint64Value:    uint64Value,
		Float32Value:   float32Value,
		Float64Value:   float64Value,
		DateValue:      dateValue,
		DatetimeValue:  datetimeValue,
		PBoolValue:     pBoolValue,
		PIntValue:      pIntValue,
		PInt8Value:     pInt8Value,
		PInt16Value:    pInt16Value,
		PInt32Value:    pInt32Value,
		PInt64Value:    pInt64Value,
		PUintValue:     pUIntValue,
		PUint8Value:    pUInt8Value,
		PUint16Value:   pUInt16Value,
		PUint32Value:   pUInt32Value,
		PUint64Value:   pUInt64Value,
		PFloat32Value:  pFloat32Value,
		PFloat64Value:  pFloat64Value,
		PDateValue:     pDateValue,
		PDatetimeValue: pDatetimeValue,
	}
	if err := typeTest._asyncAdd(_TypeTest_maxRowNo + 1); err != nil {
		return nil, err
	}
	_TypeTest_maxRowNo++
	_TypeTest_cache[typeTest.ID] = typeTest
	_TypeTest_rowNoMap[typeTest.ID] = _TypeTest_maxRowNo
	return typeTest, nil
}

// UpdateTypeTest updates typeTest.
// If it can not be found, this function returns sheetdb.NotFoundError.
// If any fields are invalid, this function returns error.
func UpdateTypeTest(id int, stringValue string, boolValue bool, intValue int, int8Value int8, int16Value int16, int32Value int32, int64Value int64, uintValue uint, uint8Value uint8, uint16Value uint16, uint32Value uint32, uint64Value uint64, float32Value float32, float64Value float64, dateValue sheetdb.Date, datetimeValue sheetdb.Datetime, pBoolValue *bool, pIntValue *int, pInt8Value *int8, pInt16Value *int16, pInt32Value *int32, pInt64Value *int64, pUIntValue *uint, pUInt8Value *uint8, pUInt16Value *uint16, pUInt32Value *uint32, pUInt64Value *uint64, pFloat32Value *float32, pFloat64Value *float64, pDateValue *sheetdb.Date, pDatetimeValue *sheetdb.Datetime) (*TypeTest, error) {
	_TypeTest_mutex.Lock()
	defer _TypeTest_mutex.Unlock()
	typeTest, ok := _TypeTest_cache[id]
	if !ok {
		return nil, &sheetdb.NotFoundError{Model: "TypeTest"}
	}
	if err := _TypeTest_validateStringValue(stringValue); err != nil {
		return nil, err
	}
	typeTestCopy := *typeTest
	typeTestCopy.StringValue = stringValue
	typeTestCopy.BoolValue = boolValue
	typeTestCopy.IntValue = intValue
	typeTestCopy.Int8Value = int8Value
	typeTestCopy.Int16Value = int16Value
	typeTestCopy.Int32Value = int32Value
	typeTestCopy.Int64Value = int64Value
	typeTestCopy.UintValue = uintValue
	typeTestCopy.Uint8Value = uint8Value
	typeTestCopy.Uint16Value = uint16Value
	typeTestCopy.Uint32Value = uint32Value
	typeTestCopy.Uint64Value = uint64Value
	typeTestCopy.Float32Value = float32Value
	typeTestCopy.Float64Value = float64Value
	typeTestCopy.DateValue = dateValue
	typeTestCopy.DatetimeValue = datetimeValue
	typeTestCopy.PBoolValue = pBoolValue
	typeTestCopy.PIntValue = pIntValue
	typeTestCopy.PInt8Value = pInt8Value
	typeTestCopy.PInt16Value = pInt16Value
	typeTestCopy.PInt32Value = pInt32Value
	typeTestCopy.PInt64Value = pInt64Value
	typeTestCopy.PUintValue = pUIntValue
	typeTestCopy.PUint8Value = pUInt8Value
	typeTestCopy.PUint16Value = pUInt16Value
	typeTestCopy.PUint32Value = pUInt32Value
	typeTestCopy.PUint64Value = pUInt64Value
	typeTestCopy.PFloat32Value = pFloat32Value
	typeTestCopy.PFloat64Value = pFloat64Value
	typeTestCopy.PDateValue = pDateValue
	typeTestCopy.PDatetimeValue = pDatetimeValue
	if err := (&typeTestCopy)._asyncUpdate(); err != nil {
		return nil, err
	}
	typeTest = &typeTestCopy
	return typeTest, nil
}

// DeleteTypeTest deletes typeTest.
// If it can not be found, this function returns sheetdb.NotFoundError.
func DeleteTypeTest(id int) error {
	_TypeTest_mutex.Lock()
	defer _TypeTest_mutex.Unlock()
	typeTest, ok := _TypeTest_cache[id]
	if !ok {
		return &sheetdb.NotFoundError{Model: "TypeTest"}
	}
	if err := typeTest._asyncDelete(); err != nil {
		return err
	}
	delete(_TypeTest_cache, id)
	return nil
}

func _TypeTest_validateStringValue(stringValue string) error {
	if stringValue == "" {
		return &sheetdb.EmptyStringError{FieldName: "StringValue"}
	}
	return nil
}

func _TypeTest_parseID(id string) (int, error) {
	v, err := strconv.Atoi(id)
	if err != nil {
		return 0, &sheetdb.InvalidValueError{FieldName: "ID", Err: err}
	}
	return v, nil
}

func _TypeTest_parseBoolValue(boolValue string) (bool, error) {
	v, err := strconv.ParseBool(boolValue)
	if err != nil {
		return false, &sheetdb.InvalidValueError{FieldName: "BoolValue", Err: err}
	}
	return v, nil
}

func _TypeTest_parseIntValue(intValue string) (int, error) {
	v, err := strconv.Atoi(intValue)
	if err != nil {
		return 0, &sheetdb.InvalidValueError{FieldName: "IntValue", Err: err}
	}
	return v, nil
}

func _TypeTest_parseInt8Value(int8Value string) (int8, error) {
	v, err := strconv.ParseInt(int8Value, 10, 8)
	if err != nil {
		return 0, &sheetdb.InvalidValueError{FieldName: "Int8Value", Err: err}
	}
	return int8(v), nil
}

func _TypeTest_parseInt16Value(int16Value string) (int16, error) {
	v, err := strconv.ParseInt(int16Value, 10, 16)
	if err != nil {
		return 0, &sheetdb.InvalidValueError{FieldName: "Int16Value", Err: err}
	}
	return int16(v), nil
}

func _TypeTest_parseInt32Value(int32Value string) (int32, error) {
	v, err := strconv.ParseInt(int32Value, 10, 32)
	if err != nil {
		return 0, &sheetdb.InvalidValueError{FieldName: "Int32Value", Err: err}
	}
	return int32(v), nil
}

func _TypeTest_parseInt64Value(int64Value string) (int64, error) {
	v, err := strconv.ParseInt(int64Value, 10, 64)
	if err != nil {
		return 0, &sheetdb.InvalidValueError{FieldName: "Int64Value", Err: err}
	}
	return v, nil
}

func _TypeTest_parseUintValue(uintValue string) (uint, error) {
	v, err := strconv.ParseUint(uintValue, 10, 64)
	if err != nil {
		return 0, &sheetdb.InvalidValueError{FieldName: "UintValue", Err: err}
	}
	return uint(v), nil
}

func _TypeTest_parseUint8Value(uint8Value string) (uint8, error) {
	v, err := strconv.ParseUint(uint8Value, 10, 8)
	if err != nil {
		return 0, &sheetdb.InvalidValueError{FieldName: "Uint8Value", Err: err}
	}
	return uint8(v), nil
}

func _TypeTest_parseUint16Value(uint16Value string) (uint16, error) {
	v, err := strconv.ParseUint(uint16Value, 10, 16)
	if err != nil {
		return 0, &sheetdb.InvalidValueError{FieldName: "Uint16Value", Err: err}
	}
	return uint16(v), nil
}

func _TypeTest_parseUint32Value(uint32Value string) (uint32, error) {
	v, err := strconv.ParseUint(uint32Value, 10, 32)
	if err != nil {
		return 0, &sheetdb.InvalidValueError{FieldName: "Uint32Value", Err: err}
	}
	return uint32(v), nil
}

func _TypeTest_parseUint64Value(uint64Value string) (uint64, error) {
	v, err := strconv.ParseUint(uint64Value, 10, 64)
	if err != nil {
		return 0, &sheetdb.InvalidValueError{FieldName: "Uint64Value", Err: err}
	}
	return v, nil
}

func _TypeTest_parseFloat32Value(float32Value string) (float32, error) {
	v, err := strconv.ParseFloat(float32Value, 32)
	if err != nil {
		return 0, &sheetdb.InvalidValueError{FieldName: "Float32Value", Err: err}
	}
	return float32(v), nil
}

func _TypeTest_parseFloat64Value(float64Value string) (float64, error) {
	v, err := strconv.ParseFloat(float64Value, 64)
	if err != nil {
		return 0, &sheetdb.InvalidValueError{FieldName: "Float64Value", Err: err}
	}
	return float64(v), nil
}

func _TypeTest_parseDateValue(dateValue string) (sheetdb.Date, error) {
	v, err := sheetdb.NewDate(dateValue)
	if err != nil {
		return sheetdb.Date{}, &sheetdb.InvalidValueError{FieldName: "DateValue", Err: err}
	}
	return v, nil
}

func _TypeTest_parseDatetimeValue(datetimeValue string) (sheetdb.Datetime, error) {
	v, err := sheetdb.NewDatetime(datetimeValue)
	if err != nil {
		return sheetdb.Datetime{}, &sheetdb.InvalidValueError{FieldName: "DatetimeValue", Err: err}
	}
	return v, nil
}

func _TypeTest_parsePBoolValue(pBoolValue string) (*bool, error) {
	var val *bool
	if pBoolValue != "" {
		v, err := strconv.ParseBool(pBoolValue)
		if err != nil {
			return nil, &sheetdb.InvalidValueError{FieldName: "PBoolValue", Err: err}
		}
		val = &v
	}
	return val, nil
}

func _TypeTest_parsePIntValue(pIntValue string) (*int, error) {
	var val *int
	if pIntValue != "" {
		v, err := strconv.Atoi(pIntValue)
		if err != nil {
			return nil, &sheetdb.InvalidValueError{FieldName: "PIntValue", Err: err}
		}
		val = &v
	}
	return val, nil
}

func _TypeTest_parsePInt8Value(pInt8Value string) (*int8, error) {
	var val *int8
	if pInt8Value != "" {
		v, err := strconv.ParseInt(pInt8Value, 10, 8)
		if err != nil {
			return nil, &sheetdb.InvalidValueError{FieldName: "PInt8Value", Err: err}
		}
		v2 := int8(v)
		val = &v2
	}
	return val, nil
}

func _TypeTest_parsePInt16Value(pInt16Value string) (*int16, error) {
	var val *int16
	if pInt16Value != "" {
		v, err := strconv.ParseInt(pInt16Value, 10, 16)
		if err != nil {
			return nil, &sheetdb.InvalidValueError{FieldName: "PInt16Value", Err: err}
		}
		v2 := int16(v)
		val = &v2
	}
	return val, nil
}

func _TypeTest_parsePInt32Value(pInt32Value string) (*int32, error) {
	var val *int32
	if pInt32Value != "" {
		v, err := strconv.ParseInt(pInt32Value, 10, 32)
		if err != nil {
			return nil, &sheetdb.InvalidValueError{FieldName: "PInt32Value", Err: err}
		}
		v2 := int32(v)
		val = &v2
	}
	return val, nil
}

func _TypeTest_parsePInt64Value(pInt64Value string) (*int64, error) {
	var val *int64
	if pInt64Value != "" {
		v, err := strconv.ParseInt(pInt64Value, 10, 64)
		if err != nil {
			return nil, &sheetdb.InvalidValueError{FieldName: "PInt64Value", Err: err}
		}
		val = &v
	}
	return val, nil
}

func _TypeTest_parsePUintValue(pUIntValue string) (*uint, error) {
	var val *uint
	if pUIntValue != "" {
		v, err := strconv.ParseUint(pUIntValue, 10, 64)
		if err != nil {
			return nil, &sheetdb.InvalidValueError{FieldName: "PUintValue", Err: err}
		}
		v2 := uint(v)
		val = &v2
	}
	return val, nil
}

func _TypeTest_parsePUint8Value(pUInt8Value string) (*uint8, error) {
	var val *uint8
	if pUInt8Value != "" {
		v, err := strconv.ParseUint(pUInt8Value, 10, 8)
		if err != nil {
			return nil, &sheetdb.InvalidValueError{FieldName: "PUint8Value", Err: err}
		}
		v2 := uint8(v)
		val = &v2
	}
	return val, nil
}

func _TypeTest_parsePUint16Value(pUInt16Value string) (*uint16, error) {
	var val *uint16
	if pUInt16Value != "" {
		v, err := strconv.ParseUint(pUInt16Value, 10, 16)
		if err != nil {
			return nil, &sheetdb.InvalidValueError{FieldName: "PUint16Value", Err: err}
		}
		v2 := uint16(v)
		val = &v2
	}
	return val, nil
}

func _TypeTest_parsePUint32Value(pUInt32Value string) (*uint32, error) {
	var val *uint32
	if pUInt32Value != "" {
		v, err := strconv.ParseUint(pUInt32Value, 10, 32)
		if err != nil {
			return nil, &sheetdb.InvalidValueError{FieldName: "PUint32Value", Err: err}
		}
		v2 := uint32(v)
		val = &v2
	}
	return val, nil
}

func _TypeTest_parsePUint64Value(pUInt64Value string) (*uint64, error) {
	var val *uint64
	if pUInt64Value != "" {
		v, err := strconv.ParseUint(pUInt64Value, 10, 64)
		if err != nil {
			return nil, &sheetdb.InvalidValueError{FieldName: "PUint64Value", Err: err}
		}
		val = &v
	}
	return val, nil
}

func _TypeTest_parsePFloat32Value(pFloat32Value string) (*float32, error) {
	var val *float32
	if pFloat32Value != "" {
		v, err := strconv.ParseFloat(pFloat32Value, 32)
		if err != nil {
			return nil, &sheetdb.InvalidValueError{FieldName: "PFloat32Value", Err: err}
		}
		v2 := float32(v)
		val = &v2
	}
	return val, nil
}

func _TypeTest_parsePFloat64Value(pFloat64Value string) (*float64, error) {
	var val *float64
	if pFloat64Value != "" {
		v, err := strconv.ParseFloat(pFloat64Value, 64)
		if err != nil {
			return nil, &sheetdb.InvalidValueError{FieldName: "PFloat64Value", Err: err}
		}
		v2 := float64(v)
		val = &v2
	}
	return val, nil
}

func _TypeTest_parsePDateValue(pDateValue string) (*sheetdb.Date, error) {
	var val *sheetdb.Date
	if pDateValue != "" {
		v, err := sheetdb.NewDate(pDateValue)
		if err != nil {
			return nil, &sheetdb.InvalidValueError{FieldName: "PDateValue", Err: err}
		}
		val = &v
	}
	return val, nil
}

func _TypeTest_parsePDatetimeValue(pDatetimeValue string) (*sheetdb.Datetime, error) {
	var val *sheetdb.Datetime
	if pDatetimeValue != "" {
		v, err := sheetdb.NewDatetime(pDatetimeValue)
		if err != nil {
			return nil, &sheetdb.InvalidValueError{FieldName: "PDatetimeValue", Err: err}
		}
		val = &v
	}
	return val, nil
}

func (m *TypeTest) _asyncAdd(rowNo int) error {
	data := []gsheets.UpdateValue{
		{
			SheetName: _TypeTest_sheetName,
			RowNo:     rowNo,
			Values: []interface{}{
				m.ID,
				m.StringValue,
				m.BoolValue,
				m.IntValue,
				m.Int8Value,
				m.Int16Value,
				m.Int32Value,
				m.Int64Value,
				m.UintValue,
				m.Uint8Value,
				m.Uint16Value,
				m.Uint32Value,
				m.Uint64Value,
				m.Float32Value,
				m.Float64Value,
				m.DateValue.String(),
				m.DatetimeValue.String(),
				m.PBoolValue,
				m.PIntValue,
				m.PInt8Value,
				m.PInt16Value,
				m.PInt32Value,
				m.PInt64Value,
				m.PUintValue,
				m.PUint8Value,
				m.PUint16Value,
				m.PUint32Value,
				m.PUint64Value,
				m.PFloat32Value,
				m.PFloat64Value,
				m.PDateValue.String(),
				m.PDatetimeValue.String(),
				time.Now(),
				"",
			},
		},
	}
	return dbClient.AsyncUpdate(data)
}

func (m *TypeTest) _asyncUpdate() error {
	data := []gsheets.UpdateValue{
		{
			SheetName: _TypeTest_sheetName,
			RowNo:     _TypeTest_rowNoMap[m.ID],
			Values: []interface{}{
				m.ID,
				m.StringValue,
				m.BoolValue,
				m.IntValue,
				m.Int8Value,
				m.Int16Value,
				m.Int32Value,
				m.Int64Value,
				m.UintValue,
				m.Uint8Value,
				m.Uint16Value,
				m.Uint32Value,
				m.Uint64Value,
				m.Float32Value,
				m.Float64Value,
				m.DateValue.String(),
				m.DatetimeValue.String(),
				m.PBoolValue,
				m.PIntValue,
				m.PInt8Value,
				m.PInt16Value,
				m.PInt32Value,
				m.PInt64Value,
				m.PUintValue,
				m.PUint8Value,
				m.PUint16Value,
				m.PUint32Value,
				m.PUint64Value,
				m.PFloat32Value,
				m.PFloat64Value,
				m.PDateValue.String(),
				m.PDatetimeValue.String(),
				time.Now(),
				"",
			},
		},
	}
	return dbClient.AsyncUpdate(data)
}

func (m *TypeTest) _asyncDelete() error {
	now := time.Now()
	data := []gsheets.UpdateValue{
		{
			SheetName: _TypeTest_sheetName,
			RowNo:     _TypeTest_rowNoMap[m.ID],
			Values: []interface{}{
				m.ID,
				m.StringValue,
				m.BoolValue,
				m.IntValue,
				m.Int8Value,
				m.Int16Value,
				m.Int32Value,
				m.Int64Value,
				m.UintValue,
				m.Uint8Value,
				m.Uint16Value,
				m.Uint32Value,
				m.Uint64Value,
				m.Float32Value,
				m.Float64Value,
				m.DateValue.String(),
				m.DatetimeValue.String(),
				m.PBoolValue,
				m.PIntValue,
				m.PInt8Value,
				m.PInt16Value,
				m.PInt32Value,
				m.PInt64Value,
				m.PUintValue,
				m.PUint8Value,
				m.PUint16Value,
				m.PUint32Value,
				m.PUint64Value,
				m.PFloat32Value,
				m.PFloat64Value,
				m.PDateValue.String(),
				m.PDatetimeValue.String(),
				now,
				now,
			},
		},
	}
	return dbClient.AsyncUpdate(data)
}
