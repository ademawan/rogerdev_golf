package pemesanan

import (
	"database/sql"
	"errors"
	"fmt"
	domain "rogerdev_golf/entities"

	"strings"
)

type PemesananRepository struct {
	db *sql.DB
}

func NewPemesananRepository(db *sql.DB) *PemesananRepository {

	return &PemesananRepository{db: db}
}
func (r *PemesananRepository) Create(pemesanan *domain.Pemesanan) error {

	query := `INSERT INTO pemesanan (
		pemesanan_id,
		pemesanan_nama,
		pemesanan_date,
		pemesanan_time,
		pemesanan_email,
		pemesanan_no_hp,
		pemesanan_status,
		pemesanan_alamat,
		created_at
		) VALUES (?,?,?,?,?,?,?,?,?)`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		fmt.Println(query)

		return err

	}
	defer stmt.Close()
	_, err = stmt.Exec(pemesanan.PemesananId, pemesanan.PemesananNama, pemesanan.PemesananDate, pemesanan.PemesananTime, pemesanan.PemesananEmail, pemesanan.PemesananNoHp, pemesanan.PemesananStatus, pemesanan.PemesananAlamat, pemesanan.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *PemesananRepository) Get(pemesananId string) (domain.PemesananResponseFormat, error) {
	var err error

	pemesanan := domain.PemesananResponseFormat{}

	query := `SELECT * FROM pemesanan where pemesanan_id=?  `

	// execute
	err = r.db.QueryRow(query, pemesananId).Scan(&pemesanan.PemesananId, &pemesanan.PemesananNama, &pemesanan.PemesananDate, &pemesanan.PemesananTime, &pemesanan.PemesananEmail, &pemesanan.PemesananEmail, &pemesanan.PemesananNoHp, &pemesanan.PemesananStatus, &pemesanan.PemesananAlamat, &pemesanan.CreatedAt)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return pemesanan, errors.New("pemesanan is not found")
		}
		return pemesanan, errors.New("failed get")
	}

	return pemesanan, nil
}

func (r *PemesananRepository) GetAll() ([]domain.PemesananResponseFormat, error) {

	pemesanann := []domain.PemesananResponseFormat{}

	query := `SELECT * from pemesanan `

	rows1, err := r.db.Query(query)
	if err != nil {
		return pemesanann, errors.New("failed get")
	}

	defer rows1.Close()
	for rows1.Next() {
		pemesanan := domain.PemesananResponseFormat{}
		err := rows1.Scan(&pemesanan.PemesananId, &pemesanan.PemesananNama, &pemesanan.PemesananDate, &pemesanan.PemesananTime, &pemesanan.PemesananEmail, &pemesanan.PemesananEmail, &pemesanan.PemesananNoHp, &pemesanan.PemesananStatus, &pemesanan.PemesananAlamat, &pemesanan.CreatedAt)

		if err != nil {
			return pemesanann, errors.New("failed get")
		}
		pemesanann = append(pemesanann, pemesanan)
	}

	return pemesanann, nil
}

func (r *PemesananRepository) Update(pemesanan *domain.Pemesanan) error {

	query, params := r.QueryBuilder("update", pemesanan)
	result, err := r.db.Exec(query, params...)
	if err != nil {
		return errors.New("failed update")
	}
	count, _ := result.RowsAffected()
	if count < 1 {

		return errors.New("invalid pemesanan id")
	}

	return nil

}

func (r *PemesananRepository) Delete(pemesananId string) (int64, error) {
	var err error

	query := `DELETE FROM pemesanan WHERE pemesanan_id =? `

	result, err := r.db.Exec(query, pemesananId)
	if err != nil {
		return 0, err
	}

	count, _ := result.RowsAffected()
	if count < 1 {

		return 0, errors.New("pemesanan is not found")
	}
	return count, nil

}

func (r *PemesananRepository) QueryBuilder(scope string, fasilitas *domain.Pemesanan) (string, []interface{}) {
	queries := []string{}
	params := make([]interface{}, 0)
	query := ""
	if scope == "update" {
		query += "UPDATE fasilitas SET "
	}

	if fasilitas.PemesananNama != "" {
		queries = append(queries, " pemesanan_nama=? ")
		params = append(params, fasilitas.PemesananNama)

	}
	if fasilitas.PemesananDate != "" {
		queries = append(queries, " pemesanan_date=? ")
		params = append(params, fasilitas.PemesananDate)

	}
	if fasilitas.PemesananTime != "" {
		queries = append(queries, " pemesanan_time=? ")
		params = append(params, fasilitas.PemesananTime)
	}

	if fasilitas.PemesananEmail != "" {
		queries = append(queries, " pemesanan_email=? ")
		params = append(params, fasilitas.PemesananEmail)
	}

	if fasilitas.PemesananNoHp != "" {
		queries = append(queries, " pemesanan_no_hp=? ")
		params = append(params, fasilitas.PemesananNoHp)
	}

	if fasilitas.PemesananStatus != "" {
		queries = append(queries, " pemesanan_status=? ")
		params = append(params, fasilitas.PemesananStatus)
	}
	if fasilitas.PemesananAlamat != "" {
		queries = append(queries, " pemesanan_alamat=? ")
		params = append(params, fasilitas.PemesananAlamat)
	}

	query += strings.Join(queries, ",")
	query += " WHERE pemesanan_id =?  "
	params = append(params, fasilitas.PemesananId)

	return query, params
}
func (r *PemesananRepository) CheckName(pemesanannName string) (string, bool, error) {
	var pemesananId string
	query := `SELECT pemesanan_id AS pemesananId FROM pemesanan `
	query += `WHERE pemesanan_nama=? `

	err := r.db.QueryRow(query, pemesanannName).Scan(&pemesananId)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return pemesananId, false, nil
		}
		return pemesananId, false, err
	}
	return pemesananId, true, nil

}
func (r *PemesananRepository) GetAllDatatables() ([]domain.PemesananResponseFormatDatatables, int, error) {
	pemesanann := []domain.PemesananResponseFormatDatatables{}

	res, err := r.GetCount()
	if err != nil {
		return pemesanann, 0, errors.New("Internal Server Eroor")
	}
	if res == 0 {
		return pemesanann, 0, errors.New("user is empty")
	}
	query := `SELECT pemesanan_id,
	pemesanan_nama,
	pemesanan_date,
	pemesanan_time,
	pemesanan_email,
	pemesanan_no_hp,
	pemesanan_status,
	pemesanan_alamat, 
	created_at
	 from
	 pemesanan `
	rows1, err := r.db.Query(query)
	if err != nil {

		return pemesanann, 0, errors.New("failed get")
	}

	defer rows1.Close()
	for rows1.Next() {
		pemesanan := domain.PemesananResponseFormatDatatables{}
		err := rows1.Scan(&pemesanan.PemesananId, &pemesanan.PemesananNama, &pemesanan.PemesananDate, &pemesanan.PemesananTime, &pemesanan.PemesananEmail, &pemesanan.PemesananNoHp, &pemesanan.PemesananStatus, &pemesanan.PemesananAlamat, &pemesanan.CreatedAt)

		if err != nil {

			return pemesanann, 0, errors.New("failed get")
		}

		pemesanan.Action += `<a href="javascript:void(0)" data-toggle="tooltip"  data-id="` + pemesanan.PemesananId + `" data-original-title="Edit" class="edit btn btn-primary btn-sm editPemesanan">Edit</a>`

		pemesanan.Action += `<a href="javascript:void(0)" data-toggle="tooltip"  data-id="` + pemesanan.PemesananId + `" data-original-title="Delete" class="btn btn-danger btn-sm deletePemesanan">Delete</a>`
		pemesanann = append(pemesanann, pemesanan)
	}

	return pemesanann, res, nil
}

// Get Count
func (r *PemesananRepository) GetCount() (int, error) {

	var count int

	// cipher := os.Getenv("CHIPER_MYSQL")

	query := "SELECT count(pemesanan_id) FROM pemesanan "

	err := r.db.QueryRow(query).Scan(&count)

	if err != nil {

		return 0, errors.New("internal server error")
	}
	if count == 0 {

		return 0, errors.New("pemesanan is empty")

	}
	return count, nil
}
