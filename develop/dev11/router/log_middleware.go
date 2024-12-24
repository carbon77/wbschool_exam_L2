package router

import (
	"log"
	"net/http"
)

func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request[addr=%s, method=%s, path=%s, query=[%s]]\n", r.RemoteAddr, r.Method, r.URL.Path, r.URL.RawQuery)
		next.ServeHTTP(w, r)
	}
}
