package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/Lucas-Melo0/goChatApp/services/auth/internal/models"
	"github.com/golang-jwt/jwt/v4"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Username == "" || user.Password == "" {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}

	models.Users = append(models.Users, &user)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "User created successfully")

}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, u := range models.Users {
		if u.Username == user.Username && u.Password == user.Password {
			tokenString, err := GenerateToken(u.ID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			io.WriteString(w, tokenString)
			return
		}
	}

	http.Error(w, "Invalid username or password", http.StatusUnauthorized)
}
func GenerateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
