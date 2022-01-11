package controllers

import (
	"airport_tp/cmd/api/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func SensorsIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Le serveur est bien démarré !")
}


func AverageForDay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)

	average := models.GetAverageForADay(vars["airport_id"])
	json.NewEncoder(w).Encode(average)
}

func GetMesureFromTypeInRange(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)

	start,_ := strconv.ParseInt(vars["start"],10,64)
	end,_ := strconv.ParseInt(vars["end"],10,64)

	result := models.GetMesureFromTypeInRange(vars["airport_id"],vars["measureType"], start, end)

	json.NewEncoder(w).Encode(result)
}
