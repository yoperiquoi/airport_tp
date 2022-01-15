package subscribersUtils

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
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
	log.Println("Trying to connect to broker : " + brokerURI + ", with subscriber : " + clientId + "")
	opts := createClientOptions(brokerURI, clientId)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}
