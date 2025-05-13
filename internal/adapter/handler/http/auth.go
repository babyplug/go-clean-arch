package http

import (
	"errors"
	"net/http"

	"github.com/babyplug/go-clean-arch/internal/core/port"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service   port.AuthService
	jwtSecret string
}

type AuthReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthHandler(service port.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var creds AuthReq

	if err := ctx.ShouldBindJSON(&creds); err != nil {
		handleError(ctx, http.StatusBadRequest, errors.New("Invalid request payload"))
		return
	}

	token, err := h.service.Login(ctx, creds.Email, creds.Password)
	if err != nil {
		handleError(ctx, http.StatusUnauthorized, errors.New("Invalid credentials"))
		return
	}

	rsp := newAuthResponse(token)

	handleSuccess(ctx, rsp)
}
