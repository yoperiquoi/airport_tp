package models

import (
	"airport_tp/internal/database"
	"encoding/json"
	redistimeseries "github.com/RedisTimeSeries/redistimeseries-go"
	"time"
)

type Average struct {
	measureType string
	value float64
}

type Averages []Average

type Mesure struct {
	date time.Time
	value float64
}

type Mesures []Mesure

func GetAverageForADay(airport_id string) string{
	actualDate := time.Now().Unix()
	passedDate := actualDate - 86400

	rangeOptions := redistimeseries.DefaultMultiRangeOptions
	rangeOptions.SetWithLabels(true)
	rangeOptions.SetAggregation(redistimeseries.AvgAggregation, 86400)
	result, _ := database.CreateConnexion().MultiRangeWithOptions(passedDate, actualDate, rangeOptions, "airport_id=" + airport_id)
	resultJson, _ := json.Marshal(result)
	return string(resultJson)
}


func GetMesureFromTypeInRange(airport_id string,measureType string, start int64, end int64) string {
	rangeOptions := redistimeseries.DefaultMultiRangeOptions
	rangeOptions.WithLabels = true
	result, _ := database.CreateConnexion().MultiRangeWithOptions(start, end, rangeOptions, "airport_id=" + airport_id, "sensor_id=202101122")
	resultJson, _ := json.Marshal(result)

	return string(resultJson)
}