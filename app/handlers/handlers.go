package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gimaevra94/test-effective-mobile/app/database"
	"github.com/gimaevra94/test-effective-mobile/app/errs"
	"github.com/gimaevra94/test-effective-mobile/app/structs"
	"github.com/pkg/errors"
)

func CreateSubscription(w http.ResponseWriter, r *http.Request) {
	var sub structs.Subscription

	if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
		errs.JSONErr(w, err.Error(), http.StatusBadRequest)
		log.Printf("%+v", errors.WithStack(err))
		return
	}

	if err := database.CreateSubscription(&sub); err != nil {
		log.Printf("%+v", err)
	}
}
