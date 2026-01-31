// Пакет предоставляет хендлеры для обработки CRUDL операций
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gimaevra94/test-effective-mobile/app/consts"
	"github.com/gimaevra94/test-effective-mobile/app/database"
	"github.com/gimaevra94/test-effective-mobile/app/errs"
	"github.com/gimaevra94/test-effective-mobile/app/structs"
	"github.com/pkg/errors"
)

// Функция реализует api "/api/v1/subscription".
// Приходящий запрос декодируется.
// Проверяется на наличие пустых полей.
// Поле даты проветяется на соответствие формату.
// Если все проверки пройдены переменная,
// в которую был декодирован запрос передается в CreateSubscription
func CreateSubscription(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			err := errors.New(consts.MethodNotAllowed)
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.MethodNotAllowed, http.StatusMethodNotAllowed)
			return
		}

		var sub structs.Subscription
		if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.BadInput, http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		if sub.Price <= 0 || sub.ServiceName == "" || sub.StartDate == "" {
			err := errors.New(consts.EmptyValue)
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.EmptyValue, http.StatusBadRequest)
			return
		}

		if _, err := time.Parse("01-2006", sub.StartDate); err != nil {
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.InvalidDate, http.StatusBadRequest)
			return
		}

		if err := db.CreateSubscription(&sub); err != nil {
			errs.ErrLogAndResp(w, err, consts.InternalServerError, 500)
			return
		}

		location := fmt.Sprintf("/api/v1/subscription/%s_%s", sub.UserId, sub.ServiceName)
		w.Header().Set("Location", location)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(sub)
	}
}

func GetSubscription(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var sub structs.Subscription

		if r.Method != http.MethodGet {
			err := errors.New(consts.MethodNotAllowed)
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.MethodNotAllowed, http.StatusMethodNotAllowed)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.BadInput, http.StatusBadRequest)
			return
		}

		if sub.Price <= 0 || sub.ServiceName == "" || sub.StartDate == "" || sub.UserId == "" {
			err := errors.New(consts.EmptyValue)
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.EmptyValue, http.StatusBadRequest)
			return
		}

		if _, err := time.Parse("01-2006", sub.StartDate); err != nil {
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.InvalidDate, http.StatusBadRequest)
			return
		}

		
	}
}
