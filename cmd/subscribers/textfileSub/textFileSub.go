package main

import (
	subconfig "airport_tp/infernal/config/subscribersConfig"
	"airport_tp/infernal/utils"
	subutils "airport_tp/infernal/utils/subscribersUtils"
	"encoding/csv"
	"errors"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	_ "time"
)

func fileExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

type parsedMessage struct {
	captorId    string
	airportId   string
	measureType string
	value       string
	date        time.Time
}

func insertIntoCsv(csvFile *os.File, message parsedMessage) {
	csvWriter := csv.NewWriter(csvFile)
	line := []string{
		message.captorId,
		message.airportId,
		message.measureType,
		message.value,
		message.date.Format("2006-01-02-15-04-05"),
	}
	log.Println("added row : " + fmt.Sprint(line) + " to file : " + csvFile.Name())
	err := csvWriter.Write(line)
	if err != nil {
		fmt.Println(err)
	}
	csvWriter.Flush()
}

var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	parsedString := strings.Split(string(msg.Payload()), "|")

	a, _ := time.Parse("2006-01-02-15-04-05", parsedString[4])

	parsedMessage := parsedMessage{
		captorId:    parsedString[0],
		airportId:   parsedString[1],
		measureType: parsedString[2],
		value:       parsedString[3],
		date:        a,
	}

	filename := parsedMessage.airportId + "-" + parsedMessage.date.Format("2006-01-02") + "-" + msg.Topic() + ".csv"
	exists, _ := fileExists(filename)
	var csvFile *os.File
	if exists {
		csvFile, _ = os.OpenFile(filename, os.O_APPEND, 0644)
	} else {
		csvFile, _ = os.Create(filename)
	}

	insertIntoCsv(csvFile, parsedMessage)
	err := csvFile.Close()
	if err != nil {
		log.Println(err)
	}
}

func main() {
	config := subconfig.LoadConfig("textfile")
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
