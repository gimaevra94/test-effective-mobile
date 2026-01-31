// Пакет предоставляет модуль для обработки ошибок и отправки JSON ответа клиенту.
package errs

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gimaevra94/test-effective-mobile/app/structs"
	"github.com/pkg/errors"
)

// Функция предоставляет JSON ответ клиенту в случае ошибки и логирует ее.
// Принимает на вход Response для записи сообщения об ошибке, http статус
// И саму ошибку для логирования.
// Если ошибка приходит без трассировки функция добавляет к ошибке трассирвоку
// перед логированием.
func ErrLogAndResp(w http.ResponseWriter, err error, msg string, statusCode int) {
	resp := structs.Responce{
		Msg: msg,
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)

	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	if stErr, ok := err.(stackTracer); ok {
		log.Printf("%+v", stErr)
	} else {
		err := errors.New(err.Error())
		stErr := errors.WithStack(err)
		log.Printf("%+v", stErr)
	}
}
