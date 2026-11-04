package handler

import (
	"encoding/json"
	"net/http"

	"github.com/deev/data/backend/internal/service"
)

type Handler struct {
	srv *service.Service
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{srv: srv}
}

func (h *Handler) Login(w http.ResponseWriter,r *http.Request) {
	var inp struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&inp); err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}

	if err := h.srv.Login(r.Context(),service.LoginData{
		UName: inp.UserName,
		Password: inp.Password,
	}); err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(inp)
}

func (h *Handler) SendMsg(w http.ResponseWriter,r *http.Request) {
	var req struct {
		Email string `json:"email"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}

	srv := service.NewEmailService()

	go srv.SendMSG(req.Email,"link")
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(req)
}