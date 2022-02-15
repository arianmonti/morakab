package handlers

import (
	"encoding/json"
	"fmt"
	"morakab/models"
	"morakab/pkg"
	"net/http"
)

func (h *HTTPHandler) Register(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error parsing form"))
		return
	}
	username := r.PostForm.Get("username")
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")

	if username == "" || email == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing fields"))
		return
	}

	if err := h.Morakab.RegisterUser(username, email, password); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error registering user: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User registered"))
}

func (h *HTTPHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error parsing data"))
		return
	}
	username := user.Username
	password := user.Password

	if username == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing fields"))
		return
	}

	if err := h.Morakab.LoginUser(username, password); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	token, err := pkg.GenerateToken(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", token)
}
