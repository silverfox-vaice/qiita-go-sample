package stock

import (
	"silverfox/rest"
	"silverfox/sample1/stock/service"

	"golang.org/x/net/context"
)

const Stock = "stock"

type StockRest struct {
	Ctx context.Context
}

func (s *StockRest) SetContext(context context.Context) {
	s.Ctx = context
}

func (s *StockRest) Options() rest.AccessControlHeaders {
	headers := rest.NewAccessControlHeaders()
	// sampleなので*で許可
	headers.AllowHeaders("Content-Type, Content-Length")
	headers.AllowMethodsAll()
	return headers

}

func (s *StockRest) Get(param *rest.UrlParam) interface{} {
	if len(param.Keys) == 0 {
		return nil
	}
	if param.Keys[0] == "list" {
		entitys, _ := service.GetList(s.Ctx)
		return interface{}(entitys)
	} else if param.Keys[0] == "code" {
		entitys, _ := service.GetItem(s.Ctx, param.GetParam(0))
		return interface{}(entitys)
	}
	return nil
}

func (s *StockRest) Post(param *rest.UrlParam) interface{} {
	// 未実装
	return nil
}

func (s *StockRest) Put(param *rest.UrlParam, json rest.JsonParam) interface{} {
	if len(param.Keys) == 0 {
		return nil
	}
	if param.Keys[0] == "code" {
		entity, _ := service.JsonToEntity(json)
		entitys, _ := service.DiffUpdate(s.Ctx, param.GetParam(0), entity)
		return interface{}(entitys)
	}
	return nil
}
func (s *StockRest) Delete(param *rest.UrlParam) interface{} {
	// 未実装
	return nil
}
