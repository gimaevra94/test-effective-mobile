package errs

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gimaevra94/test-effective-mobile/app/structs"
)

func JSONErr(w http.ResponseWriter, err string, statusCode int) {
	resp := structs.Responce{
		Msg: err,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)

	log.Printf(err)
}
