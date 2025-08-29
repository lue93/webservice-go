package http

import "net/http"

type HttpController interface {
	Handle(w http.ResponseWriter, r *http.Request)
}
