package common

import (
	"encoding/json"
	"log"
	"os"
)

type configuration struct {
	Port string
}

//AppConfig configuration for the redis
var AppConfig configuration

//InitConfig initialize the config
func InitConfig() {
	loadAppConfig()
}

func loadAppConfig() {
	file, err := os.Open("config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}

}
