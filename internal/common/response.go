package common

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

func MakeError(w http.ResponseWriter, code int, domain string, message string, method string) {
	logrus.WithFields(
		logrus.Fields{
			"type":   code,
			"domain": domain,
			"method": method,
		}).Error(strings.ToLower(message))
	http.Error(w, message, code)
}

func EncodeResponse(w http.ResponseWriter, response interface{}) {
	enc := json.NewEncoder(w)
	//enc.SetEscapeHTML(false)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	enc.Encode(response)
}
