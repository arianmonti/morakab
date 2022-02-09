package handlers

import (
	"fmt"
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
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error parsing form"))
		return
	}
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

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

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User Logged in successfully"))
}
