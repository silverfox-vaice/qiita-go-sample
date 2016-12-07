package sample1

import (
	"io"
	"net/http/httptest"
	"reflect"
	"testing"

	"google.golang.org/appengine/aetest"
)

func TestRestHanderFactory(t *testing.T) {

	testCases := []struct {
		url      string
		respType string
	}{
		{
			"stock",
			"*stock.StockRest",
		},
		{
			"",
			"",
		},
		{
			"none",
			"",
		},
	}

	for _, tc := range testCases {
		resp, err := restHanderFactory(tc.url)
		result := ""
		if err == nil {
			result = reflect.TypeOf(resp).String()
		}
		if result != tc.respType {
			t.Errorf("invalid type. url = %s, Expected = %v, Result = %v", tc.url, result, tc.respType)
		}
	}
}

func TestApiHandler(t *testing.T) {
	testCases := []struct {
		method string
		url    string
		body   io.Reader
	}{
		{
			"GET",
			"/api/stock/list",
			nil,
		}, {
			"GET",
			"/api/stock/",
			nil,
		}, {
			"OPTIONS",
			"/api/stock/",
			nil,
		}, {
			"DELETE",
			"/api/stock/",
			nil,
		}, {
			"PUT",
			"/api/stock/",
			nil,
		}, {
			"POST",
			"/api/stock/",
			nil,
		}, {
			"UNKNOWN",
			"/api/none/",
			nil,
		}, {
			"UNKNOWN",
			"/none",
			nil,
		},
	}

	instance, _ := aetest.NewInstance(&aetest.Options{StronglyConsistentDatastore: true})
	defer instance.Close()

	for _, tc := range testCases {
		r, _ := instance.NewRequest(tc.method, tc.url, tc.body)
		w := httptest.NewRecorder()
		apiHandler(w, r)
	}
}
