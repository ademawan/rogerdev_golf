package entities

import "mime/multipart"

type ProfilPerusahaan struct {
	ProfilPerusahaanId        string
	ProfilPerusahaanNama      string
	ProfilPerusahaanEmail     string
	ProfilPerusahaanAlamat    string
	ProfilPerusahaanNoHp      string
	ProfilPerusahaanDeskripsi string
	ProfilPerusahaanImage     string
}
type ProfilPerusahaanRequestCreateFormat struct {
	ProfilPerusahaanNama      string                `json:"profil_perusahaan_nama" form:"profil_perusahaan_nama"`
	ProfilPerusahaanEmail     string                `json:"profil_perusahaan_email" form:"profil_perusahaan_email"`
	ProfilPerusahaanAlamat    string                `json:"profil_perusahaan_alamat" form:"profil_perusahaan_alamat"`
	ProfilPerusahaanNoHp      string                `json:"profil_perusahaan_no_hp" form:"profil_perusahaan_no_hp"`
	ProfilPerusahaanDeskripsi string                `json:"profil_perusahaan_deskripsi" form:"profil_perusahaan_deskripsi"`
	ProfilPerusahaanImage     *multipart.FileHeader `json:"profil_perusahaan_image" form:"profil_perusahaan_image"`
}

type ProfilPerusahaanRequestUpdateFormat struct {
	ProfilPerusahaanNama      string `json:"profil_perusahaan_nama" form:"profil_perusahaan_nama"`
	ProfilPerusahaanEmail     string `json:"profil_perusahaan_email" form:"profil_perusahaan_email"`
	ProfilPerusahaanAlamat    string `json:"profil_perusahaan_alamat" form:"profil_perusahaan_alamat"`
	ProfilPerusahaanNoHp      string `json:"profil_perusahaan_no_hp" form:"profil_perusahaan_no_hp"`
	ProfilPerusahaanDeskripsi string `json:"profil_perusahaan_deskripsi" form:"profil_perusahaan_deskripsi"`
	ProfilPerusahaanImage     string `json:"profil_perusahaan_image" form:"profil_perusahaan_image"`
}

type ProfilPerusahaanResponseFormat struct {
	ProfilPerusahaanId        string `json:"profil_perusahaan_id"`
	ProfilPerusahaanNama      string `json:"profil_perusahaan_nama"`
	ProfilPerusahaanEmail     string `json:"profil_perusahaan_email"`
	ProfilPerusahaanAlamat    string `json:"profil_perusahaan_alamat"`
	ProfilPerusahaanNoHp      string `json:"profil_perusahaan_no_hp"`
	ProfilPerusahaanDeskripsi string `json:"profil_perusahaan_deskripsi"`
	ProfilPerusahaanImage     string `json:"profil_perusahaan_image"`
}

type ProfilPerusahaanResponseFormatDatatables struct {
	ProfilPerusahaanId        string `json:"profil_perusahaan_id"`
	ProfilPerusahaanNama      string `json:"profil_perusahaan_nama"`
	ProfilPerusahaanEmail     string `json:"profil_perusahaan_email"`
	ProfilPerusahaanAlamat    string `json:"profil_perusahaan_alamat"`
	ProfilPerusahaanNoHp      string `json:"profil_perusahaan_no_hp"`
	ProfilPerusahaanDeskripsi string `json:"profil_perusahaan_deskripsi"`
	ProfilPerusahaanImage     string `json:"profil_perusahaan_image"`
	Action                    string `json:"action"`
}
