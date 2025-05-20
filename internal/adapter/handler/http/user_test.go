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

	handler "go-hexagonal-architecture/internal/adapter/handler/http"
	"go-hexagonal-architecture/internal/core/domain"
	"go-hexagonal-architecture/internal/core/port"
	"go-hexagonal-architecture/internal/core/port/mock"
)

func TestUserHandler_Register(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		body       map[string]interface{}
		dependency func(ctrl *gomock.Controller) port.UserService
		wantStatus int
	}{
		{
			name: "success",
			body: map[string]interface{}{"name": "John", "email": "john@example.com", "password": "password1234"},
			dependency: func(ctrl *gomock.Controller) port.UserService {
				m := mock.NewMockUserService(ctrl)
				m.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
				return m
			},
			wantStatus: 200,
		},
		{
			name: "validation error",
			body: map[string]interface{}{"name": "", "email": "", "password": ""},
			dependency: func(ctrl *gomock.Controller) port.UserService {
				return nil
			},
			wantStatus: 400,
		},
		{
			name: "service error",
			body: map[string]interface{}{"name": "John", "email": "john@example.com", "password": "password"},
			dependency: func(ctrl *gomock.Controller) port.UserService {
				m := mock.NewMockUserService(ctrl)
				m.EXPECT().Create(gomock.Any(), gomock.Any()).Return(domain.ErrInternal)
				return m
			},
			wantStatus: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			h := handler.NewUserHandler(tt.dependency(ctrl))
			r := gin.New()
			r.POST("/register", h.Register)

			w := httptest.NewRecorder()
			body, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

func TestUserHandler_List(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tests := []struct {
		name       string
		query      string
		dependency func(ctrl *gomock.Controller) port.UserService
		wantStatus int
	}{
		{
			name:  "success",
			query: "?page=1&size=2",
			dependency: func(ctrl *gomock.Controller) port.UserService {
				m := mock.NewMockUserService(ctrl)
				m.EXPECT().List(gomock.Any(), gomock.Eq(int64(1)), gomock.Eq(int64(2))).Return([]*domain.User{{ID: "1", Name: "A", Email: "a@b.com"}}, nil)
				return m
			},
			wantStatus: 200,
		},
		{
			name:  "query not valid",
			query: "?page=0&size=2",
			dependency: func(ctrl *gomock.Controller) port.UserService {
				m := mock.NewMockUserService(ctrl)
				return m
			},
			wantStatus: 400,
		},
		{
			name:  "service error",
			query: "?page=1&size=2",
			dependency: func(ctrl *gomock.Controller) port.UserService {
				m := mock.NewMockUserService(ctrl)
				m.EXPECT().List(gomock.Any(), int64(1), int64(2)).Return(nil, assert.AnError)
				return m
			},
			wantStatus: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			h := handler.NewUserHandler(tt.dependency(ctrl))
			r := gin.New()
			r.GET("/users", h.List)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/users"+tt.query, nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

func TestUserHandler_GetByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tests := []struct {
		name       string
		id         string
		dependency func(ctrl *gomock.Controller) port.UserService
		wantStatus int
	}{
		{
			name: "success",
			id:   "1",
			dependency: func(ctrl *gomock.Controller) port.UserService {
				m := mock.NewMockUserService(ctrl)
				m.EXPECT().GetByID(gomock.Any(), "1").Return(&domain.User{ID: "1", Name: "A", Email: "a@b.com"}, nil)
				return m
			},
			wantStatus: 200,
		},
		{
			name: "not found",
			id:   "2",
			dependency: func(ctrl *gomock.Controller) port.UserService {
				m := mock.NewMockUserService(ctrl)
				m.EXPECT().GetByID(gomock.Any(), "2").Return(nil, assert.AnError)
				return m
			},
			wantStatus: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			h := handler.NewUserHandler(tt.dependency(ctrl))
			r := gin.New()
			r.GET("/users/:id", h.GetByID)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/users/"+tt.id, nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

func TestUserHandler_Update(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tests := []struct {
		name       string
		id         string
		body       map[string]interface{}
		dependency func(ctrl *gomock.Controller) port.UserService
		wantStatus int
	}{
		{
			name: "success",
			id:   "1",
			body: map[string]interface{}{"name": "A", "email": "test@gmail.com"},
			dependency: func(ctrl *gomock.Controller) port.UserService {
				m := mock.NewMockUserService(ctrl)
				m.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
				return m
			},
			wantStatus: 200,
		},
		{
			name: "validation error",
			id:   "1",
			body: map[string]interface{}{"name": "", "email": ""},
			dependency: func(ctrl *gomock.Controller) port.UserService {
				m := mock.NewMockUserService(ctrl)

				return m
			},
			wantStatus: 400,
		},
		{
			name: "service error",
			id:   "1",
			body: map[string]interface{}{"name": "A", "email": "a@b.com"},
			dependency: func(ctrl *gomock.Controller) port.UserService {
				m := mock.NewMockUserService(ctrl)
				m.EXPECT().Update(gomock.Any(), gomock.Any()).Return(assert.AnError)
				return m
			},
			wantStatus: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			h := handler.NewUserHandler(tt.dependency(ctrl))

			r := gin.New()
			r.PUT("/users/:id", h.Update)
			w := httptest.NewRecorder()
			body, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("PUT", "/users/"+tt.id, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			r.ServeHTTP(w, req)
			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

func TestUserHandler_Delete(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tests := []struct {
		name       string
		id         string
		dependency func(ctrl *gomock.Controller) port.UserService
		wantStatus int
	}{
		{
			name: "success",
			id:   "1",
			dependency: func(ctrl *gomock.Controller) port.UserService {
				m := mock.NewMockUserService(ctrl)
				m.EXPECT().Delete(gomock.Any(), "1").Return(nil)
				return m
			},
			wantStatus: 200,
		},
		{
			name: "service error",
			id:   "2",
			dependency: func(ctrl *gomock.Controller) port.UserService {
				m := mock.NewMockUserService(ctrl)
				m.EXPECT().Delete(gomock.Any(), "2").Return(assert.AnError)
				return m
			},
			wantStatus: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			h := handler.NewUserHandler(tt.dependency(ctrl))
			r := gin.New()
			r.DELETE("/users/:id", h.Delete)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", "/users/"+tt.id, nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}
