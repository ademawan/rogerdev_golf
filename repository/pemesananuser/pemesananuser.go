package pemesananuser

import (
	"database/sql"
	"errors"
	"fmt"
	domain "rogerdev_golf/entities"

	"strings"
)

type PemesananUserRepository struct {
	db *sql.DB
}

func NewPemesananUserRepository(db *sql.DB) *PemesananUserRepository {

	return &PemesananUserRepository{db: db}
}
func (r *PemesananUserRepository) Create(pemesanan *domain.PemesananUser) error {
	pemesanan.SecondPlayer = "0"
	pemesanan.PemesananStatus = "menunggu konfirmasi"
	query := `INSERT INTO pemesanan (
		pemesanan_id,
		pemesanan_nama,
		pemesanan_date,
		pemesanan_time,
		pemesanan_email,
		pemesanan_no_hp,
		pemesanan_status,
		pemesanan_alamat,
			first_player,
			second_player,
			third_player,
			fourth_player,
			user_tipe_id1,
			user_tipe_id2,
			user_tipe_id3,
			user_tipe_id4,
			user_id,
			status_pembayaran
		) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		fmt.Println(query)

		return err

	}
	defer stmt.Close()
	_, err = stmt.Exec(pemesanan.PemesananId, pemesanan.PemesananNama, pemesanan.PemesananDate, pemesanan.PemesananTime, pemesanan.PemesananEmail, pemesanan.PemesananNoHp, pemesanan.PemesananStatus, pemesanan.PemesananAlamat, pemesanan.FirstPlayer, pemesanan.SecondPlayer, pemesanan.ThirdPlayer, pemesanan.FourthPlayer, pemesanan.UserTipeId1, pemesanan.UserTipeId2, pemesanan.UserTipeId3, pemesanan.UserTipeId4, pemesanan.UserId, pemesanan.StatusPembayaran)
	if err != nil {
		return err
	}

	return nil
}

func (r *PemesananUserRepository) Get(pemesananId string) (domain.PemesananUserResponseFormat, error) {
	var err error

	pemesanan := domain.PemesananUserResponseFormat{}

	query := `SELECT * FROM pemesanan where pemesanan_id=?  `

	// execute
	err = r.db.QueryRow(query, pemesananId).Scan(&pemesanan.PemesananId, &pemesanan.PemesananNama, &pemesanan.PemesananDate, &pemesanan.PemesananTime, &pemesanan.PemesananEmail, &pemesanan.PemesananNoHp, &pemesanan.PemesananStatus, &pemesanan.PemesananAlamat, &pemesanan.CreatedAt, &pemesanan.FirstPlayer, &pemesanan.SecondPlayer, &pemesanan.ThirdPlayer, &pemesanan.FourthPlayer, &pemesanan.UserTipeId1, &pemesanan.UserTipeId2, &pemesanan.UserTipeId3, &pemesanan.UserTipeId4, &pemesanan.UserId, &pemesanan.StatusPembayaran)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return pemesanan, errors.New("pemesanan is not found")
		}
		return pemesanan, errors.New("failed get")
	}

	return pemesanan, nil
}

func (r *PemesananUserRepository) GetAll() ([]domain.PemesananUserResponseFormat, error) {

	pemesanann := []domain.PemesananUserResponseFormat{}

	query := `SELECT * from pemesanan `

	rows1, err := r.db.Query(query)
	if err != nil {
		return pemesanann, errors.New("failed get")
	}

	defer rows1.Close()
	for rows1.Next() {
		pemesanan := domain.PemesananUserResponseFormat{}
		err := rows1.Scan(&pemesanan.PemesananId, &pemesanan.PemesananNama, &pemesanan.PemesananDate, &pemesanan.PemesananTime, &pemesanan.PemesananEmail, &pemesanan.PemesananNoHp, &pemesanan.PemesananStatus, &pemesanan.PemesananAlamat, &pemesanan.CreatedAt, &pemesanan.FirstPlayer, &pemesanan.SecondPlayer, &pemesanan.ThirdPlayer, &pemesanan.FourthPlayer, &pemesanan.UserTipeId1, &pemesanan.UserTipeId2, &pemesanan.UserTipeId3, &pemesanan.UserTipeId4, &pemesanan.UserId, &pemesanan.StatusPembayaran)

		if err != nil {
			return pemesanann, errors.New("failed get")
		}
		pemesanann = append(pemesanann, pemesanan)
	}

	return pemesanann, nil
}

func (r *PemesananUserRepository) Update(pemesanan *domain.PemesananUser) error {

	query, params := r.QueryBuilder("update", pemesanan)
	result, err := r.db.Exec(query, params...)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("failed update")
	}
	count, _ := result.RowsAffected()
	if count < 1 {

		return errors.New("invalid pemesanan id")
	}

	return nil

}

func (r *PemesananUserRepository) Delete(pemesananId string) (int64, error) {
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

func (r *PemesananUserRepository) QueryBuilder(scope string, pemesanan *domain.PemesananUser) (string, []interface{}) {
	queries := []string{}
	params := make([]interface{}, 0)
	query := ""
	if scope == "update" {
		query += "UPDATE pemesanan SET "
	}

	if pemesanan.PemesananNama != "" {
		queries = append(queries, " pemesanan_nama=? ")
		params = append(params, pemesanan.PemesananNama)

	}
	if pemesanan.PemesananDate != "" {
		queries = append(queries, " pemesanan_date=? ")
		params = append(params, pemesanan.PemesananDate)

	}
	if pemesanan.PemesananTime != "" {
		queries = append(queries, " pemesanan_time=? ")
		params = append(params, pemesanan.PemesananTime)
	}

	if pemesanan.PemesananEmail != "" {
		queries = append(queries, " pemesanan_email=? ")
		params = append(params, pemesanan.PemesananEmail)
	}

	if pemesanan.PemesananNoHp != "" {
		queries = append(queries, " pemesanan_no_hp=? ")
		params = append(params, pemesanan.PemesananNoHp)
	}

	if pemesanan.PemesananStatus != "" {
		queries = append(queries, " pemesanan_status=? ")
		params = append(params, pemesanan.PemesananStatus)
	}
	if pemesanan.PemesananAlamat != "" {
		queries = append(queries, " pemesanan_alamat=? ")
		params = append(params, pemesanan.PemesananAlamat)
	}
	if pemesanan.StatusPembayaran != "" {
		queries = append(queries, " status_pembayaran=? ")
		params = append(params, pemesanan.StatusPembayaran)
	}

	query += strings.Join(queries, ",")
	query += " WHERE pemesanan_id =?  "

	params = append(params, pemesanan.PemesananId)

	return query, params
}
func (r *PemesananUserRepository) CheckName(pemesanannName string) (string, bool, error) {
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
func (r *PemesananUserRepository) GetAllDatatables() ([]domain.PemesananUserResponseFormatDatatables, int, error) {
	pemesanann := []domain.PemesananUserResponseFormatDatatables{}

	res, err := r.GetCount()
	if err != nil {
		if err.Error() == "pemesanan is empty" {
			return pemesanann, 0, nil
		}
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
	created_at,
		first_player,
		second_player,
		third_player,
		fourth_player,
		user_tipe_id1,
		user_tipe_id2,
		user_tipe_id3,
		user_tipe_id4,
		user_id,
		status_pembayaran
	 from
	 pemesanan `
	rows1, err := r.db.Query(query)
	if err != nil {

		return pemesanann, 0, errors.New("failed get")
	}

	defer rows1.Close()
	for rows1.Next() {
		pemesanan := domain.PemesananUserResponseFormatDatatables{}
		err := rows1.Scan(&pemesanan.PemesananId, &pemesanan.PemesananNama, &pemesanan.PemesananDate, &pemesanan.PemesananTime, &pemesanan.PemesananEmail, &pemesanan.PemesananNoHp, &pemesanan.PemesananStatus, &pemesanan.PemesananAlamat, &pemesanan.CreatedAt, &pemesanan.FirstPlayer, &pemesanan.SecondPlayer, &pemesanan.ThirdPlayer, &pemesanan.FourthPlayer, &pemesanan.UserTipeId1, &pemesanan.UserTipeId2, &pemesanan.UserTipeId3, &pemesanan.UserTipeId4, &pemesanan.UserId, &pemesanan.StatusPembayaran)

		if err != nil {
			fmt.Println(err.Error())

			return pemesanann, 0, errors.New("failed get")
		}
		if pemesanan.StatusPembayaran == "1" {
			pemesanan.StatusPembayaran = "Sudah Dibayar"
		} else {
			if pemesanan.PemesananStatus == "menunggu konfirmasi" {
				pemesanan.StatusPembayaran = "menunggu konfirmasi"
			} else {
				pemesanan.StatusPembayaran = "Belum Dibayar"
				pemesanan.Action += `<a href="javascript:void(0)" data-toggle="tooltip"  data-id="` + pemesanan.PemesananId + `" data-original-title="Bayar" class="bayar btn btn-primary btn-sm editPemesanan">Bayar</a>`
			}

		}

		// pemesanan.Action += `<a href="javascript:void(0)" data-toggle="tooltip"  data-id="` + pemesanan.PemesananId + `" data-original-title="Edit" class="edit btn btn-primary btn-sm editPemesanan">Edit</a>`

		// pemesanan.Action += `<a href="javascript:void(0)" data-toggle="tooltip"  data-id="` + pemesanan.PemesananId + `" data-original-title="Delete" class="btn btn-danger btn-sm deletePemesanan">Delete</a>`

		pemesanann = append(pemesanann, pemesanan)
	}

	return pemesanann, res, nil
}

// Get Count
func (r *PemesananUserRepository) GetCount() (int, error) {

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
