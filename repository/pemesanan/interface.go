package pemesanan

import "rogerdev_golf/entities"

type Pemesanan interface {
	Create(pemesanan *entities.Pemesanan) error
	Get(pemesananId string) (entities.PemesananResponseFormat, error)
	GetAll() ([]entities.PemesananResponseFormat, error)
	GetAllDatatables() ([]entities.PemesananResponseFormatDatatables, int, error)
	Update(pemesanan *entities.Pemesanan) error
	Delete(pemesananId string) (int64, error)
}
