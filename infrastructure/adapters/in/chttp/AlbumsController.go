package chttp

import (
	"encoding/json"
	"example/web-service-go/application/core/domain/models"
	"example/web-service-go/application/core/domain/services"
	"net/http"
)

type AllAlbumController struct {
	service services.AllAlbumsService
}

type AllAlbumsResponse struct {
	album []models.Album
}

func (a AllAlbumController) Handle(w http.ResponseWriter, r *http.Request) {
	data := a.service.GetAllAlbumsExec()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responseBody, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		http.Error(w, "Erro ao gerar JSON", http.StatusInternalServerError)
		return
	}
	w.Write(responseBody)
}
