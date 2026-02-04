// Пакет предоставляет хендлеры для обработки CRUDL операций
package handlers

import (
	"database/sql"
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

		if sub.Price <= 0 || sub.ServiceName == "" || sub.UserID == "" || sub.StartDate == "" {
			err := errors.New(consts.EmptyValue)
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.EmptyValue, http.StatusBadRequest)
			return
		}

		if _, err := time.Parse(consts.TimeFormat, sub.StartDate); err != nil {
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.InvalidDate, http.StatusBadRequest)
			return
		}

		if err := db.CreateSubscription(&sub); err != nil {
			errs.ErrLogAndResp(w, err, consts.InternalServerError, http.StatusInternalServerError)
			return
		}

		location := fmt.Sprintf(consts.APIPathV1+"/%s_%s", sub.UserID, sub.ServiceName)
		w.Header().Set("Location", location)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(sub)
	}
}

func GetSubscription(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			err := errors.New(consts.MethodNotAllowed)
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.MethodNotAllowed, http.StatusMethodNotAllowed)
			return
		}

		userID := r.PathValue(consts.UserID)
		serviceName := r.PathValue(consts.ServiceName)

		if userID == "" || serviceName == "" {
			err := errors.New(consts.EmptyValue)
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.EmptyValue, http.StatusBadRequest)
			return
		}

		sub := structs.Subscription{
			UserID:      userID,
			ServiceName: serviceName,
		}

		result, err := db.GetSubscription(&sub)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				errs.ErrLogAndResp(w, err, consts.NotExist, http.StatusNotFound)
				return
			}
			errs.ErrLogAndResp(w, err, consts.InternalServerError, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func UpdateSubscription(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			err := errors.New(consts.MethodNotAllowed)
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.MethodNotAllowed, http.StatusBadRequest)
			return
		}

		var sub structs.Subscription
		if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.BadInput, http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		if sub.Price <= 0 {
			err := errors.New(consts.EmptyValue)
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.EmptyValue, http.StatusBadRequest)
			return
		}

		userID, serviceName := r.PathValue(consts.UserID), r.PathValue(consts.ServiceName)

		if userID == "" || serviceName == "" {
			err := errors.New(consts.EmptyValue)
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.EmptyValue, http.StatusBadRequest)
			return
		}

		sub.UserID, sub.ServiceName = userID, serviceName

		result, err := db.UpdateSubscription(&sub)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				errs.ErrLogAndResp(w, err, consts.NotExist, http.StatusNotFound)
				return
			}
			errs.ErrLogAndResp(w, err, consts.InternalServerError, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func DeleteSubscription(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			err := errors.New(consts.MethodNotAllowed)
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.MethodNotAllowed, http.StatusBadRequest)
			return
		}

		var sub structs.Subscription
		if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.BadInput, http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		userID, serviceName := r.PathValue(consts.UserID), r.PathValue(consts.ServiceName)

		if userID == "" || serviceName == "" {
			err := errors.New(consts.EmptyValue)
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.EmptyValue, http.StatusBadRequest)
			return
		}

		sub.UserID, sub.ServiceName = userID, serviceName
		err := db.DeleteSubscription(&sub)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				errs.ErrLogAndResp(w, err, consts.NotExist, http.StatusFound)
				return
			}
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.InternalServerError, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}
