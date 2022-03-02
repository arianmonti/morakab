package handlers

import (
	"encoding/json"
	"fmt"
	"morakab/models"
	"net/http"
)

func (h *HTTPHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error parsing data"))
		return
	}

	title := book.Title
	author := book.Author

	if title == "" || author == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Empty"))
		return
	}

	if err := h.Morakab.CreateBook(title, author); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("%s", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
