package response

type Success struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
type SuccessGetAll struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type SuccessGetAllWithPaginator struct {
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Paginator interface{} `json:"paginator"`
}

type RerponseFailedFormat struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
