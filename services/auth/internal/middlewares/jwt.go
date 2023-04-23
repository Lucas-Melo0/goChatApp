package middlewares

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
			return
		}
		tokenString := authHeaderParts[1]
		token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected string method : %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})
		if err != nil {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(*jwt.RegisteredClaims)
		if !ok {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "userId", claims.ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
