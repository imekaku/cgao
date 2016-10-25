package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type GolablConf struct {
	Dbaddr   string `json:"dbaddr"`
	Dbpasswd string `json:"dbpasswd"`
	Dbname   string `json:"dbname"`
}

var (
	config *GolablConf
)

func Config() *GolablConf {
	return config
}

func ParseConfig(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("os.Open err=", err)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("ioutil.ReadAll err=", err)
	}

	var cfg GolablConf
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
	}
	config = &cfg
}
