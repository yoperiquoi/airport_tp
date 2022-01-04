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

func generateRandomWind() float64 {
	min := utils.WindMin
	max := utils.WindMax
	return (rand.Float64() * (max - min)) + min
}

func main() {
	config := pubconfig.LoadConfig("wind")
	log.Println("Load configuration of the publisher")
	publisher := pubutils.Connect("tcp://"+config.BrokerHost+":"+strconv.Itoa(config.BrokerPort), config.ClientId)
	defer func() {
		publisher.Disconnect(250)
		log.Println(config.ClientId + " disconnect from the broker")
	}()

	for {
		message := pubutils.FormatMessage(config.CaptorId, config.IataCode, config.MeasureType, generateRandomWind(), time.Now())
		publisher.Publish(utils.TopicWind, byte(config.Qos), false, message)
		log.Println(config.ClientId + " publish : " + message)
		time.Sleep(time.Second * time.Duration(config.PublishDelai))
	}
}
