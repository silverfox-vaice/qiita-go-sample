package service

import (
	"reflect"
	"silverfox/rest"
	"testing"

	"google.golang.org/appengine/aetest"
)

func TestJsonToEntity(t *testing.T) {
	testCases := []struct {
		input  rest.JsonParam
		result *StockEntity
	}{
		{
			rest.JsonParam{
				Body: map[string]interface{}{
					"code":         "Code1",
					"name":         "Name1",
					"market":       "Market1",
					"openingPrice": 100.5,
					"highprice":    "100.5",
					"lowPrice":     100.5,
					"closingPrice": 100.5,
					"volume":       999999999,
					"tradingValue": "999999999",
				},
			},
			&StockEntity{Code: "Code1",
				Name:         "Name1",
				Market:       "Market1",
				OpeningPrice: 100.5,
				Highprice:    100.5,
				LowPrice:     100.5,
				ClosingPrice: 100.5,
				Volume:       999999999,
				TradingValue: 999999999,
			},
		},
	}

	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()
	setUp(t, ctx)

	for i, tc := range testCases {
		caseNum := i + 1
		result, err := JsonToEntity(tc.input)
		if err != nil {
			t.Errorf("case:%d invalid error. %#v", caseNum, err)
		}
		if !reflect.DeepEqual(result, tc.result) {
			t.Errorf("case:%d faild.\n Expected = %#v\n Result   = %#v", caseNum, tc.result, result)
		}
	}
}
func TestCsvToEntity(t *testing.T) {
	//dxo := NewStockEntityDxo()
	testCases := []struct {
		input  []string
		result *StockEntity
		errNum int
	}{
		{
			[]string{"Code1", "Name1", "Market1", "100.5", "100.5", "100.5", "100.5", "999999999", "999999999"},
			&StockEntity{Code: "Code1",
				Name:         "Name1",
				Market:       "Market1",
				OpeningPrice: 100.5,
				Highprice:    100.5,
				LowPrice:     100.5,
				ClosingPrice: 100.5,
				Volume:       999999999,
				TradingValue: 999999999,
			},
			0,
		},
		{
			[]string{"Code1", "Name1", "Market1", "aaa", "100.5", "100.5", "100.5", "999999999", "999999999"},
			&StockEntity{Code: "Code1",
				Name:         "Name1",
				Market:       "Market1",
				OpeningPrice: 0,
				Highprice:    100.5,
				LowPrice:     100.5,
				ClosingPrice: 100.5,
				Volume:       999999999,
				TradingValue: 999999999,
			},
			1,
		},
	}

	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()
	setUp(t, ctx)

	for i, tc := range testCases {
		caseNum := i + 1
		result, err := CsvToEntity(tc.input)
		if len(err) != tc.errNum {
			t.Errorf("case:%d invalid error. %#v", caseNum, err)
		}
		if !reflect.DeepEqual(result, tc.result) {
			t.Errorf("case:%d faild.\n Expected = %#v\n Result   = %#v", caseNum, tc.result, result)
		}
	}
}
