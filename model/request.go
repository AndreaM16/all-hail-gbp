package model

type Request struct {
	ApiKey string `url:"api_key"`
	StartDate string `url:"start_date"`
	Order string `url:"order"`
}
