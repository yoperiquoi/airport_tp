package captorConfig

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	BrokerHost   string  `json:"brokerHost"`
	BrokerPort   int     `json:"brokerPort"`
	Qos          int     `json:"qos"`
	ClientId     string  `json:"clientId"`
	PublishDelai int     `json:"publishDelai"`
	CaptorId     int     `json:"captorId"`
	IataCode     string  `json:"IataCode"`
	MeasureType  string  `json:"measureType"`
	Max          float64 `json:"max"`
	Min          float64 `json:"min"`
	Variation    float64 `json:"variation"`
}

func LoadConfig(publisherType string) Config {
	var config Config
	configFile, err := os.Open(publisherType + "_publisher.json")
	if err != nil {
		log.Println(err.Error())
		panic(err)
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
