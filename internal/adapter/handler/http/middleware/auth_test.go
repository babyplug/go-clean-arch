package middleware_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"clean-arch/internal/adapter/handler/http/middleware"
	"clean-arch/internal/core/domain"
	"clean-arch/internal/core/port"
	"clean-arch/internal/core/port/mock"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		header         string
		dependency     func(ctrl *gomock.Controller) port.TokenService
		expectStatus   int
		expectResponse string
	}{
		{
			name:   "missing header",
			header: "",
			dependency: func(ctrl *gomock.Controller) port.TokenService {
				return nil
			},
			expectStatus:   401,
			expectResponse: "Authorization header is missing",
		},
		{
			name:   "invalid header format",
			header: "Bearer",
			dependency: func(ctrl *gomock.Controller) port.TokenService {
				return nil
			},
			expectStatus:   401,
			expectResponse: "Invalid authorization header format",
		},
		{
			name:   "invalid type",
			header: "Basic sometoken",
			dependency: func(ctrl *gomock.Controller) port.TokenService {
				return nil
			},
			expectStatus:   401,
			expectResponse: "Invalid authorization type",
		},
		{
			name:   "invalid token",
			header: "Bearer sometoken",
			dependency: func(ctrl *gomock.Controller) port.TokenService {
				m := mock.NewMockTokenService(ctrl)
				m.EXPECT().VerifyToken("sometoken").Return(nil, errors.New("invalid token"))
				return m
			},
			expectStatus:   401,
			expectResponse: "Invalid token",
		},
		{
			name:   "valid token",
			header: "Bearer validtoken",
			dependency: func(ctrl *gomock.Controller) port.TokenService {
				m := mock.NewMockTokenService(ctrl)
				m.EXPECT().VerifyToken("validtoken").Return(&domain.TokenPayload{ID: uuid.New()}, nil)
				return m
			},
			expectStatus:   200,
			expectResponse: "ok",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			w := httptest.NewRecorder()
			r := gin.New()

			r.GET("/protected", middleware.AuthMiddleware(test.dependency(ctrl)), func(c *gin.Context) {
				c.JSON(200, gin.H{"msg": "ok"})
			})

			req, _ := http.NewRequest("GET", "/protected", nil)
			if test.header != "" {
				req.Header.Set("Authorization", test.header)
			}
			r.ServeHTTP(w, req)

			assert.Equal(t, test.expectStatus, w.Code)
			assert.Contains(t, w.Body.String(), test.expectResponse)
		})
	}
}
