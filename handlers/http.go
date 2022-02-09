package handlers

import (
	"morakab/pkg"
	"net/http"
)

type HTTPHandler struct {
	Morakab *pkg.Morakab
}

func (h *HTTPHandler) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
