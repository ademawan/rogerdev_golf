package entities

type User struct {
	UserId       string
	UserNama     string
	UserEmail    string
	UserPassword string
	UserAlamat   string
	UserNoHp     string
	UserTipeId   string
	IsUser       string
	CreatedAt    string
}

type UserRequestCreateFormat struct {
	UserNama     string `json:"user_nama" form:"user_nama"`
	UserEmail    string `json:"user_email" form:"user_email"`
	UserPassword string `json:"user_password" form:"user_password"`
	UserAlamat   string `json:"user_alamat" form:"user_alamat"`
	UserNoHp     string `json:"user_no_hp" form:"user_no_hp"`
	UserTipeId   string `json:"user_tipe_id" form:"user_tipe_id"`
}

type UserRequestUpdateFormat struct {
	UserNama   string `json:"user_nama" form:"user_nama"`
	UserEmail  string `json:"user_email" form:"user_email"`
	UserAlamat string `json:"user_alamat" form:"user_alamat"`
	UserNoHp   string `json:"user_no_hp" form:"user_no_hp"`
	UserTipeId string `json:"user_tipe_id" form:"user_tipe_id"`
}

type UserResponseFormat struct {
	UserId       string `json:"user_id"`
	UserNama     string `json:"user_nama"`
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"-"`
	UserAlamat   string `json:"user_alamat"`
	UserNoHp     string `json:"user_no_hp"`
	UserTipeId   string `json:"user_tipe_id"`
	IsUser       string `json:"is_user"`
	CreatedAt    string `json:"created_at"`
}

type UserResponseFormatDatatables struct {
	UserId       string `json:"user_id"`
	UserNama     string `json:"user_nama"`
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"-"`
	UserAlamat   string `json:"user_alamat"`
	UserNoHp     string `json:"user_no_hp"`
	UserTipeId   string `json:"user_tipe_id"`
	UserTipeNama string `json:"user_tipe_nama"`
	IsUser       string `json:"is_user"`
	CreatedAt    string `json:"created_at"`
	Action       string `json:"action"`
}
