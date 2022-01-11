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

func generateRandomWind(lastWind float64, config pubconfig.Config) float64 {
	maxVariation := config.Variation
	if rand.Intn(2) == 0 {
		lastWind = lastWind + maxVariation
	} else {
		lastWind = lastWind - maxVariation
	}
	return lastWind
}

func main() {
	config := pubconfig.LoadConfig("wind")
	log.Println("Load configuration of the publisher")
	publisher := pubutils.Connect("tcp://"+config.BrokerHost+":"+strconv.Itoa(config.BrokerPort), config.ClientId)
	defer func() {
		publisher.Disconnect(250)
		log.Println(config.ClientId + " disconnect from the broker")
	}()
	lastWind := config.Max - ((config.Max - config.Min) / 2)
	for {
		lastWind = generateRandomWind(lastWind, config)
		message := pubutils.FormatMessage(config.CaptorId, config.IataCode, config.MeasureType, lastWind, time.Now())
		publisher.Publish(utils.TopicWind, byte(config.Qos), false, message)
		log.Println(config.ClientId + " publish : " + message)
		time.Sleep(time.Second * time.Duration(config.PublishDelai))
	}
}
