package user

import (
	"database/sql"
	"errors"
	"fmt"
	domain "rogerdev_golf/entities"
	"rogerdev_golf/middlewares"

	"strings"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {

	return &UserRepository{db: db}
}
func (r *UserRepository) Create(user domain.User) error {
	_, res, err := r.CheckEmail(user.UserEmail)
	if err != nil {
		return err
	} else if res {
		return errors.New("user email already exist")
	}

	user.UserPassword, _ = middlewares.HashPassword(user.UserPassword)
	query := `INSERT INTO user (
		user_id,
		user_nama,
		user_email,
		user_password,
		user_alamat,
		user_no_hp,
		user_tipe_id,
		is_user
		) VALUES (?,?,?,?,?,?,?,?)`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		fmt.Println(query)

		return err

	}
	defer stmt.Close()
	_, err = stmt.Exec(user.UserId, user.UserNama, user.UserEmail, user.UserPassword, user.UserAlamat, user.UserNoHp, user.UserTipeId, user.IsUser)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Get(userId string) (domain.UserResponseFormat, error) {
	var err error

	user := domain.UserResponseFormat{}

	query := `SELECT * FROM user where user_id=?  `

	// execute
	err = r.db.QueryRow(query, userId).Scan(&user.UserId, &user.UserNama, &user.UserEmail, &user.UserPassword, &user.UserAlamat, &user.UserNoHp, &user.UserTipeId, &user.IsUser, &user.CreatedAt)
	if err != nil {
		fmt.Println(err.Error())
		if err.Error() == "sql: no rows in result set" {
			return user, errors.New("user is not found")
		}
		return user, errors.New("failed get")
	}

	return user, nil
}

func (r *UserRepository) GetAll() ([]domain.UserResponseFormat, error) {

	userr := []domain.UserResponseFormat{}

	query := `SELECT * from user `

	rows1, err := r.db.Query(query)
	if err != nil {
		return userr, errors.New("failed get")
	}

	defer rows1.Close()
	for rows1.Next() {
		user := domain.UserResponseFormat{}
		err := rows1.Scan(&user.UserId, &user.UserNama, &user.UserEmail, &user.UserPassword, &user.UserAlamat, &user.UserNoHp, &user.UserTipeId, &user.IsUser, user.CreatedAt)

		if err != nil {
			return userr, errors.New("failed get")
		}
		userr = append(userr, user)
	}

	return userr, nil
}

func (r *UserRepository) GetAllDatatables() ([]domain.UserResponseFormatDatatables, int, error) {
	userr := []domain.UserResponseFormatDatatables{}

	res, err := r.GetCount()
	if err != nil {
		if err.Error() == "user is not found" {
			return userr, 0, nil
		}
		return userr, 0, errors.New("Internal Server Eroor")
	}
	if res == 0 {
		return userr, 0, errors.New("user is empty")
	}
	query := `SELECT u.user_id,
	u.user_nama,
	u.user_email,
	u.user_password,
	u.user_alamat,
	u.user_no_hp,
	u.user_tipe_id,
	u.is_user , 
	u.created_at,
	COALESCE((select user_tipe_nama from user_tipe where user_tipe_id=u.user_tipe_id), 'Guest') from
	 user as u `
	rows1, err := r.db.Query(query)
	if err != nil {
		return userr, 0, errors.New("failed get")
	}

	defer rows1.Close()
	for rows1.Next() {
		user := domain.UserResponseFormatDatatables{}
		err := rows1.Scan(&user.UserId, &user.UserNama, &user.UserEmail, &user.UserPassword, &user.UserAlamat, &user.UserNoHp, &user.UserTipeId, &user.IsUser, &user.CreatedAt, &user.UserTipeNama)

		if err != nil {
			fmt.Println(err.Error())

			return userr, 0, errors.New("failed get")
		}

		user.Action += `<a href="javascript:void(0)" data-toggle="tooltip"  data-id="` + user.UserId + `" data-original-title="Edit" class="edit btn btn-primary btn-sm editUser">Edit</a>`

		user.Action += `<a href="javascript:void(0)" data-toggle="tooltip"  data-id="` + user.UserId + `" data-original-title="Delete" class="btn btn-danger btn-sm deleteUser">Delete</a>`
		userr = append(userr, user)
	}

	return userr, res, nil
}

func (r *UserRepository) Update(user *domain.User) error {

	ifuserId, res, err := r.CheckEmail(user.UserEmail)
	if err != nil {
		return err
	} else if res {
		if ifuserId != user.UserId {
			return errors.New("user email already exist")
		}
	}

	query, params := r.QueryBuilder("update", user)
	result, err := r.db.Exec(query, params...)
	if err != nil {
		return errors.New("failed update")
	}
	count, _ := result.RowsAffected()
	if count < 1 {

		return errors.New("invalid user id")
	}

	return nil

}

func (r *UserRepository) Delete(userId string) (int64, error) {
	var err error

	query := `DELETE FROM user WHERE user_id =? `

	result, err := r.db.Exec(query, userId)
	if err != nil {
		return 0, err
	}

	count, _ := result.RowsAffected()
	if count < 1 {

		return 0, errors.New("user is not found")
	}
	return count, nil

}

func (r *UserRepository) QueryBuilder(scope string, user *domain.User) (string, []interface{}) {
	queries := []string{}
	params := make([]interface{}, 0)
	query := ""
	if scope == "update" {
		query += "UPDATE user SET "
	}

	if user.UserNama != "" {
		queries = append(queries, " user_nama=? ")
		params = append(params, user.UserNama)

	}
	if user.UserEmail != "" {
		queries = append(queries, " user_email=? ")
		params = append(params, user.UserEmail)

	}
	if user.UserAlamat != "" {
		queries = append(queries, " user_alamat=? ")
		params = append(params, user.UserAlamat)
	}

	if user.UserNoHp != "" {
		queries = append(queries, " user_no_hp=? ")
		params = append(params, user.UserNoHp)
	}
	if user.UserTipeId != "" {
		queries = append(queries, " user_tipe_id=? ")
		params = append(params, user.UserTipeId)
	}
	query += strings.Join(queries, ",")
	query += " WHERE user_id =?  "
	params = append(params, user.UserId)

	return query, params
}
func (r *UserRepository) CheckEmail(userEmail string) (string, bool, error) {
	var userId string
	query := `SELECT user_id AS userId FROM user `
	query += `WHERE user_email=? `
	fmt.Println(query, userEmail)

	err := r.db.QueryRow(query, userEmail).Scan(&userId)
	if err != nil {
		fmt.Println(query, err.Error())

		if err.Error() == "sql: no rows in result set" {
			return userId, false, nil
		}
		return userId, false, err
	}
	return userId, true, nil

}

// Get Count
func (r *UserRepository) GetCount() (int, error) {

	var count int

	// cipher := os.Getenv("CHIPER_MYSQL")

	query := "SELECT count(user_id) FROM user "

	err := r.db.QueryRow(query).Scan(&count)

	if err != nil {
		return 0, errors.New("internal server error")
	}
	if count == 0 {
		return 0, errors.New("user is not found")

	}
	return count, nil
}
