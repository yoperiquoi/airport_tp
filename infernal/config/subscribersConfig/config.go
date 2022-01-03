package subscribersConfig

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	BrokerHost string `json:"brokerHost"`
	BrokerPort int    `json:"BrokerPort"`
	Qos        int    `json:"qos"`
	ClientId   string `json:"clientId"`
}

func LoadConfig(subscriberType string) Config {
	var config Config
	configFile, err := os.Open("infernal/config/subscribersConfig/" + subscriberType + "_subscriber.json")
	if err != nil {
		log.Println(err.Error())
		return Config{}
	}
	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(configFile)

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		log.Println(err.Error())
		return Config{}
	}
	return config
}
