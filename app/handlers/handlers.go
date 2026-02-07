// Пакет предоставляет хендлеры для обработки CRUDL операций
package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gimaevra94/test-effective-mobile/app/consts"
	"github.com/gimaevra94/test-effective-mobile/app/database"
	"github.com/gimaevra94/test-effective-mobile/app/errs"
	"github.com/gimaevra94/test-effective-mobile/app/structs"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Функция реализует 'create' API.
// Приходящий запрос декодируется.
// В запросе структура которую необходимо сохранить в базу данных.
// Проверяется на наличие пустых полей.
// Поле даты проветяется на соответствие формату.
// Если все проверки пройдены переменная,
// в которую был декодирован запрос в виде структуры передается в db.CreateSubscription
// для работы с базой данных.
// Выполняется запрос к базе данных, результат запроса кодируется в json и отправляется клиенту.
func CreateSubscription(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			err := errors.New(consts.MethodNotAllowed)
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.MethodNotAllowed, http.StatusMethodNotAllowed)
			return
		}

		var sub structs.Subscription
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.BadInput, http.StatusBadRequest)
			return
		}

		if sub.ServiceName == "" || sub.Price <= 0 || sub.UserID == "" || sub.StartDate == "" {
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

		location := fmt.Sprintf(consts.APIPathV1+"/%s/%s", sub.UserID, sub.ServiceName)
		w.Header().Set("Location", location)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(sub)
	}
}

// Функция реализует 'get' API.
// Из пути запроса берутся поля составляющие ключ для поиска в базе данных.
// Проверяются на наличие пустых полей и передаются в виде структуры в db.GetSubscriprion для работы с базой данных.
// Выполняется запрос к базе данных, результат запроса кодируется в json и отправляется клиенту.
func GetSubscription(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			err := errors.New(consts.MethodNotAllowed)
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.MethodNotAllowed, http.StatusMethodNotAllowed)
			return
		}

		serviceName, userID := r.PathValue(consts.UserID), r.PathValue(consts.ServiceName)
		if serviceName == "" || userID == "" {
			err := errors.New(consts.EmptyValue)
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.EmptyValue, http.StatusBadRequest)
			return
		}

		sub := structs.Subscription{
			ServiceName: serviceName,
			UserID:      userID,
		}

		result, err := db.GetSubscription(&sub)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				errs.ErrLogAndResp(w, err, consts.NotExist, http.StatusNotFound)
			}
			errs.ErrLogAndResp(w, err, consts.InternalServerError, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

// Функция реализует 'update' API.
// Приходящий запрос декодируется.
// В запросе хранится новое значение которым нужно заменить значение из базы данных.
// Проверяется на наличие пустых полей и в виде структуры передается в db.UpdateSubscription
// для работы с базой данных.
// Выполняется запрос к базе данных, результат запроса кодируется в json и отправляется клиенту.
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

		serviceName, userID := r.PathValue(consts.ServiceName), r.PathValue(consts.UserID)
		if serviceName == "" || userID == "" || sub.Price <= 0 {
			err := errors.New(consts.EmptyValue)
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.EmptyValue, http.StatusBadRequest)
			return
		}

		sub.ServiceName, sub.UserID = serviceName, userID
		result, err := db.UpdateSubscription(&sub)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				errs.ErrLogAndResp(w, err, consts.NotExist, http.StatusNotFound)
			}
			errs.ErrLogAndResp(w, err, consts.InternalServerError, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

// Функция реализует 'delete' API.
// Приходящий запрос декодируется.
// Он хранит в себе поля составляющие ключ для поиска в базе данных по которым необходимо найти строку и удалить.
// Проверяется на наличие пустых полей и в виде структуры передается в db.UpdateSubscription
// для работы с базой данных.
// В результате успешного запроса к базе данных клиенту отправляется статусОК
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
			}
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.InternalServerError, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

// Функция реализует 'list' API.
// В качестве соединения с базой данных принимает не *database.DB, а *grom.DB.
// Фильтры из запроса проверяются на соответствие разрешенным фильтрам.
// Все фильтры прошедшие проверку добавляются к последовательности запроса к базе данных.
// Выполняется запрос к базе данных, результат запроса кодируется в json и отправляется клиенту.
func ListSubscription(gdb *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			err := errors.New(consts.MethodNotAllowed)
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.MethodNotAllowed, http.StatusBadRequest)
			return
		}

		query := r.URL.Query()
		dbQuery := gdb.Model(&structs.Subscription{})
		allowedFilters := []string{
			"service_name",
			"price",
			"user_id",
			"start_date",
		}

		for _, v := range allowedFilters {
			if q_v, ok := query[v]; ok && len(q_v) > 0 && q_v[0] != "" {

				switch v {
				case "price":
					if price, err := strconv.Atoi(q_v[0]); err == nil {
						dbQuery = dbQuery.Where(v+" = $1", price)
					}
				default:
					dbQuery = dbQuery.Where(v+" = $1", q_v[0])
				}
			}
		}

		subs := []structs.Subscription{}
		if err := dbQuery.Find(&subs); err != nil {
			errs.ErrLogAndResp(w, errors.WithStack(err.Error), consts.NotExist, http.StatusNotFound)
			return
		}

		if len(subs) == 0 {
			err := errors.New(consts.NotExist)
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.NotExist, http.StatusNotFound)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(subs)
	}
}

// Функция принимает из запроса ключи, проверяет их на соответствие формату,
// и вызывает db.GetPeriodPricesSum для поиска строк по этим ключам.
// База данных возвращает стоимость подписки за 1 месяц.
// После этого вызывается priceSum которая считает сумму подписок за период.
// Результат кодируется в JSON и отправляется клиенту.
func GetPeriodPricesSum(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			err := errors.New(consts.MethodNotAllowed)
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.MethodNotAllowed, http.StatusBadRequest)
			return
		}

		serviceName := r.PathValue(consts.ServiceName)
		userID := r.PathValue(consts.UserID)
		startDate := r.PathValue(consts.StartDate)
		if serviceName == "" || userID == "" {
			err := errors.New(consts.EmptyValue)
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.EmptyValue, http.StatusBadRequest)
			return
		}

		startDateTime, err := time.Parse(consts.TimeFormat, startDate)
		if err != nil {
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.InvalidDate, http.StatusBadRequest)
			return
		}

		sub := structs.Subscription{
			ServiceName: serviceName,
			UserID:      userID,
		}

		result, err := db.GetPeriodPrices(sub)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				errs.ErrLogAndResp(w, errors.WithStack(err), consts.NotExist, http.StatusNotFound)
				return
			}

			err := errors.New(consts.InternalServerError)
			errs.ErrLogAndResp(w, errors.WithStack(err), consts.InternalServerError, http.StatusInternalServerError)
			return
		}

		resSum := priceSum(result, startDateTime)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resSum)
	}
}

// Функция принимает стоимость подписки за один месяц и начало периода подписки.
// Вычисляет колличество месяцев которые прошли с начала периода подписки,
// и умножает их на стоимость подписки за один месяц.
func priceSum(oneMouthPrice int, startDate time.Time) int {
	totalDuration := time.Now().Sub(startDate)
	oneMonthsAsHourses := 30 * 24
	totalMonthsAsHourses := totalDuration.Hours()
	mouthCount := totalMonthsAsHourses / float64(oneMonthsAsHourses)
	priceSum := mouthCount * float64(oneMouthPrice)
	return int(priceSum)
}
