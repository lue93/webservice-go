package services

import (
	"example/web-service-go/application/core/domain/models"
	"example/web-service-go/application/port/out/repository"
)

type AllAlbumsService struct {
	Repo repository.GetAllAlbumsRepo
}

func (a AllAlbumsService) GetAllAlbumsExec() []models.Album {
	a.Repo.GetAllAlbumsData()
}
