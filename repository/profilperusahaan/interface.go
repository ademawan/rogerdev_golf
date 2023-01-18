package profilperusahaan

import (
	"rogerdev_golf/entities"
)

type ProfilPerusahaan interface {
	Create(profilPrusahaan entities.ProfilPerusahaan) error
	Get(profilPerusahaan string) (entities.ProfilPerusahaanResponseFormat, error)
	GetProfilPerusahaan() (entities.ProfilPerusahaanResponseFormat, error)
	GetAll() ([]entities.ProfilPerusahaanResponseFormat, error)
	GetAllDatatables() ([]entities.ProfilPerusahaanResponseFormatDatatables, int, error)
	Update(profilPrusahaan *entities.ProfilPerusahaan) error
	Delete(profilPerusahaan string) (int64, error)
}
