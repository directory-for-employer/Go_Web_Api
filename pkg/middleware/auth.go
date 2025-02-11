package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		token := strings.TrimPrefix(bearer, "Bearer ")
		if bearer == "" {
			next.ServeHTTP(w, r)
			return
		}
		fmt.Println(token)
	})
}
