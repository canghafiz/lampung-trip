package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"lampung_trip/controller"
	"lampung_trip/middleware"
	"net/http"
	"time"
)

type Router struct {
	UserCont         controller.UserCont
	CustomerCont     controller.CustomerCont
	KaryawanCont     controller.KaryawanCont
	WisataCont       controller.WisataCont
	UlasanWisataCont controller.UlasanWisataCont
	VillaCont        controller.VillaCont
	UlasanVillaCont  controller.UlasanVillaCont
	FileCont         controller.FileCont

	SecretKey string
	Engine    *gin.Engine
}

func NewRouter(r Router) *Router {
	// CORS middleware
	r.Engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "ngrok-skip-browser-warning"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Route
	userGroup := r.Engine.Group("/user")
	userGroup.Use(middleware.ApiMiddleware(r.SecretKey))
	{
		userGroup.POST("/login-root", r.UserCont.LoginAdminKaryawan)
		userGroup.POST("/login", r.UserCont.Login)
		userGroup.POST("/admin/register", r.UserCont.RegisterAdmin)
		userGroup.POST("/karyawan/register", r.UserCont.RegisterKaryawan)
		userGroup.POST("/customer/register", r.UserCont.RegisterCustomer)

		userGroup.PUT("/password", r.UserCont.UpdatePw)

		userGroup.DELETE("/karyawan/:id", r.UserCont.DeleteKaryawan)
		userGroup.DELETE("/admin/:id", r.UserCont.DeleteAdmin)
	}

	custGroup := r.Engine.Group("/customer")
	custGroup.Use(middleware.ApiMiddleware(r.SecretKey))
	{
		custGroup.PUT("/", r.CustomerCont.UpdateData)

		custGroup.GET("/:id", r.CustomerCont.GetSingleByUserId)
		custGroup.GET("/", r.CustomerCont.GetAll)
	}

	karyawanGroup := r.Engine.Group("/karyawan")
	karyawanGroup.Use(middleware.ApiMiddleware(r.SecretKey))
	{
		karyawanGroup.PUT("/", r.KaryawanCont.UpdateData)

		karyawanGroup.GET("/:id", r.KaryawanCont.GetSingleByUserId)
		karyawanGroup.GET("/", r.KaryawanCont.GetAll)
	}

	wisataGroup := r.Engine.Group("/wisata")
	wisataGroup.Use(middleware.ApiMiddleware(r.SecretKey))
	{
		wisataGroup.POST("/", r.WisataCont.Create)

		wisataGroup.PUT("/", r.WisataCont.Update)

		wisataGroup.GET("/", r.WisataCont.GetAll)

		wisataGroup.DELETE("/:id", r.WisataCont.Delete)
	}

	ulasanWisataGroup := r.Engine.Group("/ulasan-wisata")
	ulasanWisataGroup.Use(middleware.ApiMiddleware(r.SecretKey))
	{
		ulasanWisataGroup.POST("/", r.UlasanWisataCont.Create)

		ulasanWisataGroup.DELETE("/", r.UlasanWisataCont.Delete)
	}

	villaGroup := r.Engine.Group("/villa")
	villaGroup.Use(middleware.ApiMiddleware(r.SecretKey))
	{
		villaGroup.POST("/", r.VillaCont.Create)

		villaGroup.PUT("/", r.VillaCont.Update)

		villaGroup.GET("/", r.VillaCont.GetAll)

		villaGroup.DELETE("/:id", r.VillaCont.Delete)
	}

	ulasanVillaGroup := r.Engine.Group("/ulasan-villa")
	ulasanVillaGroup.Use(middleware.ApiMiddleware(r.SecretKey))
	{
		ulasanVillaGroup.POST("/", r.UlasanVillaCont.Create)

		ulasanVillaGroup.DELETE("/", r.UlasanVillaCont.Delete)
	}

	fileGroup := r.Engine.Group("/file")
	fileGroup.Use(middleware.ApiMiddleware(r.SecretKey))
	{
		fileGroup.POST("/upload", r.FileCont.UploadFile)

		fileGroup.DELETE("/delete", r.FileCont.DeleteFile)
	}

	r.Engine.StaticFS("/assets", http.Dir("./assets"))

	return &Router{
		UserCont:         r.UserCont,
		CustomerCont:     r.CustomerCont,
		KaryawanCont:     r.KaryawanCont,
		WisataCont:       r.WisataCont,
		UlasanWisataCont: r.UlasanWisataCont,
		VillaCont:        r.VillaCont,
		UlasanVillaCont:  r.UlasanVillaCont,
		FileCont:         r.FileCont,

		SecretKey: r.SecretKey,
		Engine:    r.Engine,
	}
}
