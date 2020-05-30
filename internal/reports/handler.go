package reports

import (
	"encoding/json"
	"jobreport/internal/common"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type reportsHandler struct {
	service Service
}

func MakeReportHandler(mr *mux.Router, s Service) http.Handler {
	h := &reportsHandler{
		service: s,
	}
	mr.HandleFunc("/report/initialize-page", h.initializePage).Methods("GET")
	mr.HandleFunc("/report/generatereport", h.generateReport).Methods("POST")

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

func (h *reportsHandler) generateReport(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reportEntity := Report{}

	if err := json.NewDecoder(r.Body).Decode(&reportEntity); err != nil {
		logrus.WithError(err).Error("unable to unmarshal ReportDetails entry")
		common.MakeError(w, http.StatusBadRequest, "generateReport", "Bad Request", "create")
		return
	}
	h.service.generateReport(ctx, reportEntity)
}
