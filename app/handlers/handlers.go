package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gimaevra94/test-effective-mobile/app/database"
	"github.com/gimaevra94/test-effective-mobile/app/errs"
	"github.com/gimaevra94/test-effective-mobile/app/structs"
)

func CreateSubscription(w http.ResponseWriter, r *http.Request) {
	var sub structs.Subscription

	if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
		errs.ErrLogAndResp(w, err, "Bad input", http.StatusBadRequest)
		return
	}

	if err := database.CreateSubscription(&sub); err != nil {
		errs.ErrLogAndResp(w, err, "Internal server error", 500)
	}
}
