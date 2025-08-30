package services

import (
	"example/web-service-go/infrastructure/adapters/out/repo"
)

type AllAlbumsService struct {
	Repo repo.Album
}

func (a AllAlbumsService) GetAllAlbumsExec() []repo.Album {
	return a.Repo.GetAllAlbumsData()
}
