package routes

import (
	"rogerdev_golf/delivery/controllers/auth"

	"rogerdev_golf/delivery/controllers/fasilitas"
	"rogerdev_golf/delivery/controllers/galeri"
	"rogerdev_golf/delivery/controllers/pemesanan"
	"rogerdev_golf/delivery/controllers/user"

	"rogerdev_golf/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo,
	aa *auth.AuthController,
	uc *user.UserController,
	pc *pemesanan.PemesananController,
	gc *galeri.GaleriController,
	fc *fasilitas.FasilitasController,

) {

	//CORS
	e.Use(middleware.CORS())

	//LOGGER
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	//ROUTE index
	e.GET("/", aa.Index())
	e.GET("/tentang", aa.About())
	e.GET("/fasilitas", aa.Service())
	e.GET("/galeri", aa.Project())
	e.GET("/blog", aa.Blog())
	e.GET("/kontak", aa.Contact())
	e.GET("/daftar", aa.Daftar())

	e.GET("/login", aa.LoginPage())
	e.GET("/register", aa.RegisterPage())

	e.GET("/admin/dashboard", aa.Dashboard())

	//ROUTE REGISTER - LOGIN USERS
	e.POST("/users/register", uc.Register())
	e.POST("/users/login", aa.Login())
	// e.GET("google/login", aa.LoginGoogle())
	// e.GET("google/callback", aa.LoginGoogleCallback())
	e.POST("users/logout", aa.Logout(), middlewares.JwtMiddleware())

	//ROUTE USERS
	e.GET("/users/me", uc.GetByUid(), middlewares.JwtMiddleware())
	e.PUT("/users/me", uc.Update(), middlewares.JwtMiddleware())
	e.DELETE("/users/me", uc.Delete(), middlewares.JwtMiddleware())

	e.GET("/admin/user", uc.UI())
	e.GET("/admin/user/datatables", uc.GetAllDatatables())
	e.POST("/admin/user/ajax", uc.Register())
	e.GET("/admin/user/ajax", uc.GetAll())
	e.GET("/admin/user/ajax/:userid/edit", uc.GetByUid())
	e.PUT("/admin/user/ajax/:userid", uc.Update())
	e.DELETE("/admin/user/ajax/:userid", uc.Delete())

	//ROUTE PEMESANAN
	e.GET("/admin/pemesanan", pc.UI())
	e.GET("/admin/pemesanan/datatables", pc.GetAllDatatables())
	e.POST("/admin/pemesanan/ajax", pc.Create())
	e.GET("/admin/pemesanan/ajax", pc.GetAll())
	e.GET("/admin/pemesanan/ajax/:pemesananid/edit", pc.Get())
	e.PUT("/admin/pemesanan/ajax/:pemesananid", pc.Update())
	e.DELETE("/admin/pemesanan/ajax/:pemesananid", pc.Delete())

	//ROUTE FASILITAS
	e.GET("/admin/fasilitas", fc.UI())
	e.GET("/admin/fasilitas/datatables", fc.GetAllDatatables())
	e.POST("/admin/fasilitas/ajax", fc.Create())
	e.GET("/admin/fasilitas/ajax", fc.GetAll())
	e.GET("/admin/fasilitas/ajax/:fasilitasid/edit", pc.Get())
	e.PUT("/admin/fasilitas/ajax/:fasilitasid", fc.Update())
	e.DELETE("/admin/fasilitas/ajax/:fasilitasid", fc.Delete())

	//ROUTE GALERI
	e.GET("/admin/galeri", gc.UI())
	e.GET("/admin/galeri/datatables", gc.GetAllDatatables())
	e.POST("/admin/galeri/ajax", gc.Create())
	e.GET("/admin/galeri/ajax", gc.GetAll())
	e.GET("/admin/galeri/ajax/:galeriid/edit", pc.Get())
	e.PUT("/admin/galeri/ajax/:galeriid", gc.Update())
	e.DELETE("/admin/galeri/ajax/:galeriid", gc.Delete())

	//ROUTE PROFIL PERUSAHAAN
	e.GET("/admin/profile-perusahaan", pc.UI())
	e.GET("/admin/profile-perusahaan/datatables", pc.GetAllDatatables())
	e.POST("/admin/profile-perusahaan/ajax", pc.Create())
	e.GET("/admin/profile-perusahaan/ajax", pc.GetAll())
	e.GET("/admin/profile-perusahaan/ajax/:profileperusahaanid/edit", pc.Get())
	e.PUT("/admin/profile-perusahaan/ajax/:profileperusahaanid", pc.GetAllDatatables())
	e.DELETE("/admin/profile-perusahaan/ajax/:profileperusahaanid", pc.GetAllDatatables())

	// e.GET("/users/me/search", pc.Search())
	// e.GET("/users/me/dummy", pc.Dummy())

}
