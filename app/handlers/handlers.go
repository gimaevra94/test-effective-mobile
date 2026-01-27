package handlers

import (
	"encoding/json"
	"net/http"
)

type Subscription struct {
	ServiceName string `json:"service_name"`
	Price       int    `json:"price"`
	UserId      string `json:"user_id"`
	StartDate   string `json:"start_date"`
}

type Responce struct {
	Msg string `json:"msg"`
}

func JSONErr(w http.ResponseWriter, msg string, statusCode int) {
	resp := Responce{
		Msg: msg,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}

func CreateSubscription(w http.ResponseWriter, r *http.Request) {
	var sub Subscription

	if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
		JSONErr(w, err.Error(), http.StatusBadRequest)
		return
	}

	
}
