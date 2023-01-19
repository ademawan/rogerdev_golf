package main

import (
	"fmt"
	"io"
	"rogerdev_golf/configs"
	ac "rogerdev_golf/delivery/controllers/auth"
	fc "rogerdev_golf/delivery/controllers/fasilitas"
	gc "rogerdev_golf/delivery/controllers/galeri"
	pc "rogerdev_golf/delivery/controllers/pemesanan"
	puc "rogerdev_golf/delivery/controllers/pemesananuser"
	pruc "rogerdev_golf/delivery/controllers/profilperusahaan"

	uc "rogerdev_golf/delivery/controllers/user"
	"rogerdev_golf/delivery/routes"
	authRepo "rogerdev_golf/repository/auth"
	fasilitasRepo "rogerdev_golf/repository/fasilitas"
	galeriRepo "rogerdev_golf/repository/galeri"
	pemesananRepo "rogerdev_golf/repository/pemesanan"
	pemesananuserRepo "rogerdev_golf/repository/pemesananuser"
	profilperusahaanRepo "rogerdev_golf/repository/profilperusahaan"

	userRepo "rogerdev_golf/repository/user"
	"rogerdev_golf/utils"
	"text/template"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"

	"github.com/labstack/gommon/log"
)

type M map[string]interface{}

type Renderer struct {
	template *template.Template
	debug    bool
	location string
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	config := configs.GetConfig()

	db := utils.InitDB(config)
	defer db.Close()

	authRepo := authRepo.New(db)
	userRepo := userRepo.NewUserRepository(db)
	pemesananRepo := pemesananRepo.NewPemesananRepository(db)
	pemesananuserRepo := pemesananuserRepo.NewPemesananUserRepository(db)
	profilperusahaanRepo := profilperusahaanRepo.NewProfilPerusahaanRepository(db)

	galeriRepo := galeriRepo.NewGaleriRepository(db)
	fasilitasRepo := fasilitasRepo.NewFasilitasRepository(db)

	authController := ac.New(authRepo, profilperusahaanRepo)
	userController := uc.New(userRepo)
	pemesananController := pc.New(pemesananRepo)
	pemesananuserController := puc.New(pemesananuserRepo)
	profilperusahaanController := pruc.New(profilperusahaanRepo)

	galeriController := gc.New(galeriRepo)
	fasilitasController := fc.New(fasilitasRepo)

	e := echo.New()
	e.Static("/", "assets")
	e.Renderer = NewRenderer("./views/*.html", true)
	e.Validator = &CustomValidator{validator: validator.New()}

	routes.RegisterPath(e, authController, userController, pemesananController, galeriController, fasilitasController, pemesananuserController, profilperusahaanController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}

func NewRenderer(location string, debug bool) *Renderer {
	tpl := new(Renderer)
	tpl.location = location
	tpl.debug = debug

	tpl.ReloadTemplates()

	return tpl
}

func (t *Renderer) ReloadTemplates() {
	t.template = template.Must(template.ParseGlob(t.location))
}
func (t *Renderer) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context,
) error {
	if t.debug {
		t.ReloadTemplates()
	}

	return t.template.ExecuteTemplate(w, name, data)
}
