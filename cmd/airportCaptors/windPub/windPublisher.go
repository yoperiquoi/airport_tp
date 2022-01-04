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

func generateRandomWind(lastWind float64) float64 {
	maxVariation := utils.WindMaxVariation
	if rand.Intn(2) == 0 {
		lastWind = lastWind + maxVariation
	}else{
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
	lastWind := utils.WindMax - ((utils.WindMax - utils.WindMin)/2)
	for {
		lastWind = generateRandomWind(lastWind)
		message := pubutils.FormatMessage(config.CaptorId, config.IataCode, config.MeasureType, lastWind, time.Now())
		publisher.Publish(utils.TopicWind, byte(config.Qos), false, message)
		log.Println(config.ClientId + " publish : " + message)
		time.Sleep(time.Second * time.Duration(config.PublishDelai))
	}
}
