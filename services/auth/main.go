package main

import (
	"log"
	"net/http"

	"github.com/Lucas-Melo0/goChatApp/services/auth/internal/handlers"
	"github.com/Lucas-Melo0/goChatApp/services/auth/internal/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/register", handlers.Register)
	r.Post("/login", handlers.Login)

	protectedRoutes := chi.NewRouter()
	protectedRoutes.Use(middlewares.JWTMiddleware)
	r.Mount("/protected", protectedRoutes)

	log.Fatal(http.ListenAndServe(":8080", r))
}
