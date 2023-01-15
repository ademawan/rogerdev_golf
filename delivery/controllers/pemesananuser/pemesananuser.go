package pemesananuser

import (
	"net/http"
	"os"
	"rogerdev_golf/delivery/controllers/common"
	"rogerdev_golf/entities"
	"rogerdev_golf/middlewares"
	"rogerdev_golf/repository/pemesananuser"
	"strconv"
	"time"

	// utils "todo-list-app/utils/aws_S3"

	// "github.com/aws/aws-sdk-go/aws/session"
	"github.com/labstack/echo/v4"
)

type PemesananUserController struct {
	repo pemesananuser.PemesananUser
	// conn *session.Session
}

func New(repository pemesananuser.PemesananUser /*, S3 *session.Session*/) *PemesananUserController {
	return &PemesananUserController{
		repo: repository,
		// conn: S3,
	}
}
func (uc *PemesananUserController) UI() echo.HandlerFunc {
	return func(c echo.Context) error {
		path := os.Getenv("BASE_URL")
		var dataMap = make(map[string]interface{})
		dataMap["path"] = path

		return c.Render(http.StatusOK, "pemesananuser.html", dataMap)

	}
}
func (uc *PemesananUserController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {

		req := entities.PemesananUserRequestCreateFormat{}

		c.Bind(&req)
		err := c.Validate(&req)

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ResponseUser(http.StatusBadRequest, "There is some problem from input", nil))
		}

		// file, errO := c.FormFile("image")
		// if errO != nil {
		// 	log.Info(errO)
		// }

		// if file != nil {
		// 	src, _ := file.Open()
		// 	link, errU := utils.Upload(uc.conn, src, *file)
		// 	if errU != nil {
		// 		return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "Upload Failed", nil))
		// 	}
		// 	user.Image = link
		// } else if file == nil {
		// 	user.Image = ""
		// }
		// timeLocation, _ := time.LoadLocation("Asia/Jakarta")
		// timeNow := time.Now().In(timeLocation).Unix()

		pemesananId := "PID00" + strconv.Itoa(int(time.Now().Unix()))
		err_repo := uc.repo.Create(&entities.PemesananUser{
			PemesananId:     pemesananId,
			PemesananNama:   req.PemesananNama,
			PemesananDate:   req.PemesananDate,
			PemesananTime:   req.PemesananTime,
			PemesananEmail:  req.PemesananEmail,
			PemesananNoHp:   req.PemesananNoHp,
			PemesananStatus: req.PemesananStatus,
			PemesananAlamat: req.PemesananAlamat,
		})
		if err_repo != nil {
			return c.JSON(http.StatusConflict, common.ResponseUser(http.StatusConflict, err_repo.Error(), nil))
		}

		return c.JSON(http.StatusCreated, common.ResponseUser(http.StatusCreated, "Success create pemesanan", nil))

	}
}

func (uc *PemesananUserController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		pemesanaId := c.Param("pemesananid")

		res, err := uc.repo.Get(pemesanaId)

		if err != nil {
			statusCode := http.StatusInternalServerError
			errorMessage := "There is some problem from the server"
			if err.Error() == "record not found" {
				statusCode = http.StatusNotFound
				errorMessage = err.Error()
			}
			return c.JSON(statusCode, common.ResponseUser(http.StatusNotFound, errorMessage, nil))
		}

		// res.CreatedAt = uc.TimeToUser(res.CreatedAt.(int64))
		// res.UpdatedAt = uc.TimeToUser(res.UpdatedAt.(int64))
		// res.DeletedAt = uc.TimeToUser(res.DeletedAt.(int64))

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success get pemesanan", res))
	}
}

func (uc *PemesananUserController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {

		res, err := uc.repo.GetAll()

		if err != nil {
			statusCode := http.StatusInternalServerError
			errorMessage := "There is some problem from the server"
			if err.Error() == "record not found" {
				statusCode = http.StatusNotFound
				errorMessage = err.Error()
			}
			return c.JSON(statusCode, common.ResponseUser(http.StatusNotFound, errorMessage, nil))
		}

		// res.CreatedAt = uc.TimeToUser(res.CreatedAt.(int64))
		// res.UpdatedAt = uc.TimeToUser(res.UpdatedAt.(int64))
		// res.DeletedAt = uc.TimeToUser(res.DeletedAt.(int64))

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success get pemesanan", res))
	}
}

func (uc *PemesananUserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		pemesanaId := c.Param("pemesananid")
		var req = entities.PemesananUserRequestCreateFormat{}

		err_repo := uc.repo.Update(&entities.PemesananUser{
			PemesananId:     pemesanaId,
			PemesananNama:   req.PemesananNama,
			PemesananDate:   req.PemesananDate,
			PemesananTime:   req.PemesananTime,
			PemesananEmail:  req.PemesananEmail,
			PemesananNoHp:   req.PemesananNoHp,
			PemesananStatus: req.PemesananStatus,
			PemesananAlamat: req.PemesananAlamat,
		})

		if err_repo != nil {
			return c.JSON(http.StatusInternalServerError, common.ResponseUser(http.StatusInternalServerError, "There is some error on server", nil))
		}

		// response.Roles = res.Roles
		// response.Image = res.Image

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success update pemesanan", nil))
	}
}

func (uc *PemesananUserController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		pemesanaId := c.Param("pemesananid")
		_, err := uc.repo.Delete(pemesanaId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.ResponseUser(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success delete fasilitas", nil))
	}
}

func (uc *PemesananUserController) GetAllDatatables() echo.HandlerFunc {
	return func(c echo.Context) error {

		role := middlewares.ExtractRoles(c)
		if role == "1" {
			mapping := make(map[string]interface{})
			mapping["message"] = "unauthorize"
			mapping["login"] = "1"

			return c.JSON(http.StatusUnauthorized, mapping)
		}

		output := make(map[string]interface{})
		output["draw"] = 1
		output["recordsTotal"] = 0
		output["recordsFiltered"] = 0
		data := make([]interface{}, 0)
		output["data"] = data
		output["error"] = ""
		res, count, err := uc.repo.GetAllDatatables()

		if err != nil {
			statusCode := http.StatusInternalServerError
			errorMessage := "There is some problem from the server"
			if err.Error() == "record not found" {
				statusCode = http.StatusNotFound
				errorMessage = err.Error()
			}
			return c.JSON(statusCode, common.ResponseUser(http.StatusNotFound, errorMessage, nil))
		}

		// res.CreatedAt = uc.TimeToUser(res.CreatedAt.(int64))
		// res.UpdatedAt = uc.TimeToUser(res.UpdatedAt.(int64))
		// res.DeletedAt = uc.TimeToUser(res.DeletedAt.(int64))
		if len(res) != 0 {
			output["start"] = 1
			output["draw"] = 20
			output["data"] = res
			output["recordsTotal"] = count
			output["recordsFiltered"] = 20
		}

		return c.JSON(http.StatusOK, output)
	}
}

// func (uc *PemesananUserController) Dummy() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		q, _ := strconv.Atoi(c.QueryParam("length"))

// 		result := uc.repo.Dummy(q)
// 		if !result {
// 			return c.JSON(http.StatusInternalServerError, common.ResponseUser(http.StatusInternalServerError, "There is some error on server", nil))
// 		}
// 		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success create user dummy", nil))

//		}
//	}
func (c *PemesananUserController) TimeToUser(timeInt int64) string {
	if timeInt <= 0 {
		return ""
	}
	i, err := strconv.ParseInt(strconv.Itoa(int(timeInt)), 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	year := strconv.Itoa(tm.Year())
	month := strconv.Itoa(int(tm.Month()))
	h, m, s := tm.Clock()
	hour := strconv.Itoa(h)
	if len(hour) == 1 {
		hour = "0" + hour
	}
	minute := strconv.Itoa(m)
	if len(minute) == 1 {
		minute = "0" + minute
	}
	second := strconv.Itoa(s)

	switch month {
	case "1":
		month = "Januari"
	case "2":
		month = "Februari"
	case "3":
		month = "Maret"
	case "4":
		month = "April"
	case "5":
		month = "Mei"
	case "6":
		month = "Juni"
	case "7":
		month = "Juli"
	case "8":
		month = "Agustus"
	case "9":
		month = "September"
	case "10":
		month = "Oktober"
	case "11":
		month = "November"
	case "12":
		month = "Desember"
	}
	day := strconv.Itoa(tm.Day())
	createdOnString := day + " " + month + " " + year + " Pukul " + hour + ":" + minute + ":" + second + " WIB"
	return createdOnString
}
