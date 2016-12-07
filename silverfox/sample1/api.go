package sample1

import (
	"errors"
	"log"
	"net/http"
	"silverfox/rest"
	"silverfox/sample1/stock"

	"google.golang.org/appengine"
)

// API用のハンドラ
func apiHandler(w http.ResponseWriter, r *http.Request) {

	param, err := rest.NewUrlParam(r.URL.Path, 2)
	if err != nil {
		log.Printf("%s is invalid url.\n%#v", r.URL.Path, err)
		return
	}

	restHandler, err := restHanderFactory(param.Kind)
	if err != nil {
		log.Printf("%#v", err)
		return
	}

	ctx := appengine.NewContext(r)
	restHandler.SetContext(ctx)

	resp := rest.Response{}

	json, _ := rest.ParseHTTPBody(r)

	switch r.Method {
	case rest.Get:
		resp.Entry = restHandler.Get(param)
	case rest.Put:
		resp.Entry = restHandler.Put(param, json)
	case rest.Post:
		resp.Entry = restHandler.Post(param)
	case rest.Delete:
		resp.Entry = restHandler.Delete(param)
	case rest.Options:
		headers := restHandler.Options()
		for k, v := range headers.Get() {
			w.Header().Add(k, v)
		}
	default:
		log.Printf("%#v", r.Method)
	}

	rest.Respond(w, 200, resp)
	return
}

// handlerの種類が増えていく予定
func restHanderFactory(kind string) (rest.RestHandler, error) {
	switch kind {
	case stock.Stock:
		return &stock.StockRest{}, nil
	default:
		return nil, errors.New("この種別には対応していません。")
	}
}
