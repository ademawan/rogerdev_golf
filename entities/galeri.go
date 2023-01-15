package entities

type Galeri struct {
	GaleriId    string
	GaleriNama  string
	GaleriImage string
}

type GaleriRequestCreateFormat struct {
	GaleriNama  string `json:"galeri_nama"`
	GaleriImage string `json:"galeri_image"`
}

type GaleriRequestUpdateFormat struct {
	GaleriNama  string `json:"galeri_nama"`
	GaleriImage string `json:"galeri_image"`
}

type GaleriResponseFormat struct {
	GaleriId    string `json:"galeri_id"`
	GaleriNama  string `json:"galeri_nama"`
	GaleriImage string `json:"galeri_image"`
}
