package service

import "silverfox/rest"

func CsvToEntity(row []string) (*StockEntity, []error) {
	dxo := rest.Dxo{}
	entity := &StockEntity{}
	entity.Code = row[0]
	entity.Name = row[1]
	entity.Market = row[2]
	entity.OpeningPrice = dxo.StrToFloat(row[3])
	entity.Highprice = dxo.StrToFloat(row[4])
	entity.LowPrice = dxo.StrToFloat(row[5])
	entity.ClosingPrice = dxo.StrToFloat(row[6])
	entity.Volume = dxo.StrToInt64(row[7])
	entity.TradingValue = dxo.StrToInt64(row[8])
	return entity, dxo.Err
}

func JsonToEntity(json rest.JsonParam) (*StockEntity, []error) {
	dxo := rest.Dxo{}
	entity := &StockEntity{}
	entity.Code = dxo.ToString(json.Body["code"])
	entity.Name = dxo.ToString(json.Body["name"])
	entity.Market = dxo.ToString(json.Body["market"])
	entity.OpeningPrice = dxo.ToFloat(json.Body["openingPrice"])
	entity.Highprice = dxo.ToFloat(json.Body["highprice"])
	entity.LowPrice = dxo.ToFloat(json.Body["lowPrice"])
	entity.ClosingPrice = dxo.ToFloat(json.Body["closingPrice"])
	entity.Volume = dxo.ToInt(json.Body["volume"])
	entity.TradingValue = dxo.ToInt(json.Body["tradingValue"])
	return entity, dxo.Err
}
