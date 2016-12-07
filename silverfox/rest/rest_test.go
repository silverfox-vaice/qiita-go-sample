package rest

import (
	"reflect"
	"testing"
)

func TestNewAccessControlHeaders(t *testing.T) {
	testCases := []struct {
		method string
		value  string
		result map[string]string
	}{
		{ // 正常系: AllowOrigin
			"AllowOrigin",
			"*",
			map[string]string{allowOrigin: "*"},
		},
		{ // 正常系: AllowHeaders
			"AllowHeaders",
			"Content-Type",
			map[string]string{allowHeaders: "Content-Type"},
		},
		{ // 正常系: AllowMethods
			"AllowMethods",
			"GET",
			map[string]string{allowMethods: "GET"},
		},
		{ // 正常系: AllowAllowExposeHeaders
			"AllowExposeHeaders",
			"Location",
			map[string]string{allowExposeHeaders: "Location"},
		},
		{ // 正常系: AllowMethodsAll
			"AllowMethodsAll",
			"#EMPTY#",
			map[string]string{allowMethods: "GET,POST,PUT,DELETE"},
		},
	}

	for i, tc := range testCases {
		caseNum := i + 1
		headers := NewAccessControlHeaders()
		method := reflect.ValueOf(&headers).MethodByName(tc.method)
		if tc.value == "#EMPTY#" {
			method.Call([]reflect.Value{})
		} else {
			method.Call([]reflect.Value{reflect.ValueOf(tc.value)})
		}
		result := headers.Get()
		if !reflect.DeepEqual(result, tc.result) {
			t.Errorf("case:%d faild. method = %s\n Expected = %#v\nResult = %#v", caseNum, tc.method, tc.result, result)
		}
	}
}
