package main

import (
	"net/http"

	"github.com/ersa97/NATS/publisher/controllers"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Content-Type", "application/json")
			h.ServeHTTP(rw, r)
		})
	})

	r.HandleFunc("/", controllers.TestApi).Methods("GET")

	http.ListenAndServe(":8000", r)
}
