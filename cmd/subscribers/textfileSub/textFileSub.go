package main

import (
	subconfig "airport_tp/internal/config/subscribersConfig"
	"airport_tp/internal/utils"
	subutils "airport_tp/internal/utils/subscribersUtils"
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

/*
	Function which inform if the file already exists
*/
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

/*
	Struct used to store parsed information of the message
*/
type parsedMessage struct {
	captorId    string
	airportId   string
	measureType string
	value       string
	date        time.Time
}

/*
	Function which inserts into the specified file the parsed message received
*/
func insertIntoCsv(csvFile *os.File, message parsedMessage) {
	// Create the writer on the file
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
	unixDate, _ := strconv.ParseInt(parsedString[4], 10, 64)
	date := time.Unix(unixDate, 0)
	parsedMessage := parsedMessage{
		captorId:    parsedString[0],
		airportId:   parsedString[1],
		measureType: parsedString[2],
		value:       parsedString[3],
		date:        date,
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
	// Connect to the broker
	subscriber := subutils.Connect("tcp://"+config.BrokerHost+":"+strconv.Itoa(config.BrokerPort), config.ClientId)
	// Disconnect of the broker at the end of the program
	defer func() {
		subscriber.Disconnect(250)
		log.Println(config.ClientId + " disconnect from the broker")
	}()
	// Subscribe to the 3 topics
	tokenSubscriberTemp := subscriber.Subscribe(utils.TopicTemp, byte(config.Qos), messageHandler)
	tokenSubscriberPressure := subscriber.Subscribe(utils.TopicPressure, byte(config.Qos), messageHandler)
	tokenSubscriberWind := subscriber.Subscribe(utils.TopicWind, byte(config.Qos), messageHandler)

	for {
		// Waiting on each topic it subscribed
		tokenSubscriberTemp.Wait()
		tokenSubscriberPressure.Wait()
		tokenSubscriberWind.Wait()
	}
}
