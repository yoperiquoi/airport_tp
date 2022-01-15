package database

import (
	redistimeseries "github.com/RedisTimeSeries/redistimeseries-go"
	"log"
)

/*
	Function which returns the connection to the redis database
*/
func CreateConnexion() *redistimeseries.Client {
	log.Println("Connecting to the database")
	return redistimeseries.NewClient("localhost:6379", "nohelp", nil)
}
