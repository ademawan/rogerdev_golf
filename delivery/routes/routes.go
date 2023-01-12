package routes

import (
	"rogerdev_golf/delivery/controllers/auth"

	"rogerdev_golf/delivery/controllers/user"
	"rogerdev_golf/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo,
	aa *auth.AuthController,
	uc *user.UserController,

) {

	//CORS
	e.Use(middleware.CORS())

	//LOGGER
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	//ROUTE index
	e.GET("/bogor-raya-club", aa.Index())
	e.GET("/bogor-raya-club/about", aa.About())
	e.GET("/bogor-raya-club/services", aa.Service())
	e.GET("/bogor-raya-club/projects", aa.Project())
	e.GET("/bogor-raya-club/blog", aa.Blog())
	e.GET("/bogor-raya-club/contact", aa.Contact())

	e.GET("/login", aa.LoginPage())
	e.GET("/register", aa.RegisterPage())

	e.GET("/undangan/khitanan", aa.Post())
	e.GET("/undangan/cover", aa.CoverPage())

	e.GET("/dashboard", aa.Dashboard())

	//ROUTE REGISTER - LOGIN USERS
	e.POST("users/register", uc.Register())
	e.POST("users/login", aa.Login())
	// e.GET("google/login", aa.LoginGoogle())
	// e.GET("google/callback", aa.LoginGoogleCallback())
	e.POST("users/logout", aa.Logout(), middlewares.JwtMiddleware())

	//ROUTE USERS
	e.GET("/users/me", uc.GetByUid(), middlewares.JwtMiddleware())
	e.PUT("/users/me", uc.Update(), middlewares.JwtMiddleware())
	e.DELETE("/users/me", uc.Delete(), middlewares.JwtMiddleware())
	// e.GET("/users/me/search", uc.Search())
	// e.GET("/users/me/dummy", uc.Dummy())

}
