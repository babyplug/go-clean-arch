package http

import (
	"github.com/babyplug/go-clean-arch/internal/core/port"
)

type authHandler struct {
	service   *port.UserService
	jwtSecret string
}

// func NewAuthHandler(r *mux.Router, service *port.UserService, jwtSecret string) {
// 	h := &authHandler{service: service, jwtSecret: jwtSecret}
// 	r.HandleFunc("/auth/register", h.Register).Methods("POST")
// 	r.HandleFunc("/auth/login", h.Login).Methods("POST")
// }

// func (h *authHandler) Register(w http.ResponseWriter, r *http.Request) {
// 	var input domain.User
// 	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
// 		http.Error(w, "Invalid request", http.StatusBadRequest)
// 		return
// 	}

// 	user, err := h.service.RegisterUser(&input)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(user)
// }

// func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
// 	var creds struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}

// 	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
// 		http.Error(w, "Invalid request", http.StatusBadRequest)
// 		return
// 	}

// 	user, err := h.service.AuthenticateUser(creds.Email, creds.Password)
// 	if err != nil {
// 		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
// 		return
// 	}

// 	token, err := jwt.GenerateToken(user.ID, h.jwtSecret, time.Hour*24)
// 	if err != nil {
// 		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(map[string]string{"token": token})
// }
