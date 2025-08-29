package usecase

import "example/web-service-go/application/core/domain/models"

type GetAllAlbumsUsecase interface {
	GetAllAlbumsExec() []models.Album
}
