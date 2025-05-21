package token_test

import (
	"log"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"go-hexagonal-architecture/internal/adapter/auth/token"
	"go-hexagonal-architecture/internal/adapter/config"
	"go-hexagonal-architecture/internal/core/domain"
)

func TestTokenAuth_New(t *testing.T) {
	tests := []struct {
		name      string
		secret    string
		duration  string
		expectErr bool
	}{
		{
			name:      "valid config",
			secret:    "secret",
			duration:  "1h",
			expectErr: false,
		},
		{
			name:      "invalid duration",
			secret:    "secret",
			duration:  "invalid",
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cfg := &config.Config{
				JWTSecret:     test.secret,
				JWTExpiration: test.duration,
			}
			tokenAuth, err := token.New(cfg)

			if test.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, tokenAuth)
		})
	}
}

func TestTokenAuth_CreateToken(t *testing.T) {
	tests := []struct {
		name      string
		user      *domain.User
		secret    string
		duration  string
		expectErr bool
	}{
		{
			name:      "valid token",
			user:      &domain.User{ID: uuid.NewString()},
			secret:    "secret",
			duration:  "1h",
			expectErr: false,
		},
		{
			name:      "empty secret",
			user:      &domain.User{ID: uuid.NewString()},
			secret:    "",
			duration:  "1h",
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			log.Println("secret", test.secret)
			cfg := &config.Config{
				JWTSecret:     test.secret,
				JWTExpiration: test.duration,
			}
			tokenAuth, err := token.New(cfg)
			assert.NoError(t, err)

			token, err := tokenAuth.CreateToken(test.user)
			if test.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.NotEmpty(t, token)
		})
	}
}

func TestTokenAuth_VerifyToken(t *testing.T) {
	secretKey := "TEST_SECRET"

	mockUserID := "MOCK_USER_ID"
	claims := jwt.MapClaims{
		"sub": mockUserID,
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ := jwtToken.SignedString([]byte(secretKey))
	failedToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2V4YW1wbGUuYXV0aDAuY29tLyIsImF1ZCI6Imh0dHBzOi8vYXBpLmV4YW1wbGUuY29tL2NhbGFuZGFyL3YxLyIsInN1YiI6InVzcl8xMjMiLCJpYXQiOjE0NTg3ODU3OTYsImV4cCI6MTQ1ODg3MjE5Nn0.CA7eaHjIHz5NxeIJoFK9krqaeZrPLwmMmgI_XiQiIkQ"

	tests := []struct {
		name      string
		token     string
		secret    string
		duration  string
		expectErr bool
	}{
		{
			name:      "valid token",
			token:     tokenStr,
			secret:    secretKey,
			duration:  "1h",
			expectErr: false,
		},
		{
			name:      "failed token",
			token:     failedToken,
			secret:    secretKey,
			duration:  "1h",
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cfg := &config.Config{
				JWTSecret:     test.secret,
				JWTExpiration: test.duration,
			}

			tokenAuth, err := token.New(cfg)
			assert.NoError(t, err)
			payload, err := tokenAuth.VerifyToken(test.token)

			if test.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.NotEmpty(t, payload)
			assert.Equal(t, mockUserID, payload.ID)
		})
	}
}
