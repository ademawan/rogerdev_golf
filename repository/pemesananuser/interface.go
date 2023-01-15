package pemesananuser

import "rogerdev_golf/entities"

type PemesananUser interface {
	Create(pemesanan *entities.PemesananUser) error
	Get(pemesananId string) (entities.PemesananUserResponseFormat, error)
	GetAll() ([]entities.PemesananUserResponseFormat, error)
	GetAllDatatables() ([]entities.PemesananUserResponseFormatDatatables, int, error)
	Update(pemesanan *entities.PemesananUser) error
	Delete(pemesananId string) (int64, error)
}
