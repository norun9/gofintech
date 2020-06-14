package models

type DataFrameCandle struct {
	ProductCode string   `json:"product_code"`
	Duration    string   `json:"duration"`
	Candles     []Candle `json:"candles"`
}
