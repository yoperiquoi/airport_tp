package captorConfig

import (
	"encoding/json"
	"log"
	"os"
)

/*
	Struct associated to JSON object into the config file of a captor Publisher
*/
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

/*
	Function which permits to read into the config file and associate it to the struct config
*/
func LoadConfig(publisherType string) Config {
	var config Config
	// Open the config file in the same directory of the executable program
	configFile, err := os.Open(publisherType + "_publisher.json")
	// If the file can't be found / open stop the program
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	// Closing the config file at the end of the function
	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(configFile)

	// Read the json object into the configFile
	jsonParser := json.NewDecoder(configFile)
	// Associate each field of the object to the struct
	err = jsonParser.Decode(&config)
	if err != nil {
		log.Println(err.Error())
		return Config{}
	}
	return config
}
