package entities

type Fasilitas struct {
	FasilitasId    string
	FasilitasNama  string
	FasilitasIcon  string
	FasilitasImage string
	FasilitasBody  string
}

type FasilitasRequestCreateFormat struct {
	FasilitasNama  string `json:"fasilitas_nama" form:"fasilitas_nama"`
	FasilitasIcon  string `json:"fasilitas_icon" form:"fasilitas_icon"`
	FasilitasImage string `json:"fasilitas_image" form:"fasilitas_image"`
	FasilitasBody  string `json:"fasilitas_body" form:"fasilitas_body"`
}

type FasilitasRequestUpdateFormat struct {
	FasilitasNama  string `json:"fasilitas_nama" form:"fasilitas_nama"`
	FasilitasIcon  string `json:"fasilitas_icon" form:"fasilitas_icon"`
	FasilitasImage string `json:"fasilitas_image" form:"fasilitas_image"`
	FasilitasBody  string `json:"fasilitas_body" form:"fasilitas_body"`
}

type FasilitasResponseFormat struct {
	FasilitasId    string `json:"fasilitas_id"`
	FasilitasNama  string `json:"fasilitas_nama"`
	FasilitasIcom  string `json:"fasilitas_icon"`
	FasilitasImage string `json:"fasilitas_image"`
	FasilitasBody  string `json:"fasilitas_body"`
}
type FasilitasResponseFormatDatatables struct {
	FasilitasId    string `json:"fasilitas_id"`
	FasilitasNama  string `json:"fasilitas_nama"`
	FasilitasIcon  string `json:"fasilitas_icon"`
	FasilitasImage string `json:"fasilitas_image"`
	FasilitasBody  string `json:"fasilitas_body"`
	Action         string `json:"action"`
}
