package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	// Key to authenticate the cookie store
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

// User represents a user model
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Mocked user database
var users = map[string]string{
	"user@example.com": "password123",
}

// LoginHandler handles user login
func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user exists and the password is correct
	if password, ok := users[user.Email]; ok && password == user.Password {
		session, _ := store.Get(r, "session")
		session.Values["authenticated"] = true
		session.Save(r, w)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Login successful")
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}

// LogoutHandler handles user logout
func (s *Server) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	session.Values["authenticated"] = false
	session.Save(r, w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Logout successful")
}
