package database

import (
	redistimeseries "github.com/RedisTimeSeries/redistimeseries-go"
)

func CreateConnexion() *redistimeseries.Client {
	return redistimeseries.NewClient("localhost:6379", "nohelp", nil)
}

