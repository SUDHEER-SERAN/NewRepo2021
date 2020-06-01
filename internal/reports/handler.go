package reports

import (
	"encoding/json"
	"jobreport/internal/common"
	"net/http"
	"strconv"

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
	mr.HandleFunc("/report/getreports", h.getReports).Methods("GET")
	mr.HandleFunc("/report/getjrdropdownlist/{id}", h.getjrList).Methods("GET")
	return mr
}

func (h *reportsHandler) initializePage(w http.ResponseWriter, r *http.Request) {
	context := r.Context()
	list, err := h.service.initializePage(context)

	if err != nil {
		common.MakeError(w, http.StatusUnauthorized, "", "Canot create the user", "login")
		return
	}
	json.NewEncoder(w).Encode(list)
}

func (h *reportsHandler) generateReport(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reportEntity := JobReportBasicDetails{}

	if err := json.NewDecoder(r.Body).Decode(&reportEntity); err != nil {
		logrus.WithError(err).Error("Unable to unmarshal ReportDetails entry")
		common.MakeError(w, http.StatusBadRequest, "generateReport", "Bad Request", "create")
		return
	}
	if err := h.service.generateReport(ctx, reportEntity); err != nil {
		common.MakeError(w, http.StatusBadRequest, "generateReport", "Bad Request", "create")
		return
	}
	json.NewEncoder(w).Encode("{}")
}
func (h *reportsHandler) getReports(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := h.service.getReports(ctx); err != nil {
		common.MakeError(w, http.StatusBadRequest, "generateReport", "Bad Request", "create")
		return
	}
	json.NewEncoder(w).Encode("{}")

}
func (h *reportsHandler) getjrList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)
	idStr := params["id"]

	if idStr == "" {
		json.NewEncoder(w).Encode("[]")
		return
	}
	id, _ := strconv.Atoi(idStr)
	keys, ok := r.URL.Query()["searchKey"]
	var searchKey string
	if !ok {
		searchKey = ""
	}
	searchKey = keys[0]
	list, err := h.service.getjrList(ctx, id, searchKey)
	if err != nil {
		common.MakeError(w, http.StatusBadRequest, "getjrList", "Bad Request", "fetch")
		return
	}
	json.NewEncoder(w).Encode(list)
}
