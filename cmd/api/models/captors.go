package models

import (
	"time"
)

type parsedMessage struct {
	captorId    string 		`json:"captorId"`
	airportId   string		`json:"airportId"`
	measureType string		`json:"measureType"`
	value       string		`json:"value"`
	date        time.Time	`json:"date"`
}

type Message []parsedMessage

/*
func GetTemp(start time.Time , end time.Time ) Message {

	result,err := config.CreateConnexion().Do()

	if err != nil {
		log.Fatal(err)
	}
	json

	return result;

}
 */