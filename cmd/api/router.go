package main

import (
	"airport_tp/cmd/api/controllers"
	"github.com/gorilla/mux"
)

func InitializeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/").Name("Index").HandlerFunc(controllers.SensorsIndex)
	router.HandleFunc("/AverageForDay/{airport_id}",controllers.AverageForDay)
	router.HandleFunc("/GetMesureFromTypeInRange/{airport_id}/{measureType}/{start}/{end}",controllers.GetMesureFromTypeInRange)

	return router
}
