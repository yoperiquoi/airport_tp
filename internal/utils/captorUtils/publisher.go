package captorUtils

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"strconv"
	"time"
)

func createClientOptions(brokerURI string, clientId string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURI)
	opts.SetClientID(clientId)
	return opts
}

func Connect(brokerURI string, clientId string) mqtt.Client {
	log.Println("Trying to connect to broker : " + brokerURI + ", with publisher : " + clientId + "")
	opts := createClientOptions(brokerURI, clientId)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}

func FormatMessage(captorId int, airportID string, measureType string, value float64, date time.Time) string {
	return strconv.Itoa(captorId) + "|" + airportID + "|" + measureType + "|" + fmt.Sprintf("%.2f", value) + "|" + strconv.Itoa(int(date.Unix()))
}
