package service

import (
	"reflect"
	"testing"

	"golang.org/x/net/context"

	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
)

func setUp(t *testing.T, ctx context.Context) {
	entitys := []struct {
		entity *StockEntity
	}{
		{
			&StockEntity{
				Code:         "Code",
				Name:         "Name",
				Market:       "Market",
				OpeningPrice: 100,
				Highprice:    100,
				LowPrice:     100,
				ClosingPrice: 100,
				Volume:       100,
				TradingValue: 100,
			},
		},
		{
			&StockEntity{
				Code:         "Code2",
				Name:         "Name2",
				Market:       "Market2",
				OpeningPrice: 0,
				Highprice:    100.15,
				LowPrice:     100.15,
				ClosingPrice: 100.15,
				Volume:       100000000000,
				TradingValue: 100000000000,
			},
		},
	}

	for _, e := range entitys {
		key, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "Stock", nil), e.entity)
		if err != nil {
			t.Errorf("datastore error.\n%#v", err)
		}
		// 一度取得しないとコミットされないので取得する
		datastore.Get(ctx, key, &StockEntity{})
	}
}

func TestGetList(t *testing.T) {
	testCases := []struct {
		result *[]StockEntity
	}{
		{
			&[]StockEntity{
				StockEntity{
					Code:         "Code2",
					Name:         "Name2",
					Market:       "Market2",
					OpeningPrice: 0,
					Highprice:    100.15,
					LowPrice:     100.15,
					ClosingPrice: 100.15,
					Volume:       100000000000,
					TradingValue: 100000000000,
				},

				StockEntity{
					Code:         "Code",
					Name:         "Name",
					Market:       "Market",
					OpeningPrice: 100,
					Highprice:    100,
					LowPrice:     100,
					ClosingPrice: 100,
					Volume:       100,
					TradingValue: 100,
				},
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
		result, err := GetList(ctx)
		if err != nil {
			t.Errorf("case:%d invalid error. %#v", caseNum, err)
		}
		if !reflect.DeepEqual(result, tc.result) {
			t.Errorf("case:%d faild.\n Expected = %#v\n Result   = %#v", caseNum, tc.result, result)
		}
	}
}

func TestGetItem(t *testing.T) {
	testCases := []struct {
		code   string
		result *StockEntity
	}{
		{
			"Code",
			&StockEntity{
				Code:         "Code",
				Name:         "Name",
				Market:       "Market",
				OpeningPrice: 100,
				Highprice:    100,
				LowPrice:     100,
				ClosingPrice: 100,
				Volume:       100,
				TradingValue: 100,
			},
		},
		{
			"None",
			&StockEntity{
				Code:         "",
				Name:         "",
				Market:       "",
				OpeningPrice: 0,
				Highprice:    0,
				LowPrice:     0,
				ClosingPrice: 0,
				Volume:       0,
				TradingValue: 0,
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
		result, err := GetItem(ctx, tc.code)
		if err != nil {
			t.Errorf("case:%d invalid error. %#v", caseNum, err)
		}
		if !reflect.DeepEqual(result, tc.result) {
			t.Errorf("case:%d faild.\n Expected = %#v\n Result   = %#v", caseNum, tc.result, result)
		}
	}
}

func TestGetPut(t *testing.T) {
	testCases := []struct {
		code   string
		input  *StockEntity
		result *StockEntity
	}{
		{
			"Code",
			&StockEntity{
				Code:         "Code",
				Name:         "Name",
				Market:       "Market",
				OpeningPrice: 5555,
				Highprice:    100,
				LowPrice:     100,
				ClosingPrice: 100,
				Volume:       100,
				TradingValue: 100,
			},

			&StockEntity{
				Code:         "Code",
				Name:         "Name",
				Market:       "Market",
				OpeningPrice: 5555,
				Highprice:    100,
				LowPrice:     100,
				ClosingPrice: 100,
				Volume:       100,
				TradingValue: 100,
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
		result, err := FullUpdate(ctx, tc.code, tc.input)
		if err != nil {
			t.Errorf("case:%d invalid error. %#v", caseNum, err)
		}
		if !reflect.DeepEqual(result, tc.result) {
			t.Errorf("case:%d faild.\n Expected = %#v\n Result   = %#v", caseNum, tc.result, result)
		}
	}
}
