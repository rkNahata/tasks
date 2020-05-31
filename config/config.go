package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configuration struct {
	ServerPort string
}

var config Configuration
var err error

func ReadConfigFile(fileName string) Configuration {
	conf, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(conf, &config)
	if err != nil {
		log.Fatalln(err)
	}
	return config
}
