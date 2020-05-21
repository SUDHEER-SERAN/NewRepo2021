package user

import (
	"encoding/json"
	"jobreport/internal/common"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type loginHandler struct {
	service Service
}

func MakeLoginHandler(mr *mux.Router, s Service) http.Handler {
	h := &loginHandler{
		service: s,
	}
	mr.HandleFunc("/createUser", h.CreateUser).Methods("POST")
	mr.HandleFunc("/login", h.Login).Methods("POST")

	return mr
}

func (h *loginHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	entity := User{}

	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		logrus.WithError(err).Error("unable to unmarshal User entry")
		common.MakeError(w, http.StatusBadRequest, "CreateUser", "Bad Request", "create")
		return
	}

	if entity.Username == "" || entity.Password == "" {
		logrus.Warn("Missing Passowrd or Username")
		common.MakeError(w, http.StatusUnauthorized, "CreateUser", "missing auth Details", "login")
	}
	if err := h.service.CreateUser(ctx, entity); err != nil {
		common.MakeError(w, http.StatusUnauthorized, "", "Canot create the user", "login")
		return
	}
	common.EncodeResponse(w, "{}")

}

func (h *loginHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	entity := User{}

	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		logrus.WithError(err).Error("unable to unmarshal User entry")
		common.MakeError(w, http.StatusBadRequest, "User", "Bad Request", "create")
		return
	}

	if entity.Username == "" || entity.Password == "" {
		logrus.Warn("Missing Passowrd or Username")
		common.MakeError(w, http.StatusUnauthorized, "login", "missing auth Details", "login")
	}

	if err := h.service.AuthenticateUser(ctx, entity); err != nil {
		logrus.Warn("Unatorized Access")
		common.MakeError(w, http.StatusUnauthorized, "login", "missing auth Details", "login")
	}

}