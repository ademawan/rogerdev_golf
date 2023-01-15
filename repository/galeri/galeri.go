package galeri

import (
	"database/sql"
	"errors"
	"fmt"
	domain "rogerdev_golf/entities"

	"strings"
)

type GaleriRepository struct {
	db *sql.DB
}

func NewGaleriRepository(db *sql.DB) *GaleriRepository {

	return &GaleriRepository{db: db}
}
func (r *GaleriRepository) Create(galeri *domain.Galeri) error {

	_, res, err := r.CheckName(galeri.GaleriNama)
	if err != nil {
		return err
	} else if res {
		return errors.New("galeri name already exist")
	}
	query := `INSERT INTO galeri (
		galeri_id,
		galeri_nama,
		galeri_image
		) VALUES (?,?,?)`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		fmt.Println(query)

		return err

	}
	defer stmt.Close()
	_, err = stmt.Exec(galeri.GaleriId, galeri.GaleriNama, galeri.GaleriImage)
	if err != nil {
		return err
	}

	return nil
}

func (r *GaleriRepository) Get(GaleriId string) (domain.GaleriResponseFormat, error) {
	var err error

	galeri := domain.GaleriResponseFormat{}

	query := `SELECT * FROM galeri where galeri_id=?  `

	// execute
	err = r.db.QueryRow(query, GaleriId).Scan(&galeri.GaleriId, &galeri.GaleriNama, &galeri.GaleriImage)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return galeri, errors.New("galeri is not found")
		}
		return galeri, errors.New("failed get")
	}

	return galeri, nil
}

func (r *GaleriRepository) GetAll() ([]domain.GaleriResponseFormat, error) {

	galerii := []domain.GaleriResponseFormat{}

	query := `SELECT * from galeri `

	rows1, err := r.db.Query(query)
	if err != nil {
		return galerii, errors.New("failed get")
	}

	defer rows1.Close()
	for rows1.Next() {
		galeri := domain.GaleriResponseFormat{}
		err := rows1.Scan(&galeri.GaleriId, &galeri.GaleriNama, &galeri.GaleriImage)

		if err != nil {
			return galerii, errors.New("failed get")
		}
		galerii = append(galerii, galeri)
	}

	return galerii, nil
}
func (r *GaleriRepository) GetAllDatatables() ([]domain.GaleriResponseFormatDatatables, int, error) {
	galerii := []domain.GaleriResponseFormatDatatables{}

	res, err := r.GetCount()
	if err != nil {
		return galerii, 0, errors.New("Internal Server Eroor")
	}
	if res == 0 {
		return galerii, 0, errors.New("user is empty")
	}
	query := `SELECT galeri_id,
	galeri_nama,
	galeri_image from
	 galeri `
	rows1, err := r.db.Query(query)
	if err != nil {
		return galerii, 0, errors.New("failed get")
	}

	defer rows1.Close()
	for rows1.Next() {
		galeri := domain.GaleriResponseFormatDatatables{}
		err := rows1.Scan(&galeri.GaleriId, &galeri.GaleriNama, &galeri.GaleriImage)

		if err != nil {

			return galerii, 0, errors.New("failed get")
		}

		galeri.Action += `<a href="javascript:void(0)" data-toggle="tooltip"  data-id="` + galeri.GaleriId + `" data-original-title="Edit" class="edit btn btn-primary btn-sm editUser">Edit</a>`

		galeri.Action += `<a href="javascript:void(0)" data-toggle="tooltip"  data-id="` + galeri.GaleriId + `" data-original-title="Delete" class="btn btn-danger btn-sm deleteUser">Delete</a>`
		galerii = append(galerii, galeri)
	}

	return galerii, res, nil
}

func (r *GaleriRepository) Update(galeri *domain.Galeri) error {

	ifgaleriId, res, err := r.CheckName(galeri.GaleriNama)
	if err != nil {
		return err
	} else if res {
		if ifgaleriId != galeri.GaleriId {
			return errors.New("galeri name already exist")
		}
	}

	query, params := r.QueryBuilder("update", galeri)
	result, err := r.db.Exec(query, params...)
	if err != nil {
		return errors.New("failed update")
	}
	count, _ := result.RowsAffected()
	if count < 1 {

		return errors.New("invalid galeri id")
	}

	return nil

}

func (r *GaleriRepository) Delete(GaleriId string) (int64, error) {
	var err error

	query := `DELETE FROM galeri WHERE galeri_id =? `

	result, err := r.db.Exec(query, GaleriId)
	if err != nil {
		return 0, err
	}

	count, _ := result.RowsAffected()
	if count < 1 {

		return 0, errors.New("galeri is not found")
	}
	return count, nil

}

func (r *GaleriRepository) QueryBuilder(scope string, galeri *domain.Galeri) (string, []interface{}) {
	queries := []string{}
	params := make([]interface{}, 0)
	query := ""
	if scope == "update" {
		query += "UPDATE galeri SET "
	}

	if galeri.GaleriNama != "" {
		queries = append(queries, " galeri_nama=? ")
		params = append(params, galeri.GaleriNama)

	}
	if galeri.GaleriImage != "" {
		queries = append(queries, " galeri_image=? ")
		params = append(params, galeri.GaleriImage)

	}
	query += strings.Join(queries, ",")
	query += " WHERE galeri_id =?  "
	params = append(params, galeri.GaleriId)

	return query, params
}
func (r *GaleriRepository) CheckName(fasilitasName string) (string, bool, error) {
	var galeriId string
	query := `SELECT galeri_id AS galeriId FROM galeri `
	query += `WHERE AND galeri_nama=? `

	err := r.db.QueryRow(query, fasilitasName).Scan(&galeriId)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return galeriId, false, nil
		}
		return galeriId, false, err
	}
	return galeriId, true, nil

}

// Get Count
func (r *GaleriRepository) GetCount() (int, error) {

	var count int

	// cipher := os.Getenv("CHIPER_MYSQL")

	query := "SELECT count(galeri_id) FROM galeri "

	err := r.db.QueryRow(query).Scan(&count)

	if err != nil {
		return 0, errors.New("internal server error")
	}
	if count == 0 {
		return 0, errors.New("user is not found")

	}
	return count, nil
}
