package main

import (
	"example/web-service-go/application/core/domain/services"
	"example/web-service-go/infrastructure/adapters/in/chttp"
	"example/web-service-go/infrastructure/adapters/out/repo"
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

func buildContainer() Container {
	repository := repo.Album{}
	service := services.AllAlbumsService{repository}
	controller := chttp.AllAlbumController{service}

	return Container{controller}
}

// A standalone program (as opposed to a library) is always in package main.
func main() {
	container := buildContainer()

	r := mux.NewRouter()
	r.Use(logger)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the home page with logging!")
	})

	r.HandleFunc("/albums", func(w http.ResponseWriter, r *http.Request) {
		container.allAlbumsController.Handle(w, r)
	})

	fmt.Println("🚀 Servidor iniciado em http://localhost:8080")

	// Start the server.
	http.ListenAndServe(":8080", r)

}
