package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandleBookRequest(w http.ResponseWriter, r *http.Request) {

	title := chi.URLParam(r, "title")
	page := chi.URLParam(r, "page")

	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)

}
