package main

import (
	"fmt"
	"io"
	"rogerdev_golf/configs"
	ac "rogerdev_golf/delivery/controllers/auth"
	uc "rogerdev_golf/delivery/controllers/user"
	"rogerdev_golf/delivery/routes"
	authRepo "rogerdev_golf/repository/auth"
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

	authRepo := authRepo.New(db)
	userRepo := userRepo.New(db)

	authController := ac.New(authRepo)
	userController := uc.New(userRepo)

	e := echo.New()
	e.Static("/", "assets")
	e.Renderer = NewRenderer("./views/*.html", true)
	e.Validator = &CustomValidator{validator: validator.New()}

	routes.RegisterPath(e, authController, userController)

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
