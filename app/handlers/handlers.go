// Пакет предоставляет хендлеры для обработки CRUDL операций
package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

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
	defer r.Body.Close()

	if sub.Price == 0 || sub.ServiceName == "" || sub.StartDate == "" {
		err := errors.New("some value is empty")
		errs.ErrLogAndResp(w, err, "some value is empty", http.StatusBadRequest)
	}

	if _, err := time.Parse("01-2026", sub.StartDate); err != nil {
		errs.ErrLogAndResp(w, err, "Invalid date format. Use MM-YYYY", http.StatusBadRequest)
	}

	if err := database.CreateSubscription(&sub); err != nil {
		errs.ErrLogAndResp(w, err, "Internal server error", 500)
		return
	}

	location := fmt.Sprintf("/api/v1/subscription/%s_%s", sub.UserId, sub.ServiceName)
	w.Header().Set("Location", location)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sub)
}
