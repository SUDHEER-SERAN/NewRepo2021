package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type loginHandler struct {
	service Service
}

func MakeLoginHandler(mr *mux.Router, s Service) http.Handler {
	h := &loginHandler{
		service: s,
	}
	mr.HandleFunc("/auth/", h.Login).Methods("GET")

	return mr
}

func (h *loginHandler) Login(w http.ResponseWriter, r *http.Request) {

}
