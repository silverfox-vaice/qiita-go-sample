package rest

import (
	"reflect"
	"testing"
)

func TestStrToInt64(t *testing.T) {

	testCases := []struct {
		input  string
		result int64
		err    bool
	}{
		{ // 正常系
			"5",
			5,
			false,
		},
		{ // 正常系:0
			"0",
			0,
			false,
		},
		{ // 正常系:空文字
			"",
			0,
			false,
		},
		{ // 正常系:マイナス
			"-5",
			-5,
			false,
		},
		{ // 正常系:int最大値
			"9223372036854775807",
			9223372036854775807,
			false,
		},
		{ // 正常系:int最小値
			"-9223372036854775808",
			-9223372036854775808,
			false,
		},
		{ // 異常系:最大値超え
			"9223372036854775808",
			9223372036854775807,
			true,
		},
		{ // 異常系:最小値超え
			"-9223372036854775809",
			-9223372036854775808,
			true,
		},
		{ // 異常系:少数点
			"5.5",
			0,
			true,
		},
		{ // 異常系:文字
			"aaa",
			0,
			true,
		},
	}

	for i, tc := range testCases {
		caseNum := i + 1
		dxo := Dxo{}
		result := dxo.StrToInt64(tc.input)
		if (len(dxo.Err) == 0 && tc.err) || (len(dxo.Err) != 0 && !tc.err) {
			t.Errorf("case:%d invalid error. %#v", caseNum, dxo.Err)
		}
		if !reflect.DeepEqual(result, tc.result) {
			t.Errorf("case:%d faild.\ninput = %s\n Expected = %#v\n Result   = %#v", caseNum, tc.input, tc.result, result)
		}
	}
}

func TestStrToFloat(t *testing.T) {

	testCases := []struct {
		input  string
		result float64
		err    bool
	}{
		{ // 正常系:整数値
			"5",
			5,
			false,
		},
		{ // 正常系:少数値
			"5.5",
			5.5,
			false,
		},
		{ // 正常系:少数値（ゼロ無し）
			".5",
			0.5,
			false,
		},
		{ // 正常系:0
			"0",
			0,
			false,
		},
		{ // 正常系:空文字
			"",
			0,
			false,
		},
		{ // 正常系:マイナス整数値
			"-5",
			-5,
			false,
		},
		{ // 正常系:マイナス
			"-5.5",
			-5.5,
			false,
		},
		{ // 異常系:文字
			"aaa",
			0,
			true,
		},
	}

	for i, tc := range testCases {
		caseNum := i + 1
		dxo := Dxo{}
		result := dxo.StrToFloat(tc.input)
		if (len(dxo.Err) == 0 && tc.err) || (len(dxo.Err) != 0 && !tc.err) {
			t.Errorf("case:%d invalid error. %#v", caseNum, dxo.Err)
		}
		if !reflect.DeepEqual(result, tc.result) {
			t.Errorf("case:%d faild.\ninput = %s\n Expected = %#v\n Result   = %#v", caseNum, tc.input, tc.result, result)
		}
	}
}

func TestToFloat(t *testing.T) {

	testCases := []struct {
		input  interface{}
		result float64
		err    bool
	}{
		{ // 正常系:文字型
			string("5.5"),
			5.5,
			false,
		},
		{ // 正常系:int型
			int(-5),
			-5,
			false,
		},
		{ // 正常系:uint型
			uint(100),
			100,
			false,
		},
		{ // 正常系:float型
			float64(.5),
			0.5,
			false,
		},
		{ // 正常系:nil
			nil,
			0,
			false,
		},
		{ // 異常系:文字
			"aaa",
			0,
			true,
		},
		{ // 異常系:未対応の型
			bool(true),
			0,
			true,
		},
	}

	for i, tc := range testCases {
		caseNum := i + 1
		dxo := Dxo{}
		result := dxo.ToFloat(tc.input)
		if (len(dxo.Err) == 0 && tc.err) || (len(dxo.Err) != 0 && !tc.err) {
			t.Errorf("case:%d invalid error. %#v", caseNum, dxo.Err)
		}
		if !reflect.DeepEqual(result, tc.result) {
			t.Errorf("case:%d faild.\ninput = %s\n Expected = %#v\n Result   = %#v", caseNum, tc.input, tc.result, result)
		}
	}
}

func TestToInt(t *testing.T) {

	testCases := []struct {
		input  interface{}
		result int64
		err    bool
	}{
		{ // 正常系:int文字型
			string("5"),
			5,
			false,
		},
		{ // 正常系:int型
			int(-5),
			-5,
			false,
		},
		{ // 正常系:uint型
			uint(100),
			100,
			false,
		},
		{ // 正常系:float型
			float64(.5),
			0,
			false,
		},
		{ // 正常系:nil
			nil,
			0,
			false,
		},
		{ // 正常系:空文字
			"",
			0,
			false,
		},
		{ // 異常系:文字
			"aaa",
			0,
			true,
		},
		{ // 異常系:未対応の型
			bool(true),
			0,
			true,
		},
		{ // 異常系:小数点文字型
			string("5.5"),
			0,
			true,
		},
	}

	for i, tc := range testCases {
		caseNum := i + 1
		dxo := Dxo{}
		result := dxo.ToInt(tc.input)
		if (len(dxo.Err) == 0 && tc.err) || (len(dxo.Err) != 0 && !tc.err) {
			t.Errorf("case:%d invalid error. %#v", caseNum, dxo.Err)
		}
		if !reflect.DeepEqual(result, tc.result) {
			t.Errorf("case:%d faild.\ninput = %s\n Expected = %#v\n Result   = %#v", caseNum, tc.input, tc.result, result)
		}
	}
}

func TestToString(t *testing.T) {

	testCases := []struct {
		input  interface{}
		result string
		err    bool
	}{
		{ // 正常系:文字型
			string("str"),
			"str",
			false,
		},
		{ // 正常系:int型
			int(-5),
			"-5",
			false,
		},
		{ // 正常系:uint型
			uint(100),
			"100",
			false,
		},
		{ // 正常系:float型
			float64(.5),
			"0.5000",
			false,
		},
		{ // 正常系:nil
			nil,
			"",
			false,
		},
		{ // 正常系:空文字
			"",
			"",
			false,
		},
		{ // 正常系:boolean
			bool(true),
			"true",
			false,
		},
		{ // 異常系:未対応の型
			[]string{"array1"},
			"",
			true,
		},
	}

	for i, tc := range testCases {
		caseNum := i + 1
		dxo := Dxo{}
		result := dxo.ToString(tc.input)
		if (len(dxo.Err) == 0 && tc.err) || (len(dxo.Err) != 0 && !tc.err) {
			t.Errorf("case:%d invalid error. %#v", caseNum, dxo.Err)
		}
		if !reflect.DeepEqual(result, tc.result) {
			t.Errorf("case:%d faild.\ninput = %s\n Expected = %#v\n Result   = %#v", caseNum, tc.input, tc.result, result)
		}
	}
}
