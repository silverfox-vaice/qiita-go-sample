package stock

import (
	"context"
	"reflect"
	"silverfox/rest"
	"silverfox/sample1/stock/service"
	"testing"

	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
)

func setUp(t *testing.T, ctx context.Context) {
	entitys := []struct {
		entity *service.StockEntity
	}{
		{
			&service.StockEntity{
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
			&service.StockEntity{
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
		datastore.Get(ctx, key, &service.StockEntity{})
	}
}

func TestGet(t *testing.T) {
	testCases := []struct {
		input  *rest.UrlParam
		result interface{}
	}{
		{ // code指定:指定したcodeのレコードが取得されることを確認
			&rest.UrlParam{Kind: "stock",
				Keys:   []string{"code"},
				Params: map[string]string{"code": "Code"},
			},
			&service.StockEntity{
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
		{ // list指定:レコードが全件取得されることを確認
			&rest.UrlParam{Kind: "stock",
				Keys:   []string{"list"},
				Params: map[string]string{"list": ""},
			},
			&[]service.StockEntity{
				service.StockEntity{
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
				service.StockEntity{
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
		{ // 準正常系：kind指定のみ
			&rest.UrlParam{Kind: "stock",
				Keys:   []string{},
				Params: map[string]string{},
			},
			nil,
		},
		{ // 準正常系：未対応のキーを指定
			&rest.UrlParam{Kind: "stock",
				Keys:   []string{"none"},
				Params: map[string]string{"none": "Code"},
			},
			nil,
		},
	}

	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()
	setUp(t, ctx)

	s := StockRest{}
	s.SetContext(ctx)

	for i, tc := range testCases {
		caseNum := i + 1
		result := s.Get(tc.input)
		if err != nil {
			t.Errorf("case:%d invalid error. %#v", caseNum, err)
		}
		if !reflect.DeepEqual(result, tc.result) {
			t.Errorf("case:%d faild.\n Expected = %#v\n Result   = %#v", caseNum, tc.result, result)
		}
	}
}

func TestOptions(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	s := StockRest{}
	s.SetContext(ctx)

	// ArrrowHeaderが指定数分セットされていることを確認
	result := s.Options()
	expectNum := 2
	if len(result.Get()) != expectNum {
		t.Errorf("invalid headers count. Expected = %d, result = %d", expectNum, len(result.Get()))
	}
}

func TestPut(t *testing.T) {
	testCases := []struct {
		param  *rest.UrlParam
		json   rest.JsonParam
		result interface{}
	}{
		{ // 上書きリクエストをしても価格4種以外は無視されることを確認
			&rest.UrlParam{Kind: "stock",
				Keys:   []string{"code"},
				Params: map[string]string{"code": "Code"},
			},
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
			&service.StockEntity{
				Code:         "Code",
				Name:         "Name",
				Market:       "Market",
				OpeningPrice: 100.5,
				Highprice:    100.5,
				LowPrice:     100.5,
				ClosingPrice: 100.5,
				Volume:       100,
				TradingValue: 100,
			},
		},
		{
			&rest.UrlParam{Kind: "stock",
				Keys:   []string{},
				Params: map[string]string{},
			},
			rest.JsonParam{},
			nil,
		},
		{
			&rest.UrlParam{Kind: "stock",
				Keys:   []string{"none"},
				Params: map[string]string{"none": "Code"},
			},
			rest.JsonParam{},
			nil,
		},
	}

	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()
	setUp(t, ctx)

	s := StockRest{}
	s.SetContext(ctx)

	for i, tc := range testCases {
		caseNum := i + 1
		result := s.Put(tc.param, tc.json)
		if err != nil {
			t.Errorf("case:%d invalid error. %#v", caseNum, err)
		}
		if !reflect.DeepEqual(result, tc.result) {
			t.Errorf("case:%d faild.\n Expected = %#v\n Result   = %#v", caseNum, tc.result, result)
		}
	}
}

func TestDelete(t *testing.T) {
	testCases := []struct {
		input *rest.UrlParam
	}{
		{ // TODO:未実装
			&rest.UrlParam{Kind: "stock",
				Keys:   []string{"code"},
				Params: map[string]string{"code": "Code"},
			},
		},
	}

	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()
	setUp(t, ctx)

	s := StockRest{}
	s.SetContext(ctx)

	for _, tc := range testCases {
		s.Delete(tc.input)
	}
}

func TestPost(t *testing.T) {
	testCases := []struct {
		input *rest.UrlParam
	}{
		{ // TODO:未実装
			&rest.UrlParam{Kind: "stock",
				Keys:   []string{"code"},
				Params: map[string]string{"code": "Code"},
			},
		},
	}

	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()
	setUp(t, ctx)

	s := StockRest{}
	s.SetContext(ctx)

	for _, tc := range testCases {
		s.Post(tc.input)
	}
}
