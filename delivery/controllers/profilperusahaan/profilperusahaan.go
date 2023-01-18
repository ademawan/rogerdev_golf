package profilperusahaan

import (
	"fmt"
	"net/http"
	"os"
	"rogerdev_golf/delivery/controllers/common"
	"rogerdev_golf/entities"
	"rogerdev_golf/middlewares"
	"rogerdev_golf/repository/profilperusahaan"
	"rogerdev_golf/utils"
	utilsimage "rogerdev_golf/utils"
	"strconv"
	"time"

	// utils "todo-list-app/utils/aws_S3"

	// "github.com/aws/aws-sdk-go/aws/session"

	"github.com/labstack/echo/v4"
)

type ProfilPerusahaanController struct {
	repo profilperusahaan.ProfilPerusahaan
	// conn *session.Session
}

func New(repository profilperusahaan.ProfilPerusahaan /*, S3 *session.Session*/) *ProfilPerusahaanController {
	return &ProfilPerusahaanController{
		repo: repository,
		// conn: S3,
	}
}
func (uc *ProfilPerusahaanController) UI() echo.HandlerFunc {
	return func(c echo.Context) error {
		path := os.Getenv("BASE_URL")
		var dataMap = make(map[string]interface{})
		dataMap["path"] = path

		return c.Render(http.StatusOK, "adminprofilperusahaan.html", dataMap)

	}
}
func (uc *ProfilPerusahaanController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {

		req := entities.ProfilPerusahaanRequestCreateFormat{}

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
		image, err := utils.UploadImage(req.ProfilPerusahaanImage)
		if err != nil {
			return c.JSON(http.StatusConflict, common.ResponseUser(http.StatusConflict, err.Error(), nil))

		}

		profilPerusahaan := "PRID00" + strconv.Itoa(int(time.Now().Unix()))
		err_repo := uc.repo.Create(entities.ProfilPerusahaan{
			ProfilPerusahaanId:        profilPerusahaan,
			ProfilPerusahaanNama:      req.ProfilPerusahaanNama,
			ProfilPerusahaanEmail:     req.ProfilPerusahaanEmail,
			ProfilPerusahaanAlamat:    req.ProfilPerusahaanAlamat,
			ProfilPerusahaanNoHp:      req.ProfilPerusahaanNoHp,
			ProfilPerusahaanDeskripsi: req.ProfilPerusahaanDeskripsi,
			ProfilPerusahaanImage:     image.Path,
		})
		if err_repo != nil {
			return c.JSON(http.StatusConflict, common.ResponseUser(http.StatusConflict, err_repo.Error(), nil))
		}

		return c.JSON(http.StatusCreated, common.ResponseUser(http.StatusCreated, "Success create profil perusahaan", nil))

	}
}

func (uc *ProfilPerusahaanController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		profilPerusahaanId := c.Param("profilperusahaanid")
		fmt.Println(profilPerusahaanId, "HAAAALLLOOOOO")
		res, err := uc.repo.Get(profilPerusahaanId)
		if err != nil {
			statusCode := http.StatusInternalServerError
			errorMessage := "There is some problem from the server"
			if err.Error() == "profil perusahaan not found" {
				statusCode = http.StatusNotFound
				errorMessage = err.Error()
			}
			return c.JSON(statusCode, common.ResponseUser(http.StatusNotFound, errorMessage, nil))
		}

		// res.CreatedAt = uc.TimeToUser(res.CreatedAt.(int64))
		// res.UpdatedAt = uc.TimeToUser(res.UpdatedAt.(int64))
		// res.DeletedAt = uc.TimeToUser(res.DeletedAt.(int64))

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success get user", res))
	}
}

func (uc *ProfilPerusahaanController) GetProfilPerusahaan() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := uc.repo.GetProfilPerusahaan()
		if err != nil {
			statusCode := http.StatusInternalServerError
			errorMessage := "There is some problem from the server"
			if err.Error() == "profil perusahaan not found" {
				statusCode = http.StatusNotFound
				errorMessage = err.Error()
			}
			return c.JSON(statusCode, common.ResponseUser(http.StatusNotFound, errorMessage, nil))
		}

		// res.CreatedAt = uc.TimeToUser(res.CreatedAt.(int64))
		// res.UpdatedAt = uc.TimeToUser(res.UpdatedAt.(int64))
		// res.DeletedAt = uc.TimeToUser(res.DeletedAt.(int64))

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success get user", res))
	}
}

func (uc *ProfilPerusahaanController) GetAll() echo.HandlerFunc {
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

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success get profil perusahaan", res))
	}
}

func (uc *ProfilPerusahaanController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		profilPerusahaanId := c.Param("profilperusahaanid")

		var newProfilPrusahaan = entities.ProfilPerusahaanRequestUpdateFormat{}
		c.Bind(&newProfilPrusahaan)

		err := c.Validate(&newProfilPrusahaan)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ResponseUser(http.StatusBadRequest, "There is some problem from input", nil))
		}

		// resGet, errGet := uc.repo.GetById(user_uid)
		// if errGet != nil {
		// 	log.Info(resGet)
		// }

		// file, errO := c.FormFile("image")
		// if errO != nil {
		// 	log.Info(errO)
		// } else if errO == nil {
		// 	src, _ := file.Open()
		// 	if resGet.Image != "" {
		// 		var updateImage = resGet.Image
		// 		updateImage = strings.Replace(updateImage, "https://airbnb-app.s3.ap-southeast-1.amazonaws.com/", "", -1)

		// 		var resUp = utils.UpdateUpload(uc.conn, updateImage, src, *file)
		// 		if resUp != "success to update image" {
		// 			return c.JSON(http.StatusInternalServerError, common.InternalServerError(http.StatusInternalServerError, "There is some error on server"+resUp, nil))
		// 		}
		// 	} else if resGet.Image == "" {
		// 		var image, errUp = utils.Upload(uc.conn, src, *file)
		// 		if errUp != nil {
		// 			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "Upload Failed", nil))
		// 		}
		// 		newUser.Image = image
		// 	}
		// }

		err_repo := uc.repo.Update(&entities.ProfilPerusahaan{
			ProfilPerusahaanId:        profilPerusahaanId,
			ProfilPerusahaanNama:      newProfilPrusahaan.ProfilPerusahaanNama,
			ProfilPerusahaanEmail:     newProfilPrusahaan.ProfilPerusahaanEmail,
			ProfilPerusahaanAlamat:    newProfilPrusahaan.ProfilPerusahaanAlamat,
			ProfilPerusahaanNoHp:      newProfilPrusahaan.ProfilPerusahaanNoHp,
			ProfilPerusahaanDeskripsi: newProfilPrusahaan.ProfilPerusahaanDeskripsi,
			ProfilPerusahaanImage:     newProfilPrusahaan.ProfilPerusahaanImage,
			// Image:    newUser.Image,
		})

		if err_repo != nil {
			return c.JSON(http.StatusInternalServerError, common.ResponseUser(http.StatusInternalServerError, "There is some error on server", nil))
		}

		// response.Roles = res.Roles
		// response.Image = res.Image

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success update profil perusahaan", nil))
	}
}

func (uc *ProfilPerusahaanController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		profilPerusahaan := c.Param("profilperusahaanid")

		_, err := uc.repo.Delete(profilPerusahaan)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.ResponseUser(http.StatusInternalServerError, "There is some error on server", nil))
		}

		return c.JSON(http.StatusOK, common.ResponseUser(http.StatusOK, "Success delete profil perusahaan", nil))
	}
}

func (uc *ProfilPerusahaanController) GetAllDatatables() echo.HandlerFunc {
	return func(c echo.Context) error {
		// userUid := middlewares.ExtractTokenUserUid(c)
		// var dataas=localStorage.getItem("accessToken")

		role := middlewares.ExtractRoles(c)
		if role != "1" {
			mapping := make(map[string]interface{})
			mapping["message"] = "unauthorize"
			mapping["error"] = "error"
			if role == "2" || role == "3" {
				return c.JSON(http.StatusForbidden, mapping)
			}
			mapping["login"] = "0"
			return c.JSON(http.StatusUnauthorized, mapping)
		}

		output := make(map[string]interface{})
		output["draw"] = 1
		output["recordsTotal"] = 0
		output["recordsFiltered"] = 0
		data := make([]interface{}, 0)
		output["data"] = data
		output[""] = ""

		res, count, err := uc.repo.GetAllDatatables()
		if err != nil {
			statusCode := http.StatusInternalServerError
			errorMessage := "There is some problem from the server"
			if err.Error() == "record not found" {
				statusCode = http.StatusNotFound
				errorMessage = err.Error()
			}
			return c.JSON(statusCode, common.ResponseUser(http.StatusNotFound, errorMessage, output))

		}

		// res.CreatedAt = uc.TimeToUser(res.CreatedAt.(int64))
		// res.UpdatedAt = uc.TimeToUser(res.UpdatedAt.(int64))
		// res.DeletedAt = uc.TimeToUser(res.DeletedAt.(int64))
		if len(res) == 0 {
			output["draw"] = 1
			profilPerusahaann := []entities.ProfilPerusahaanResponseFormatDatatables{}
			profilPerusahaan := entities.ProfilPerusahaanResponseFormatDatatables{}
			profilPerusahaann = append(profilPerusahaann, profilPerusahaan)
			output["data"] = profilPerusahaann
			output["status"] = 200

			return c.JSON(http.StatusOK, output)
		}
		output["start"] = 1
		output["draw"] = 20
		output["data"] = res
		output["recordsTotal"] = count
		output["recordsFiltered"] = count

		return c.JSON(http.StatusOK, output)
	}
}

func (uc *ProfilPerusahaanController) CreateImage() echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		images, _ := utilsimage.UploadImage(file)
		fmt.Println(images)

		return c.JSON(http.StatusOK, "success")
	}
}

func (c *ProfilPerusahaanController) TimeToUser(timeInt int64) string {
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
