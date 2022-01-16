package main

import (
	pubconfig "airport_tp/internal/config/captorConfig"
	"airport_tp/internal/utils"
	pubutils "airport_tp/internal/utils/captorUtils"
	"log"
	"math/rand"
	"strconv"
	"time"
)

/*
	Function which generates a value for the temperature to send to the broker
*/
func generateRandomTemperature(lastTemperature float64, config pubconfig.Config) float64 {
	// Based on the previous value it generate a value according to the variation
	maxVariation := config.Variation
	if rand.Intn(4) == 0 {
		lastTemperature = lastTemperature + maxVariation
	} else {
		lastTemperature = lastTemperature - maxVariation
	}
	return lastTemperature
}

func main() {
	config := pubconfig.LoadConfig("temperature")
	log.Println("Load configuration of the publisher")
	// Connect to the broker
	publisher := pubutils.Connect("tcp://"+config.BrokerHost+":"+strconv.Itoa(config.BrokerPort), config.ClientId)
	// Disconnect of the broker at the end of the program
	defer func() {
		publisher.Disconnect(250)
		log.Println(config.ClientId + " disconnect from the broker")
	}()
	lastTemperature := config.Max - ((config.Max - config.Min) / 2)
	for {
		lastTemperature = generateRandomTemperature(lastTemperature, config)
		message := pubutils.FormatMessage(config.CaptorId, config.IataCode, config.MeasureType, lastTemperature, time.Now())
		// Send a message to the broker with the new value
		publisher.Publish(utils.TopicTemp, byte(config.Qos), false, message)
		log.Println(config.ClientId + " publish : " + message)
		// Loop every time defined into the config
		time.Sleep(time.Second * time.Duration(config.PublishDelai))
	}
}
