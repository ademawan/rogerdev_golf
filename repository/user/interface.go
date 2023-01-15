package user

import (
	"rogerdev_golf/entities"
)

type User interface {
	Create(user entities.User) error
	Get(userId string) (entities.UserResponseFormat, error)
	GetAll() ([]entities.UserResponseFormat, error)
	GetAllDatatables() ([]entities.UserResponseFormatDatatables, int, error)
	Update(user *entities.User) error
	Delete(userId string) (int64, error)
}
