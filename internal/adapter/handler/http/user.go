package http

import (
	"net/http"

	"github.com/babyplug/go-clean-arch/internal/core/domain"
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
	users, err := h.svc.List(ctx)
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, err)
		return
	}
	handleSuccess(ctx, users)
}

func (h *UserHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := h.svc.GetByID(ctx, id)
	if err != nil {
		handleError(ctx, http.StatusNotFound, err)
		return
	}
	handleSuccess(ctx, user)
}

func (h *UserHandler) Update(ctx *gin.Context) {
	user := &domain.User{}

	err := h.svc.Update(ctx, user)
	if err != nil {
		handleError(ctx, http.StatusNotFound, err)
		return
	}
	handleSuccess(ctx, user)
}

func (h *UserHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.svc.Delete(ctx, id)
	if err != nil {
		handleError(ctx, http.StatusNotFound, err)
		return
	}

	handleSuccess(ctx, nil)
}

func (h *UserHandler) Register(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}

	if err := h.svc.Create(ctx, &user); err != nil {
		handleError(ctx, http.StatusInternalServerError, err)
		return
	}

	handleSuccess(ctx, user)
}
