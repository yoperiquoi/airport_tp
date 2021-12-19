package captorConfig

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	BrokerHost string `json:"brokerHost"`
	BrokerPort string `json:"BrokerPort"`
	Qos        int    `json:"qos"`
	ClientId   int    `json:"clientId"`
}

func LoadConfig(publisherType string) Config {
	var config Config
	configFile, err := os.Open("infernal/config/captorConfig/" + publisherType + "_publisher.json")
	if err != nil {
		fmt.Println(err.Error())
		return Config{}
	}
	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(configFile)

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		fmt.Println(err.Error())
		return Config{}
	}
	return config
}
