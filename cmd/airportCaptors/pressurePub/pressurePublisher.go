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
	Function which generates a value for the pressure to send to the broker
*/
func generateRandomPressure(lastPressure float64, config pubconfig.Config) float64 {
	// Based on the previous value it generate a value according to the variation
	maxVariation := config.Variation
	if rand.Intn(2) == 0 {
		lastPressure = lastPressure + maxVariation
	} else {
		lastPressure = lastPressure - maxVariation
	}
	return lastPressure
}

func main() {
	config := pubconfig.LoadConfig("pressure")
	log.Println("Load configuration of the publisher")
	// Connect to the broker
	publisher := pubutils.Connect("tcp://"+config.BrokerHost+":"+strconv.Itoa(config.BrokerPort), config.ClientId)
	// Disconnect of the broker at the end of the program
	defer func() {
		publisher.Disconnect(250)
		log.Println(config.ClientId + " disconnect from the broker")
	}()
	lastPressure := config.Max - ((config.Max - config.Min) / 2)
	for {
		lastPressure = generateRandomPressure(lastPressure, config)
		message := pubutils.FormatMessage(config.CaptorId, config.IataCode, config.MeasureType, lastPressure, time.Now())
		// Send a message to the broker with the new value
		publisher.Publish(utils.TopicPressure, byte(config.Qos), false, message)
		log.Println(config.ClientId + " publish : " + message)
		// Loop every time defined into the config
		time.Sleep(time.Second * time.Duration(config.PublishDelai))
	}
}
