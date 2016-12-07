package service

type StockEntity struct {
	Code         string  `json:"code"`
	Name         string  `json:"name"`
	Market       string  `json:"market"`
	OpeningPrice float64 `json:"openingPrice"`
	Highprice    float64 `json:"highprice"`
	LowPrice     float64 `json:"lowPrice"`
	ClosingPrice float64 `json:"closingPrice"`
	Volume       int64   `json:"volume"`
	TradingValue int64   `json:"tradingValue"`
}
