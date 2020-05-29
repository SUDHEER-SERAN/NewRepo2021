package reports

import (
	"encoding/json"
	"jobreport/internal/common"
	"net/http"

	"github.com/gorilla/mux"
)

type reportsHandler struct {
	service Service
}

func MakeReportHandler(mr *mux.Router, s Service) http.Handler {
	h := &reportsHandler{
		service: s,
	}
	mr.HandleFunc("/report/initialize-page", h.initializePage).Methods("GET")

	return mr
}

func (h *reportsHandler) initializePage(w http.ResponseWriter, r *http.Request) {
	context := r.Context()
	data, err := h.service.initializePage(context)

	if err != nil {
		common.MakeError(w, http.StatusUnauthorized, "", "Canot create the user", "login")
		return
	}
	json.NewEncoder(w).Encode(data)
}
