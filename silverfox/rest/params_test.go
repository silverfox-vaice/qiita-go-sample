package rest

import (
	"reflect"
	"testing"
)

func TestRestNewUrlParam(t *testing.T) {

	testCases := []struct {
		url    string
		index  int
		result *UrlParam
		err    bool
	}{
		{ // 正常系: kindとKeyのみ
			"/api/stock/list",
			2,
			&UrlParam{
				Kind:       "stock",
				Keys:       []string{"list"},
				Params:     map[string]string{"list": ""},
				Conditions: map[string]string{},
			},
			false,
		},
		{ // 正常系: paramsが1つ
			"/api/stock/key/value",
			2,
			&UrlParam{
				Kind:       "stock",
				Keys:       []string{"key"},
				Params:     map[string]string{"key": "value"},
				Conditions: map[string]string{},
			},
			false,
		},
		{ // 正常系: paramsが複数
			"/api/stock/key/value/key2/value2",
			2,
			&UrlParam{
				Kind:       "stock",
				Keys:       []string{"key", "key2"},
				Params:     map[string]string{"key": "value", "key2": "value2"},
				Conditions: map[string]string{},
			},
			false,
		},
		{ // 正常系: kindのみ
			"/api/stock/",
			2,
			&UrlParam{Kind: "stock",
				Keys:       []string{},
				Params:     map[string]string{},
				Conditions: map[string]string{},
			},
			false,
		},
		{ // 異常系: パラメタ数がレンジ外
			"/api/",
			2,
			nil,
			true,
		},
	}

	for i, tc := range testCases {
		caseNum := i + 1
		result, err := NewUrlParam(tc.url, tc.index)
		if (err == nil && tc.err) || (err != nil && !tc.err) {
			t.Errorf("case:%d invalid error. %#v", caseNum, err)
		}
		if !reflect.DeepEqual(result, tc.result) {
			t.Errorf("case:%d faild. url = %s\n Expected = %#v\n Result   = %#v", caseNum, tc.url, tc.result, result)
		}
	}
}
