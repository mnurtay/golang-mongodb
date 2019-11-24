package auth

import (
	"io"
	"net/http"
)

// LoginFunc ...
func LoginFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
