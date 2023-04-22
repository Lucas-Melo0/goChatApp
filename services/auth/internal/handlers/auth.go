package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Println(string(body))
}

func Login(w http.ResponseWriter, r *http.Request) {
}
