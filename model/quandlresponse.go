package model

type DataResponse struct {
	Date string `json:"date"`
	Value float64 `json:"value"`
}

type QuandlResponse struct {
	DataSet struct {
		Data [][]interface{} `json:"data"`
	} `json:"dataset"`
}