package galeri

import (
	"rogerdev_golf/entities"
)

type Galeri interface {
	Create(user *entities.Galeri) error
	Get(galeriId string) (entities.GaleriResponseFormat, error)
	GetAll() ([]entities.GaleriResponseFormat, error)
	GetAllDatatables() ([]entities.GaleriResponseFormatDatatables, int, error)
	Update(galeri *entities.Galeri) error
	Delete(galeriId string) (int64, error)
}
