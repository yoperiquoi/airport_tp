package captorUtils

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"strconv"
	"time"
)

/*
	Function which returns the options for the MQTT connection to the broker
*/
func createClientOptions(brokerURI string, clientId string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURI)
	opts.SetClientID(clientId)
	return opts
}

/*
	Function which returns the connection to the MQTT Broker
*/
func Connect(brokerURI string, clientId string) mqtt.Client {
	log.Println("Trying to connect to broker : " + brokerURI + ", with publisher : " + clientId + "")
	opts := createClientOptions(brokerURI, clientId)
	client := mqtt.NewClient(opts)
	// Connect to the broker and exit the program if there are errors
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}

/*
	Function to set the format of message sent through MQTT protocol
*/
func FormatMessage(captorId int, airportID string, measureType string, value float64, date time.Time) string {
	return strconv.Itoa(captorId) + "|" + airportID + "|" + measureType + "|" + fmt.Sprintf("%.2f", value) + "|" + strconv.Itoa(int(date.Unix()))
}
