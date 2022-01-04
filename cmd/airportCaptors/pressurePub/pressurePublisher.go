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

func generateRandomPressure(lastPressure float64) float64 {
	maxVariation := utils.PressureMaxVariation
	if rand.Intn(2) == 0 {
		lastPressure = lastPressure + maxVariation
	}else{
		lastPressure = lastPressure - maxVariation
	}
	return lastPressure
}

func main() {
	config := pubconfig.LoadConfig("pressure")
	log.Println("Load configuration of the publisher")
	publisher := pubutils.Connect("tcp://"+config.BrokerHost+":"+strconv.Itoa(config.BrokerPort), config.ClientId)
	defer func() {
		publisher.Disconnect(250)
		log.Println(config.ClientId + " disconnect from the broker")
	}()
	lastPressure := utils.PressureMax - ((utils.PressureMax - utils.PressureMin)/2)
	for {
		lastPressure = generateRandomPressure(lastPressure)
		message := pubutils.FormatMessage(config.CaptorId, config.IataCode, config.MeasureType, lastPressure, time.Now())
		publisher.Publish(utils.TopicPressure, byte(config.Qos), false, message)
		log.Println(config.ClientId + " publish : " + message)
		time.Sleep(time.Second * time.Duration(config.PublishDelai))
	}
}
