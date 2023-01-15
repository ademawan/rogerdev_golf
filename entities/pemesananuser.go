package entities

type PemesananUser struct {
	PemesananId     string
	PemesananNama   string
	PemesananDate   string
	PemesananTime   string
	PemesananEmail  string
	PemesananNoHp   string
	PemesananStatus string
	PemesananAlamat string
	CreatedAt       string
}

type PemesananUserRequestCreateFormat struct {
	PemesananNama   string `json:"pemesanan_nama" form:"pemesanan_nama"`
	PemesananDate   string `json:"pemesanan_date" form:"pemesanan_date"`
	PemesananTime   string `json:"pemesanan_time" form:"pemesanan_time"`
	PemesananEmail  string `json:"pemesanan_email" form:"pemesanan_email"`
	PemesananNoHp   string `json:"pemesanan_no_hp" form:"pemesanan_no_hp"`
	PemesananStatus string `json:"pemesanan_status" form:"pemesanan_status"`
	PemesananAlamat string `json:"pemesanan_alamat" form:"pemesanan_alamat"`
}

type PemesananUserRequestUpdateFormat struct {
	PemesananNama   string `json:"pemesanan_nama" form:"pemesanan_nama"`
	PemesananDate   string `json:"pemesanan_date" form:"pemesanan_date"`
	PemesananTime   string `json:"pemesanan_time" form:"pemesanan_time"`
	PemesananEmail  string `json:"pemesanan_email" form:"pemesanan_email"`
	PemesananNoHp   string `json:"pemesanan_no_hp" form:"pemesanan_no_hp"`
	PemesananStatus string `json:"pemesanan_status" form:"pemesanan_status"`
	PemesananAlamat string `json:"pemesanan_alamat" form:"pemesanan_alamat"`
}

type PemesananUserResponseFormat struct {
	PemesananId     string `json:"pemesanan_id"`
	PemesananNama   string `json:"pemesanan_nama"`
	PemesananDate   string `json:"pemesanan_date"`
	PemesananTime   string `json:"pemesanan_time"`
	PemesananEmail  string `json:"pemesanan_email"`
	PemesananNoHp   string `json:"pemesanan_no_hp"`
	PemesananStatus string `json:"pemesanan_status"`
	PemesananAlamat string `json:"pemesanan_alamat"`
	CreatedAt       string `json:"created_at"`
}
type PemesananUserResponseFormatDatatables struct {
	PemesananId     string `json:"pemesanan_id"`
	PemesananNama   string `json:"pemesanan_nama"`
	PemesananDate   string `json:"pemesanan_date"`
	PemesananTime   string `json:"pemesanan_time"`
	PemesananEmail  string `json:"pemesanan_email"`
	PemesananNoHp   string `json:"pemesanan_no_hp"`
	PemesananStatus string `json:"pemesanan_status"`
	PemesananAlamat string `json:"pemesanan_alamat"`
	CreatedAt       string `json:"created_at"`
	Action          string `json:"action"`
}
