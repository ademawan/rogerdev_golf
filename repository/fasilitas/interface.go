package fasilitas

import (
	"rogerdev_golf/entities"
)

type Fasilitas interface {
	Create(user *entities.Fasilitas) error
	Get(fasilitasId string) (entities.FasilitasResponseFormat, error)
	GetAll() ([]entities.FasilitasResponseFormat, error)
	GetAllDatatables() ([]entities.FasilitasResponseFormatDatatables, int, error)
	Update(newFasilitas *entities.Fasilitas) error
	Delete(fasilitasId string) (int64, error)
}
