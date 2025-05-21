package token

import (
	"errors"
	"time"

	"go-hexagonal-architecture/internal/adapter/config"
	"go-hexagonal-architecture/internal/core/domain"
	"go-hexagonal-architecture/internal/core/port"

	"github.com/golang-jwt/jwt/v5"
)

type jwtTokenAuth struct {
	secret   string
	duration time.Duration
}

func New(cfg *config.Config) (port.TokenService, error) {
	duration, err := time.ParseDuration(cfg.JWTExpiration)
	if err != nil {
		return nil, errors.New("invalid duration format")
	}

	return &jwtTokenAuth{secret: cfg.JWTSecret, duration: duration}, nil
}

func (ta *jwtTokenAuth) CreateToken(user *domain.User) (string, error) {
	if len(ta.secret) == 0 {
		return "", errors.New("secret is empty")
	}

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(ta.duration)

	claims := jwt.MapClaims{
		"sub": user.ID,
		"iat": issuedAt.Unix(),
		"exp": expiredAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(ta.secret))
}

func (ta *jwtTokenAuth) VerifyToken(token string) (*domain.TokenPayload, error) {
	tk, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenMalformed
		}
		return []byte(ta.secret), nil
	})

	if err != nil || !tk.Valid {
		return nil, errors.New("invalid token")
	}
	claims := tk.Claims.(jwt.MapClaims)
	return &domain.TokenPayload{
		ID: claims["sub"].(string),
	}, nil
}
