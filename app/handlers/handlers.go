package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gimaevra94/test-effective-mobile/app/structs"
)

func JSONErr(w http.ResponseWriter, msg string, statusCode int) {
	resp := structs.Responce{
		Msg: msg,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}

func CreateSubscription(w http.ResponseWriter, r *http.Request) {
	var sub structs.Subscription

	if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
		JSONErr(w, err.Error(), http.StatusBadRequest)
		return
	}

}
