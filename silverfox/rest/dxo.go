package rest

import (
	"errors"
	"reflect"
	"strconv"
)

type Dxo struct {
	Err []error
}

// 文字列をint64に変換する
func (e *Dxo) StrToInt64(str string) int64 {
	if str == "" {
		return 0
	}
	intVal, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		e.Err = append(e.Err, err)
	}
	return intVal
}

func (e *Dxo) StrToFloat(str string) float64 {
	if str == "" {
		return 0
	}
	floatVal, err := strconv.ParseFloat(str, 64)
	if err != nil {
		e.Err = append(e.Err, err)
	}
	return floatVal
}

func (e *Dxo) ToFloat(v interface{}) float64 {
	if v == nil || v == "" {
		return 0
	}
	r := reflect.ValueOf(v)
	switch r.Kind() {
	case reflect.String:
		float, err := strconv.ParseFloat(r.String(), 64)
		if err != nil {
			e.Err = append(e.Err, err)
			return 0
		}
		return float
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(r.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(r.Uint())
	case reflect.Float32, reflect.Float64:
		return r.Float()
	default:
		e.Err = append(e.Err, errors.New("invalid type."))
		return 0
	}
}

func (e *Dxo) ToInt(v interface{}) int64 {
	if v == nil || v == "" {
		return 0
	}
	r := reflect.ValueOf(v)
	switch r.Kind() {
	case reflect.String:
		intVal, err := strconv.ParseInt(r.String(), 10, 64)
		if err != nil {
			e.Err = append(e.Err, err)
			return 0
		}
		return intVal
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int64(r.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int64(r.Uint())
	case reflect.Float32, reflect.Float64:
		return int64(r.Float())
	default:
		e.Err = append(e.Err, errors.New("invalid type."))
		return 0
	}
}
func (e *Dxo) ToString(v interface{}) string {
	if v == nil || v == "" {
		return ""
	}
	r := reflect.ValueOf(v)
	switch r.Kind() {
	case reflect.String:
		return r.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(r.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(r.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(r.Float(), 'f', 4, 64)
	case reflect.Bool:
		return strconv.FormatBool(r.Bool())
	default:
		e.Err = append(e.Err, errors.New("invalid type."))
		return ""
	}
}
