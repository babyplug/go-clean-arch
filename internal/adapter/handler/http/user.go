package http

import (
	"net/http"

	"github.com/babyplug/go-clean-arch/internal/core/port"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc port.UserService
}

func NewUserHandler(svc port.UserService) *UserHandler {
	return &UserHandler{svc}
}

func (h *UserHandler) List(ctx *gin.Context) {
	users, err := h.svc.List()
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, err)
		return
	}
	handleSuccess(ctx, users)
}

func (h *UserHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := h.svc.GetByID(id)
	if err != nil {
		handleError(ctx, http.StatusNotFound, err)
		return
	}
	handleSuccess(ctx, user)
}
