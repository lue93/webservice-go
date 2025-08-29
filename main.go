package main

import (
	"example/web-service-go/infrastructure/adapters/in/chttp"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Container struct {
	allAlbumsController chttp.AllAlbumController
}

// Logger middleware.
func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request URL:", r.URL)
		next.ServeHTTP(w, r)
	})
}

// A standalone program (as opposed to a library) is always in package main.
func main() {
	r := mux.NewRouter()
	r.Use(logger)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the home page with logging!")
	})

	r.HandleFunc("/albums", func(w http.ResponseWriter, r *http.Request) {
		controller := chttp.AllAlbumController{}
		controller.Handle(w, r)
	})
	// Start the server.
	http.ListenAndServe(":8080", r)

}
