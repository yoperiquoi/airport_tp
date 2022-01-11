package main

import (
	"airport_tp/cmd/api/controllers"
	"github.com/gorilla/mux"
)

func InitializeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/").Name("Index").HandlerFunc(controllers.SensorsIndex)
	router.Methods("GET").Path("/AverageForDay").Name("AverageForDay").HandlerFunc(controllers.AverageForDay)
	router.Methods("GET").Path("/GetMesureFromTypeInRange").Name("GetMesureFromTypeInRange").HandlerFunc(controllers.GetMesureFromTypeInRange)

	return router
}
