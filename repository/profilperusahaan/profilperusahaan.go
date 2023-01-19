package profilperusahaan

import (
	"database/sql"
	"errors"
	"fmt"
	domain "rogerdev_golf/entities"

	"strings"
)

type ProfilPerusahaanRepository struct {
	db *sql.DB
}

func NewProfilPerusahaanRepository(db *sql.DB) *ProfilPerusahaanRepository {

	return &ProfilPerusahaanRepository{db: db}
}
func (r *ProfilPerusahaanRepository) Create(profilPerusahaan domain.ProfilPerusahaan) error {

	query := `INSERT INTO profil_perusahaan (
		profil_perusahaan_id,
		profil_perusahaan_nama,
		profil_perusahaan_email,
		profil_perusahaan_alamat,
		profil_perusahaan_no_hp,
		profil_perusahaan_deskripsi,
		profil_perusahaan_image
		) VALUES (?,?,?,?,?,?,?)`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		fmt.Println(query)

		return err

	}
	defer stmt.Close()
	_, err = stmt.Exec(profilPerusahaan.ProfilPerusahaanId, profilPerusahaan.ProfilPerusahaanNama, profilPerusahaan.ProfilPerusahaanEmail, profilPerusahaan.ProfilPerusahaanAlamat, profilPerusahaan.ProfilPerusahaanNoHp, profilPerusahaan.ProfilPerusahaanDeskripsi, profilPerusahaan.ProfilPerusahaanImage)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProfilPerusahaanRepository) Get(profilPerusahaanId string) (domain.ProfilPerusahaanResponseFormat, error) {
	var err error

	profilPerusahaan := domain.ProfilPerusahaanResponseFormat{}

	query := `SELECT * FROM profil_perusahaan where profil_perusahaan_id=?  `

	// execute
	err = r.db.QueryRow(query, profilPerusahaanId).Scan(&profilPerusahaan.ProfilPerusahaanId, &profilPerusahaan.ProfilPerusahaanNama, &profilPerusahaan.ProfilPerusahaanEmail, &profilPerusahaan.ProfilPerusahaanAlamat, &profilPerusahaan.ProfilPerusahaanNoHp, &profilPerusahaan.ProfilPerusahaanDeskripsi, &profilPerusahaan.ProfilPerusahaanImage)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return profilPerusahaan, errors.New("profil perusahaan is not found")
		}
		return profilPerusahaan, errors.New("failed get")
	}

	return profilPerusahaan, nil
}

func (r *ProfilPerusahaanRepository) GetProfilPerusahaan() (domain.ProfilPerusahaanResponseFormat, error) {
	var err error

	profilPerusahaan := domain.ProfilPerusahaanResponseFormat{}
	query := `SELECT * from profil_perusahaan `

	rows1, err := r.db.Query(query)
	if err != nil {
		return domain.ProfilPerusahaanResponseFormat{}, errors.New("failed get")
	}

	defer rows1.Close()
	for rows1.Next() {
		err := rows1.Scan(&profilPerusahaan.ProfilPerusahaanId, &profilPerusahaan.ProfilPerusahaanNama, &profilPerusahaan.ProfilPerusahaanEmail, &profilPerusahaan.ProfilPerusahaanAlamat, &profilPerusahaan.ProfilPerusahaanNoHp, &profilPerusahaan.ProfilPerusahaanDeskripsi, &profilPerusahaan.ProfilPerusahaanImage)

		if err != nil {
			return domain.ProfilPerusahaanResponseFormat{}, errors.New("failed get")
		}
	}

	return profilPerusahaan, nil
}

func (r *ProfilPerusahaanRepository) GetAll() ([]domain.ProfilPerusahaanResponseFormat, error) {

	perusahaanPrusahaann := []domain.ProfilPerusahaanResponseFormat{}

	query := `SELECT * from profil_perusahaan `

	rows1, err := r.db.Query(query)
	if err != nil {
		return perusahaanPrusahaann, errors.New("failed get")
	}

	defer rows1.Close()
	for rows1.Next() {
		profilPerusahaan := domain.ProfilPerusahaanResponseFormat{}
		err := rows1.Scan(&profilPerusahaan.ProfilPerusahaanId, &profilPerusahaan.ProfilPerusahaanNama, &profilPerusahaan.ProfilPerusahaanEmail, &profilPerusahaan.ProfilPerusahaanAlamat, &profilPerusahaan.ProfilPerusahaanNoHp, &profilPerusahaan.ProfilPerusahaanDeskripsi, &profilPerusahaan.ProfilPerusahaanImage)

		if err != nil {
			return perusahaanPrusahaann, errors.New("failed get")
		}
		perusahaanPrusahaann = append(perusahaanPrusahaann, profilPerusahaan)
	}

	return perusahaanPrusahaann, nil
}

func (r *ProfilPerusahaanRepository) GetAllDatatables() ([]domain.ProfilPerusahaanResponseFormatDatatables, int, error) {
	profilPerusahaann := []domain.ProfilPerusahaanResponseFormatDatatables{}

	res, err := r.GetCount()
	if err != nil {
		if err.Error() == "profil perusahaan is not found" {
			return profilPerusahaann, 0, nil
		}
		return profilPerusahaann, 0, errors.New("Internal Server Eroor")
	}
	if res == 0 {
		return profilPerusahaann, 0, errors.New("profil_perusahaan is empty")
	}
	query := `SELECT profil_perusahaan_id,
	profil_perusahaan_nama,
	profil_perusahaan_email,
	profil_perusahaan_alamat,
	profil_perusahaan_no_hp,
	profil_perusahaan_deskripsi,
	profil_perusahaan_image from
	 profil_perusahaan `
	rows1, err := r.db.Query(query)
	if err != nil {
		return profilPerusahaann, 0, errors.New("failed get")
	}

	defer rows1.Close()
	for rows1.Next() {
		profilPerusahaan := domain.ProfilPerusahaanResponseFormatDatatables{}
		err := rows1.Scan(&profilPerusahaan.ProfilPerusahaanId, &profilPerusahaan.ProfilPerusahaanNama, &profilPerusahaan.ProfilPerusahaanEmail, &profilPerusahaan.ProfilPerusahaanAlamat, &profilPerusahaan.ProfilPerusahaanNoHp, &profilPerusahaan.ProfilPerusahaanDeskripsi, &profilPerusahaan.ProfilPerusahaanImage)

		if err != nil {
			fmt.Println(err.Error())

			return profilPerusahaann, 0, errors.New("failed get")
		}
		profilPerusahaann = append(profilPerusahaann, profilPerusahaan)

	}

	return profilPerusahaann, res, nil
}

func (r *ProfilPerusahaanRepository) Update(profilPerusahaan *domain.ProfilPerusahaan) error {

	query, params := r.QueryBuilder("update", profilPerusahaan)
	_, err := r.db.Exec(query, params...)
	if err != nil {
		return errors.New("failed update")
	}

	return nil

}

func (r *ProfilPerusahaanRepository) Delete(profilPerusahaanId string) (int64, error) {
	var err error

	query := `DELETE FROM profil_perusahaan WHERE profil_perusahaan_id =? `

	result, err := r.db.Exec(query, profilPerusahaanId)
	if err != nil {
		return 0, err
	}

	count, _ := result.RowsAffected()
	if count < 1 {

		return 0, errors.New("profil_perusahaan is not found")
	}
	return count, nil

}

func (r *ProfilPerusahaanRepository) QueryBuilder(scope string, profilPerusahaan *domain.ProfilPerusahaan) (string, []interface{}) {
	queries := []string{}
	params := make([]interface{}, 0)
	query := ""
	if scope == "update" {
		query += "UPDATE profil_perusahaan SET "
	}

	if profilPerusahaan.ProfilPerusahaanNama != "" {
		queries = append(queries, " profil_perusahaan_nama=? ")
		params = append(params, profilPerusahaan.ProfilPerusahaanNama)

	}
	if profilPerusahaan.ProfilPerusahaanEmail != "" {
		queries = append(queries, " profil_perusahaan_email=? ")
		params = append(params, profilPerusahaan.ProfilPerusahaanEmail)

	}
	if profilPerusahaan.ProfilPerusahaanAlamat != "" {
		queries = append(queries, " profil_perusahaan_alamat=? ")
		params = append(params, profilPerusahaan.ProfilPerusahaanAlamat)
	}

	if profilPerusahaan.ProfilPerusahaanNoHp != "" {
		queries = append(queries, " profil_perusahaan_no_hp=? ")
		params = append(params, profilPerusahaan.ProfilPerusahaanNoHp)
	}
	if profilPerusahaan.ProfilPerusahaanDeskripsi != "" {
		queries = append(queries, " profil_perusahaan_deskripsi=? ")
		params = append(params, profilPerusahaan.ProfilPerusahaanDeskripsi)
	}
	if profilPerusahaan.ProfilPerusahaanImage != "" {
		queries = append(queries, " profil_perusahaan_image=? ")
		params = append(params, profilPerusahaan.ProfilPerusahaanImage)
	}
	query += strings.Join(queries, ",")
	query += " WHERE profil_perusahaan_id =?  "
	params = append(params, profilPerusahaan.ProfilPerusahaanId)

	return query, params
}

// Get Count
func (r *ProfilPerusahaanRepository) GetCount() (int, error) {

	var count int

	// cipher := os.Getenv("CHIPER_MYSQL")

	query := "SELECT count(profil_perusahaan_id) FROM profil_perusahaan "

	err := r.db.QueryRow(query).Scan(&count)

	if err != nil {
		return 0, errors.New("internal server error")
	}
	if count == 0 {
		return 0, errors.New("profil perusahaan is not found")

	}
	return count, nil
}
