package galeri

import (
	"net/http"
	"os"
	"rogerdev_golf/delivery/controllers/common"
	"rogerdev_golf/entities"
	"rogerdev_golf/repository/galeri"
	"strconv"
	"time"

	// utils "todo-list-app/utils/aws_S3"

	// "github.com/aws/aws-sdk-go/aws/session"
	"github.com/labstack/echo/v4"
)

type GaleriController struct {
	repo galeri.Galeri
	// conn *session.Session
}

func New(repository galeri.Galeri /*, S3 *session.Session*/) *GaleriController {
	return &GaleriController{
		repo: repository,
		// conn: S3,
	}
}
func (uc *GaleriController) UI() echo.HandlerFunc {
	return func(c echo.Context) error {
		path := os.Getenv("BASE_URL")
		var dataMap = make(map[string]interface{})
		dataMap["path"] = path

		return c.Render(http.StatusOK, "adminfasilitas.html", dataMap)

	}
}
func (uc *GaleriController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {

		req := entities.GaleriRequestCreateFormat{}

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

		galeriId := "FID00" + strconv.Itoa(int(time.Now().Unix()))
		err_repo := uc.repo.Create(&entities.Galeri{
			GaleriId:    galeriId,
			GaleriNama:  req.GaleriNama,
			GaleriImage: req.GaleriImage,
		})
		if err_repo != nil {
			return c.JSON(http.StatusConflict, common.ResponseUser(http.StatusConflict, err_repo.Error(), nil))
		}

		return c.JSON(http.StatusCreated, common.ResponseUser(http.StatusCreated, "Success create galeri", nil))

	}
}

func (uc *GaleriController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Request().Header.Get("user_id")

		res, err := uc.repo.Get(userId)

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

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success get galeri", res))
	}
}

func (uc *GaleriController) GetAll() echo.HandlerFunc {
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

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success get galeri", res))
	}
}

func (uc *GaleriController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		galeriId := c.Request().Header.Get("user_id")
		var req = entities.GaleriRequestUpdateFormat{}

		err_repo := uc.repo.Update(&entities.Galeri{
			GaleriId:    galeriId,
			GaleriNama:  req.GaleriNama,
			GaleriImage: req.GaleriImage,
		})

		if err_repo != nil {
			return c.JSON(http.StatusInternalServerError, common.ResponseUser(http.StatusInternalServerError, "There is some error on server", nil))
		}

		// response.Roles = res.Roles
		// response.Image = res.Image

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success update fasilitas", nil))
	}
}

func (uc *GaleriController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Request().Header.Get("user_id")
		_, err := uc.repo.Delete(userId)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.ResponseUser(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success delete fasilitas", nil))
	}
}

func (uc *GaleriController) GetAllDatatables() echo.HandlerFunc {
	return func(c echo.Context) error {
		// userUid := middlewares.ExtractTokenUserUid(c)

		// res, err := uc.repo.GetByUid(userUid)

		// if err != nil {
		// 	statusCode := http.StatusInternalServerError
		// 	errorMessage := "There is some problem from the server"
		// 	if err.Error() == "record not found" {
		// 		statusCode = http.StatusNotFound
		// 		errorMessage = err.Error()
		// 	}
		// 	return c.JSON(statusCode, common.ResponseUser(http.StatusNotFound, errorMessage, nil))
		// }

		// res.CreatedAt = uc.TimeToUser(res.CreatedAt.(int64))
		// res.UpdatedAt = uc.TimeToUser(res.UpdatedAt.(int64))
		// res.DeletedAt = uc.TimeToUser(res.DeletedAt.(int64))

		output := make(map[string]interface{})
		output["draw"] = 1
		output["recordsTotal"] = 0
		output["recordsFiltered"] = 0
		data := make([]interface{}, 0)
		output["data"] = data
		output["error"] = ""

		res := make(map[string]interface{})
		res["user_nama"] = "test"
		res["user_email"] = 50
		res["user_alamat"] = "honda"
		res["user_no_hp"] = "081221997456"
		res["user_tipe"] = "Member"
		res["action"] = `<button type="submit"></button>`
		data = append(data, res)
		output["start"] = 1
		output["draw"] = 1
		output["recordsTotal"] = 2
		output["recordsFiltered"] = 2

		output["data"] = data

		return c.JSON(http.StatusOK, output)
	}
}

// func (uc *GaleriController) Dummy() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		q, _ := strconv.Atoi(c.QueryParam("length"))

// 		result := uc.repo.Dummy(q)
// 		if !result {
// 			return c.JSON(http.StatusInternalServerError, common.ResponseUser(http.StatusInternalServerError, "There is some error on server", nil))
// 		}
// 		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success create user dummy", nil))

//		}
//	}
func (c *GaleriController) TimeToUser(timeInt int64) string {
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
