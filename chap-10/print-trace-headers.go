package main

import (
	"fmt"
	"log"
	"net/http"
)

func tracingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	incomingHeaders := []string{
		"x-request-id",
		"x-b3-traceid",
		"x-b3-spanid",
		"x-b3-parentspanid",
		"x-b3-sampled",
		"x-b3-flags",
		"x-ot-span-context",
	}

	return func(w http.ResponseWriter, r *http.Request) {
		for _, th := range incomingHeaders {
			w.Header().Set(th, r.Header.Get(th))
		}
		next.ServeHTTP(w, r)
	}
}

func main() {
	http.HandleFunc("/", tracingMiddleware(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello headers, %v", r.Header)
	}))

	log.Fatal(http.ListenAndServe(":8081", nil))
}