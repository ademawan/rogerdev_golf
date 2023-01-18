package routes

import (
	"rogerdev_golf/delivery/controllers/auth"

	"rogerdev_golf/delivery/controllers/fasilitas"
	"rogerdev_golf/delivery/controllers/galeri"
	"rogerdev_golf/delivery/controllers/pemesanan"
	"rogerdev_golf/delivery/controllers/pemesananuser"
	"rogerdev_golf/delivery/controllers/profilperusahaan"

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
	puc *pemesananuser.PemesananUserController,
	pruc *profilperusahaan.ProfilPerusahaanController,

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

	e.GET("/admin/login", aa.AdminLoginPage())
	e.GET("/admin/register", aa.AdminRegisterPage())
	e.GET("/user/login", aa.LoginPage())
	e.GET("/user/register", aa.RegisterPage())

	// e.GET("/admin/dashboard", aa.Dashboard())

	//ROUTE REGISTER - LOGIN USERS
	e.POST("/admin/register", uc.Register())
	e.POST("/admin/login", aa.Login())
	e.POST("/user/register", uc.Register())
	e.POST("/user/login", aa.Login())
	// e.GET("google/login", aa.LoginGoogle())
	// e.GET("google/callback", aa.LoginGoogleCallback())
	e.POST("users/logout", aa.Logout(), middlewares.JwtMiddleware())

	//ROUTE USERS
	e.GET("/users/me", uc.GetByUid(), middlewares.JwtMiddleware())
	e.PUT("/users/me", uc.Update(), middlewares.JwtMiddleware())
	e.DELETE("/users/me", uc.Delete(), middlewares.JwtMiddleware())

	e.GET("/admin/user", uc.UI())
	e.GET("/admin/user/datatables", uc.GetAllDatatables(), middlewares.JwtMiddleware())
	e.POST("/admin/user/ajax", uc.Register())
	e.GET("/admin/user/ajax", uc.GetAll())
	e.GET("/admin/user/ajax/:userid/edit", uc.GetByUid())
	e.PUT("/admin/user/ajax/:userid", uc.Update())
	e.DELETE("/admin/user/ajax/:userid", uc.Delete())

	//ROUTE PEMESANAN
	e.GET("/admin/pemesanan", pc.UI())
	e.GET("/admin/pemesanan/datatables", pc.GetAllDatatables(), middlewares.JwtMiddleware())
	e.POST("/admin/pemesanan/ajax", pc.Create(), middlewares.JwtMiddleware())
	e.GET("/admin/pemesanan/ajax", pc.GetAll())
	e.GET("/admin/pemesanan/ajax/:pemesananid/edit", pc.Get())
	e.PUT("/admin/pemesanan/ajax/:pemesananid", pc.Update())
	e.DELETE("/admin/pemesanan/ajax/:pemesananid", pc.Delete())

	//ROUTE PEMESANAN USER
	e.GET("/user/pemesanan", puc.UI())
	e.GET("/user/pemesanan/datatables", puc.GetAllDatatables(), middlewares.JwtMiddleware())
	e.POST("/user/pemesanan/ajax", puc.Create(), middlewares.JwtMiddleware())
	// e.GET("/user/pemesanan/ajax", puc.GetAll())
	e.GET("/user/pemesanan/ajax/:pemesananid/edit", puc.Get())
	// e.PUT("/user/pemesanan/ajax/:pemesananid", puc.Update())
	e.GET("/user/pemesanan/ajax/:pemesananid/update", puc.UpdatePembayaran())
	e.DELETE("/user/pemesanan/ajax/:pemesananid", puc.Delete())

	//ROUTE FASILITAS
	e.GET("/admin/fasilitas", fc.UI())
	e.GET("/admin/fasilitas/datatables", fc.GetAllDatatables(), middlewares.JwtMiddleware())
	e.POST("/admin/fasilitas/ajax", fc.Create())
	e.GET("/admin/fasilitas/ajax", fc.GetAll())
	e.GET("/admin/fasilitas/ajax/:fasilitasid/edit", pc.Get())
	e.PUT("/admin/fasilitas/ajax/:fasilitasid", fc.Update())
	e.DELETE("/admin/fasilitas/ajax/:fasilitasid", fc.Delete())

	//ROUTE GALERI
	e.GET("/admin/galeri", gc.UI())
	e.GET("/admin/galeri/datatables", gc.GetAllDatatables(), middlewares.JwtMiddleware())
	e.POST("/admin/galeri/ajax", gc.Create())
	e.GET("/admin/galeri/ajax", gc.GetAll())
	e.GET("/admin/galeri/ajax/:galeriid/edit", pc.Get())
	e.PUT("/admin/galeri/ajax/:galeriid", gc.Update())
	e.DELETE("/admin/galeri/ajax/:galeriid", gc.Delete())

	//ROUTE PROFIL PERUSAHAAN
	e.GET("/admin/profilperusahaan", pruc.UI())
	e.GET("/admin/profilperusahaan/datatables", pruc.GetAllDatatables(), middlewares.JwtMiddleware())
	e.POST("/admin/profilperusahaan/ajax", pruc.Create())
	e.GET("/admin/profilperusahaan/ajax", pruc.GetAll())
	e.GET("/admin/profilperusahaan/ajax/get", pruc.GetProfilPerusahaan())
	e.GET("/admin/profilperusahaan/ajax/:profilperusahaanid/edit", pruc.Get())
	e.PUT("/admin/profilperusahaan/ajax/:profilperusahaanid", pruc.GetAllDatatables())
	e.DELETE("/admin/profilperusahaan/ajax/:profilperusahaanid", pruc.GetAllDatatables())

	e.POST("/admin/profilperusahaan", pruc.CreateImage())

	// e.GET("/users/me/search", pc.Search())
	// e.GET("/users/me/dummy", pc.Dummy())

}
