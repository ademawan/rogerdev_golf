package entities

type PemesananUser struct {
	PemesananId      string
	PemesananNama    string
	PemesananDate    string
	PemesananTime    string
	PemesananEmail   string
	PemesananNoHp    string
	PemesananStatus  string
	PemesananAlamat  string
	CreatedAt        string
	FirstPlayer      string
	SecondPlayer     string
	ThirdPlayer      string
	FourthPlayer     string
	UserTipeId1      string
	UserTipeId2      string
	UserTipeId3      string
	UserTipeId4      string
	UserId           string
	StatusPembayaran string
}

type PemesananUserRequestCreateFormat struct {
	PemesananNama    string `json:"pemesanan_nama" form:"pemesanan_nama"`
	PemesananDate    string `json:"pemesanan_date" form:"pemesanan_date"`
	PemesananTime    string `json:"pemesanan_time" form:"pemesanan_time"`
	PemesananEmail   string `json:"pemesanan_email" form:"pemesanan_email"`
	PemesananNoHp    string `json:"pemesanan_no_hp" form:"pemesanan_no_hp"`
	PemesananStatus  string `json:"pemesanan_status" form:"pemesanan_status"`
	PemesananAlamat  string `json:"pemesanan_alamat" form:"pemesanan_alamat"`
	FirstPlayer      string `json:"first_player" form:"first_player"`
	SecondPlayer     string `json:"second_player" form:"second_player"`
	ThirdPlayer      string `json:"third_player" form:"third_player"`
	FourthPlayer     string `json:"fourth_player" form:"fourth_player"`
	UserTipeId1      string `json:"user_tipe_id1" form:"user_tipe_id1"`
	UserTipeId2      string `json:"user_tipe_id2" form:"user_tipe_id2"`
	UserTipeId3      string `json:"user_tipe_id3" form:"user_tipe_id3"`
	UserTipeId4      string `json:"user_tipe_id4" form:"user_tipe_id4"`
	UserId           string
	StatusPembayaran string `json:"status_pembayaran" form:"status_pembayaran"`
}

type PemesananUserRequestUpdateFormat struct {
	PemesananNama    string `json:"pemesanan_nama" form:"pemesanan_nama"`
	PemesananDate    string `json:"pemesanan_date" form:"pemesanan_date"`
	PemesananTime    string `json:"pemesanan_time" form:"pemesanan_time"`
	PemesananEmail   string `json:"pemesanan_email" form:"pemesanan_email"`
	PemesananNoHp    string `json:"pemesanan_no_hp" form:"pemesanan_no_hp"`
	PemesananStatus  string `json:"pemesanan_status" form:"pemesanan_status"`
	PemesananAlamat  string `json:"pemesanan_alamat" form:"pemesanan_alamat"`
	StatusPembayaran string `json:"status_pembayaran" form:"status_pembayaran"`
}

type PemesananUserResponseFormat struct {
	PemesananId      string `json:"pemesanan_id"`
	PemesananNama    string `json:"pemesanan_nama"`
	PemesananDate    string `json:"pemesanan_date"`
	PemesananTime    string `json:"pemesanan_time"`
	PemesananEmail   string `json:"pemesanan_email"`
	PemesananNoHp    string `json:"pemesanan_no_hp"`
	PemesananStatus  string `json:"pemesanan_status"`
	PemesananAlamat  string `json:"pemesanan_alamat"`
	CreatedAt        string `json:"created_at"`
	FirstPlayer      string `json:"first_player" form:"first_player"`
	SecondPlayer     string `json:"second_player" form:"second_player"`
	ThirdPlayer      string `json:"third_player" form:"third_player"`
	FourthPlayer     string `json:"fourth_player" form:"fourth_player"`
	UserTipeId1      string `json:"user_tipe_id1" form:"user_tipe_id1"`
	UserTipeId2      string `json:"user_tipe_id2" form:"user_tipe_id2"`
	UserTipeId3      string `json:"user_tipe_id3" form:"user_tipe_id3"`
	UserTipeId4      string `json:"user_tipe_id4" form:"user_tipe_id4"`
	UserId           string `json:"user_id" form:"user_id"`
	StatusPembayaran string `json:"status_pembayaran" form:"status_pembayaran"`
}
type PemesananUserResponseFormatDatatables struct {
	PemesananId      string `json:"pemesanan_id"`
	PemesananNama    string `json:"pemesanan_nama"`
	PemesananDate    string `json:"pemesanan_date"`
	PemesananTime    string `json:"pemesanan_time"`
	PemesananEmail   string `json:"pemesanan_email"`
	PemesananNoHp    string `json:"pemesanan_no_hp"`
	PemesananStatus  string `json:"pemesanan_status"`
	PemesananAlamat  string `json:"pemesanan_alamat"`
	CreatedAt        string `json:"created_at"`
	Action           string `json:"action"`
	FirstPlayer      string `json:"first_player" form:"first_player"`
	SecondPlayer     string `json:"second_player" form:"second_player"`
	ThirdPlayer      string `json:"third_player" form:"third_player"`
	FourthPlayer     string `json:"fourth_player" form:"fourth_player"`
	UserTipeId1      string `json:"user_tipe_id1" form:"user_tipe_id1"`
	UserTipeId2      string `json:"user_tipe_id2" form:"user_tipe_id2"`
	UserTipeId3      string `json:"user_tipe_id3" form:"user_tipe_id3"`
	UserTipeId4      string `json:"user_tipe_id4" form:"user_tipe_id4"`
	UserId           string `json:"user_id" form:"user_id"`
	StatusPembayaran string `json:"status_pembayaran" form:"status_pembayaran"`
}
