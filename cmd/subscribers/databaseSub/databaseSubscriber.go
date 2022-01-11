package main

import (
	subconfig "airport_tp/internal/config/subscribersConfig"
	"airport_tp/internal/database"
	"airport_tp/internal/utils"
	subutils "airport_tp/internal/utils/subscribersUtils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"strconv"
	"strings"
	"time"
	_ "time"
)

type parsedMessage struct {
	captorId    string
	airportId   string
	measureType string
	value       string
	date        time.Time
}

var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	parsedString := strings.Split(string(msg.Payload()), "|")
	unixDate, _ := strconv.ParseInt(parsedString[4], 10, 64)
	date := time.Unix(unixDate, 0)
	parsedMessage := parsedMessage{
		captorId:    parsedString[0],
		airportId:   parsedString[1],
		measureType: parsedString[2],
		value:       parsedString[3],
		date:        date,
	}

	dbConnect := database.CreateConnexion()

	sensorTs := "sensor:" + parsedMessage.captorId + ":" + parsedMessage.airportId

	log.Println(sensorTs)
	
	value, _ := strconv.ParseFloat(parsedMessage.value, 64)
	_, err := dbConnect.Add(sensorTs, parsedMessage.date.Unix(), value)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	config := subconfig.LoadConfig("database")
	log.Println("Load configuration of the subscriber")
	subscriber := subutils.Connect("tcp://"+config.BrokerHost+":"+strconv.Itoa(config.BrokerPort), config.ClientId)

	tokenSubscriberTemp := subscriber.Subscribe(utils.TopicTemp, byte(config.Qos), messageHandler)
	tokenSubscriberPressure := subscriber.Subscribe(utils.TopicPressure, byte(config.Qos), messageHandler)
	tokenSubscriberWind := subscriber.Subscribe(utils.TopicWind, byte(config.Qos), messageHandler)


	for {
		tokenSubscriberTemp.Wait()
		tokenSubscriberPressure.Wait()
		tokenSubscriberWind.Wait()
	}



}