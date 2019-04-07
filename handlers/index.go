package handlers

import (
	"fmt"
	"net/http"
)

// IndexHandler serves default route with homepage
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "GET request OK")
}
