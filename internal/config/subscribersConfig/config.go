package subscribersConfig

import (
	"encoding/json"
	"log"
	"os"
)

/*
	Struct associated to JSON object into the config file of a Subscriber which
	stores data into database or csv files
*/
type Config struct {
	BrokerHost string `json:"brokerHost"`
	BrokerPort int    `json:"BrokerPort"`
	Qos        int    `json:"qos"`
	ClientId   string `json:"clientId"`
}

/*
	Function which permits to read into the config file and associate it to the struct config
*/
func LoadConfig(subscriberType string) Config {
	var config Config
	// Open the config file in the same directory of the executable program
	configFile, err := os.Open(subscriberType + "_subscriber.json")
	// If the file can't be found / open stop the program
	if err != nil {
		log.Println(err.Error())
		return Config{}
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
