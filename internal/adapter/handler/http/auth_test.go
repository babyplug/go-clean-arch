package http_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"go-hexagonal-architecture/internal/core/domain"
	"go-hexagonal-architecture/internal/core/port/mock"

	handler "go-hexagonal-architecture/internal/adapter/handler/http"
)

func TestAuthHandler_Login(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tests := []struct {
		name       string
		body       map[string]interface{}
		dependency func(ctrl *gomock.Controller) *mock.MockAuthService
		wantStatus int
	}{
		{
			name: "success",
			body: map[string]interface{}{"email": "john@example.com", "password": "password1234"},
			dependency: func(ctrl *gomock.Controller) *mock.MockAuthService {
				m := mock.NewMockAuthService(ctrl)
				m.EXPECT().Login(gomock.Any(), "john@example.com", "password1234").Return("token123", nil)
				return m
			},
			wantStatus: 200,
		},
		{
			name: "validation error",
			body: map[string]interface{}{"email": "", "password": ""},
			dependency: func(ctrl *gomock.Controller) *mock.MockAuthService {
				return mock.NewMockAuthService(ctrl)
			},
			wantStatus: 400,
		},
		{
			name: "invalid credentials",
			body: map[string]interface{}{"email": "john@example.com", "password": "wrongpass"},
			dependency: func(ctrl *gomock.Controller) *mock.MockAuthService {
				m := mock.NewMockAuthService(ctrl)
				m.EXPECT().Login(gomock.Any(), "john@example.com", "wrongpass").Return("", domain.ErrInvalidCredentials)
				return m
			},
			wantStatus: 401,
		},
		{
			name: "internal error",
			body: map[string]interface{}{"email": "john@example.com", "password": "password1234"},
			dependency: func(ctrl *gomock.Controller) *mock.MockAuthService {
				m := mock.NewMockAuthService(ctrl)
				m.EXPECT().Login(gomock.Any(), "john@example.com", "password1234").Return("", domain.ErrInternal)
				return m
			},
			wantStatus: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			h := handler.NewAuthHandler(tt.dependency(ctrl))
			r := gin.New()
			r.POST("/auth/login", h.Login)

			w := httptest.NewRecorder()
			body, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			r.ServeHTTP(w, req)
			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}
