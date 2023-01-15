package fasilitas

import (
	"database/sql"
	"errors"
	"fmt"
	domain "rogerdev_golf/entities"

	"strings"
)

type FasilitasRepository struct {
	db *sql.DB
}

func NewFasilitasRepository(db *sql.DB) *FasilitasRepository {

	return &FasilitasRepository{db: db}
}
func (r *FasilitasRepository) Create(fasilitas *domain.Fasilitas) error {

	_, res, err := r.CheckName(fasilitas.FasilitasNama)
	if err != nil {
		return err
	} else if res {
		return errors.New("fasilitas name already exist")
	}
	query := `INSERT INTO fasilitas (
		fasilitas_id,
		fasilitas_nama,
		fasilitas_icon,
		fasilitas_image,
		fasilitas_body
		) VALUES (?,?,?,?,?)`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		fmt.Println(query)

		return err

	}
	defer stmt.Close()
	_, err = stmt.Exec(fasilitas.FasilitasId, fasilitas.FasilitasNama, fasilitas.FasilitasIcon, fasilitas.FasilitasImage, fasilitas.FasilitasBody)
	if err != nil {
		return err
	}

	return nil
}

func (r *FasilitasRepository) Get(fasilitasId string) (domain.FasilitasResponseFormat, error) {
	var err error

	fasilitas := domain.FasilitasResponseFormat{}

	query := `SELECT * FROM fasilitas where fasilitas_id=?  `

	// execute
	err = r.db.QueryRow(query, fasilitasId).Scan(&fasilitas.FasilitasId, &fasilitas.FasilitasNama, &fasilitas.FasilitasIcom, &fasilitas.FasilitasImage, &fasilitas.FasilitasBody)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return fasilitas, errors.New("fasilitas is not found")
		}
		return fasilitas, errors.New("failed get")
	}

	return fasilitas, nil
}

func (r *FasilitasRepository) GetAll() ([]domain.FasilitasResponseFormat, error) {

	fasilitass := []domain.FasilitasResponseFormat{}

	query := `SELECT * from fasilitas `

	rows1, err := r.db.Query(query)
	if err != nil {
		return fasilitass, errors.New("failed get")
	}

	defer rows1.Close()
	for rows1.Next() {
		fasilitas := domain.FasilitasResponseFormat{}
		err := rows1.Scan(&fasilitas.FasilitasId, &fasilitas.FasilitasNama, &fasilitas.FasilitasIcom, &fasilitas.FasilitasImage, &fasilitas.FasilitasBody)

		if err != nil {
			return fasilitass, errors.New("failed get")
		}
		fasilitass = append(fasilitass, fasilitas)
	}

	return fasilitass, nil
}
func (r *FasilitasRepository) GetAllDatatables() ([]domain.FasilitasResponseFormatDatatables, int, error) {
	fasilitass := []domain.FasilitasResponseFormatDatatables{}

	res, err := r.GetCount()
	if err != nil {
		return fasilitass, 0, errors.New("Internal Server Eroor")
	}
	if res == 0 {
		return fasilitass, 0, errors.New("fasilitas is empty")
	}
	query := `SELECT fasilitas_id,
	fasilitas_nama,
	fasilitas_icon,
	fasilitas_image,
	fasilitas_body
	 from
	  fasilitas `
	fmt.Println(query)
	rows1, err := r.db.Query(query)
	if err != nil {
		return fasilitass, 0, errors.New("failed get")
	}

	defer rows1.Close()
	for rows1.Next() {
		fasilitas := domain.FasilitasResponseFormatDatatables{}
		err := rows1.Scan(&fasilitas.FasilitasId, &fasilitas.FasilitasNama, &fasilitas.FasilitasIcon, &fasilitas.FasilitasImage, &fasilitas.FasilitasBody)

		if err != nil {

			return fasilitass, 0, errors.New("failed get")
		}

		fasilitas.Action += `<a href="javascript:void(0)" data-toggle="tooltip"  data-id="` + fasilitas.FasilitasId + `" data-original-title="Edit" class="edit btn btn-primary btn-sm editFasilitas">Edit</a>`

		fasilitas.Action += `<a href="javascript:void(0)" data-toggle="tooltip"  data-id="` + fasilitas.FasilitasId + `" data-original-title="Delete" class="btn btn-danger btn-sm deleteFasilitas">Delete</a>`
		fasilitass = append(fasilitass, fasilitas)
		fmt.Println(fasilitas, "EEEERRRRROOOOOO")

	}

	return fasilitass, res, nil
}

func (r *FasilitasRepository) Update(fasilitas *domain.Fasilitas) error {

	iffasilitasId, res, err := r.CheckName(fasilitas.FasilitasNama)
	if err != nil {
		return err
	} else if res {
		if iffasilitasId != fasilitas.FasilitasId {
			return errors.New("fasilitas name already exist")
		}
	}

	query, params := r.QueryBuilder("update", fasilitas)
	result, err := r.db.Exec(query, params...)
	if err != nil {
		return errors.New("failed update")
	}
	count, _ := result.RowsAffected()
	if count < 1 {

		return errors.New("invalid fasilitas id")
	}

	return nil

}

func (r *FasilitasRepository) Delete(fasilitasId string) (int64, error) {
	var err error

	query := `DELETE FROM fasilitas WHERE fasilitas_id =? `

	result, err := r.db.Exec(query, fasilitasId)
	if err != nil {
		return 0, err
	}

	count, _ := result.RowsAffected()
	if count < 1 {

		return 0, errors.New("fasilitas is not found")
	}
	return count, nil

}

func (r *FasilitasRepository) QueryBuilder(scope string, fasilitas *domain.Fasilitas) (string, []interface{}) {
	queries := []string{}
	params := make([]interface{}, 0)
	query := ""
	if scope == "update" {
		query += "UPDATE fasilitas SET "
	}

	if fasilitas.FasilitasNama != "" {
		queries = append(queries, " fasilitas_nama=? ")
		params = append(params, fasilitas.FasilitasNama)

	}
	if fasilitas.FasilitasIcon != "" {
		queries = append(queries, " fasilitas_icon=? ")
		params = append(params, fasilitas.FasilitasIcon)

	}
	if fasilitas.FasilitasImage != "" {
		queries = append(queries, " fasilitas_image=? ")
		params = append(params, fasilitas.FasilitasImage)
	}

	if fasilitas.FasilitasBody != "" {
		queries = append(queries, " fasilitas_body=? ")
		params = append(params, fasilitas.FasilitasBody)
	}

	query += strings.Join(queries, ",")
	query += " WHERE fasilitas_id =?  "
	params = append(params, fasilitas.FasilitasId)

	return query, params
}
func (r *FasilitasRepository) CheckName(fasilitasName string) (string, bool, error) {
	var fasilitasId string
	query := `SELECT fasilitas_id AS fasilitasId FROM fasilitas `
	query += `WHERE fasilitas_nama=? `

	err := r.db.QueryRow(query, fasilitasName).Scan(&fasilitasId)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return fasilitasId, false, nil
		}
		return fasilitasId, false, err
	}
	return fasilitasId, true, nil

}

// Get Count
func (r *FasilitasRepository) GetCount() (int, error) {

	var count int

	// cipher := os.Getenv("CHIPER_MYSQL")

	query := "SELECT count(fasilitas_id) FROM fasilitas "

	err := r.db.QueryRow(query).Scan(&count)

	if err != nil {
		return 0, errors.New("internal server error")
	}
	if count == 0 {
		return 0, errors.New("fasilitas is not found")

	}
	return count, nil
}
