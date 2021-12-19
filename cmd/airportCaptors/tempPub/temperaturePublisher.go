package main

import (
	pubconfig "airport_tp/infernal/config/captorConfig"
	pubutils "airport_tp/infernal/utils/captorUtils"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func generateRandomTemperature() float64 {
	min := pubutils.TempMin
	max := pubutils.TempMax
	return (rand.Float64() * (max - min)) + min
}

func main() {
	config := pubconfig.LoadConfig("temperature")
	log.Println("Load configuration of the publisher")
	publisher := pubutils.Connect("tcp://"+config.BrokerHost+":"+strconv.Itoa(config.BrokerPort), config.ClientId)
	defer func() {
		publisher.Disconnect(250)
		log.Println(config.ClientId + " disconnect from the broker")
	}()

	for {
		message := pubutils.FormatMessage(config.CaptorId, config.IataCode, config.MeasureType, generateRandomTemperature(), time.Now())
		publisher.Publish(pubutils.TopicTemp, byte(config.Qos), false, message)
		log.Println(config.ClientId + " publish : " + message)
		time.Sleep(time.Second * time.Duration(config.PublishDelai))
	}
}
