package quandl

import (
	"github.com/andream16/all-hail-gbp/configuration"
	"github.com/andream16/all-hail-gbp/model"
	"github.com/google/go-querystring/query"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/go-errors/errors"
)

// Gets response for GBPvsEur and GBPvsUSD in []{ string : "yy-mm-dd", float64 : "0.00" }
func CrawlCurrencies(config *configuration.Configuration) ([][]model.DataResponse, error) {
	request := model.Request{ ApiKey: config.API.Key, StartDate: config.API.Query.StartDate, Order: config.API.Query.Sort }
	queryString := func (r  model.Request) string {
		v, _  := query.Values(r)
		return v.Encode()
	}(request)
	var url = config.API.URL
	var finalResponse = make([][]model.DataResponse, 2)
	for _, k := range config.API.Datasets {
		currQuery := url + k.Name + "/" + k.Comparison + ".json?" + queryString
		var currResponse model.QuandlResponse
		var cassandraFormatResponse []model.DataResponse
		response, err := http.Get(currQuery); if err != nil {
			fmt.Println(err.Error())
		}
		defer response.Body.Close()
		decodeErr := json.NewDecoder(response.Body).Decode(&currResponse); if decodeErr != nil {
			fmt.Println(decodeErr.Error())
		}
		for _, v := range currResponse.DataSet.Data {
			cassandraFormatResponse = append(cassandraFormatResponse, model.DataResponse{Date: v[0].(string), Value: v[1].(float64)})
		}
		if len(cassandraFormatResponse) > 0 {
			finalResponse = append(finalResponse, cassandraFormatResponse)
		}
	}
	if len(finalResponse) > 0 {
		return finalResponse, nil
	}
	return [][]model.DataResponse{}, errors.New("unable to parse responses")
}
