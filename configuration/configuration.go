package configuration

import (
	"fmt"
	"os"
	"runtime"
	"path"
	"path/filepath"
	"strings"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	API struct {
		Key      string `json:"Key"`
		URL      string `json:"Url"`
		Datasets []struct {
			Name       string `json:"Name"`
			Comparison string `json:"Comparison"`
		} `json:"Datasets"`
		Query struct {
			StartDate string `json:"StartDate"`
			Sort      string `json:"Sort"`
		} `json:"Query"`
	} `json:"Api"`
}


func InitConfiguration() Configuration {
	configuration := Configuration{}
	err := gonfig.GetConf(getFileName(), &configuration)
	if err != nil {
		fmt.Println("error " + err.Error())
		os.Exit(1)
	}
	return configuration
}

func getFileName() string {
	filename := []string{"configuration", ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))

	return filePath
}