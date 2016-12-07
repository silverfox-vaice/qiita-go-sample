package rest

import (
	"bytes"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestParseHTTPBody(t *testing.T) {

	testCases := []struct {
		contentType   string
		contentLength string
		jsonString    string
		result        map[string]interface{}
		err           bool
	}{
		{ // 正常系
			"application/json",
			"10",
			`{"id":55}`,
			map[string]interface{}{"id": 55.0},
			false,
		},
		{ // 正常系
			"application/json",
			"10000",
			`{"id":[100,101,102]}`,
			map[string]interface{}{"id": []interface{}{100.0, 101.0, 102.0}},
			false,
		},
		{ // 正常系(空)
			"application/json",
			"0",
			"",
			map[string]interface{}(nil),
			false,
		},
		{ // 異常系: ヘッダ不正
			"text/plain",
			"100",
			`{"id":55}`,
			nil,
			true,
		},
		{ // 異常系: lengthなし
			"application/json",
			"nil",
			`{"id":55}`,
			nil,
			true,
		},
		{ // 異常系: jsonフォーマットがおかしい
			"application/json",
			"100",
			`{"id":55`,
			nil,
			true,
		},
	}

	for i, tc := range testCases {
		caseNum := i + 1
		r := httptest.NewRequest("POST", "/", bytes.NewBuffer([]byte(tc.jsonString)))
		r.Header.Set("Content-Type", tc.contentType)
		r.Header.Set("Content-Length", tc.contentLength)
		param, err := ParseHTTPBody(r)
		if (err == nil && tc.err) || (err != nil && !tc.err) {
			t.Errorf("case:%d invalid error. %#v", caseNum, err)
		}
		result := param.Body
		if !reflect.DeepEqual(tc.result, result) {
			t.Errorf("case:%d faild. jsonString = %s\n Expected = %#v\n Result   = %#v", caseNum, tc.jsonString, tc.result, result)
		}
	}
}
