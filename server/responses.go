package server

import (
	"net/http"
	"log"
	"encoding/json"
)

// ответ в json формате
func Json(w http.ResponseWriter, data interface{}, statusCode ... uint16) {
	log.Println(statusCode)
	var body, err = json.Marshal(data)
	if err != nil {
		body = []byte(`{"message": "Internal Server Error"}`) // TODO
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		if len(statusCode) == 1 {
			w.WriteHeader(int(statusCode[0]))
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func InternalServerError(w http.ResponseWriter, err error) {
	Json(w, response{Message: err.Error()})
}

type response struct {
	Message string `json:"message"`
}
