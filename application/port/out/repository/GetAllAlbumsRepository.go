package repository

import "example/web-service-go/application/core/domain/models"

type GetAllAlbumsRepo interface {
	GetAllAlbumsData() []models.Album //qualificacao do tipo
}
