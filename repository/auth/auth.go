package auth

import (
	"database/sql"
	"fmt"
	"rogerdev_golf/entities"
	"rogerdev_golf/middlewares"

	"errors"
)

type AuthDb struct {
	db *sql.DB
}

func New(db *sql.DB) *AuthDb {
	return &AuthDb{
		db: db,
	}
}

func (ad *AuthDb) Login(email, password string) (entities.User, error) {
	user := entities.User{}
	fmt.Println(email, password)
	query := `SELECT * from user Where user_email=?`
	err := ad.db.QueryRow(query, email).Scan(&user.UserId, &user.UserNama, &user.UserEmail, &user.UserPassword, &user.UserAlamat, &user.UserNoHp, &user.UserTipeId, &user.IsUser, &user.CreatedAt)
	if err != nil {
		return user, errors.New("Internal Server Error")
	}
	match := middlewares.CheckPasswordHash(password, user.UserPassword)

	if !match {
		return user, errors.New("incorrect password")
	}

	return user, nil
}

// func (ad *AuthDb) LoginGoogle(email string) (entities.User, error) {
// 	user := entities.User{}

// 	err := ad.db.QueryRow("SELECT * FROM user WHERE email=?", email).Scan(&user.UserUid, &user.Name, &user.Email, user.Password, user.Address, user.Gender, user.CreatedAt, user.DeletedAt)
// 	if err != nil {
// 		return user, errors.New("email not found")
// 	}

// 	return user, nil
// }
