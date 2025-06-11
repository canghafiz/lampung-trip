package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"lampung_trip/app"
	"lampung_trip/controller"
	"lampung_trip/helper"
	"lampung_trip/model/domain"
	"lampung_trip/model/repository"
	"lampung_trip/model/service"
	"os"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	// Load .env file
	err := godotenv.Load(".env")
	helper.FatalError(err)

	// Other Env
	secretKey := os.Getenv("SECRET_KEY")
	port := os.Getenv("APP_PORT")        // e.g., ":3000"
	appDomain := os.Getenv("APP_DOMAIN") // http://localhost

	// Database config
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	// Open DB Connection
	db := app.OpenConnection(dbUser, dbPass, dbHost+":"+dbPort, dbName)

	// Initialize Repository
	userRepo := repository.NewUserRepoImpl(db)
	custRepo := repository.NewCustomerRepoImpl(db)
	karyawanRepo := repository.NewKaryawanRepoImpl(db)
	wisataRepo := repository.NewWisataRepoImpl(db)
	ulasanWisataRepo := repository.NewUlasanWisataRepoImpl(db)
	villaRepo := repository.NewVillaRepoImpl(db)
	ulasanVillaRepo := repository.NewUlasanVillaRepoImpl(db)

	// Initialize Service
	fileServ := service.NewFileServImpl(domain.File{}, appDomain)
	userServ := service.NewUserServImpl(userRepo, secretKey)
	custServ := service.NewCustomerServImpl(custRepo)
	karyawanServ := service.NewKaryawanServImpl(karyawanRepo)
	wisataServ := service.NewWisataServImpl(wisataRepo)
	ulasanWisataServ := service.NewUlasanWisataServImpl(ulasanWisataRepo)
	villaServ := service.NewVillaServImpl(villaRepo)
	ulasanVillaServ := service.NewUlasanVillaServImpl(ulasanVillaRepo)

	// Initialize Controller
	userCont := controller.NewUserContImpl(userServ)
	custCont := controller.NewCustomerContImpl(custServ)
	karyawanCont := controller.NewKaryawanContImpl(karyawanServ)
	wisataCont := controller.NewWisataContImpl(wisataServ)
	ulasanWisataCont := controller.NewUlasanWisataContImpl(ulasanWisataServ)
	villaCont := controller.NewVillaContImpl(villaServ)
	ulasanVillaCont := controller.NewUlasanVillaContImpl(ulasanVillaServ)
	fileCont := controller.NewFileContImpl(fileServ)

	// Setup router
	engine := gin.Default()
	routerParent := app.Router{
		UserCont:         userCont,
		CustomerCont:     custCont,
		KaryawanCont:     karyawanCont,
		WisataCont:       wisataCont,
		UlasanWisataCont: ulasanWisataCont,
		VillaCont:        villaCont,
		UlasanVillaCont:  ulasanVillaCont,
		FileCont:         fileCont,

		SecretKey: secretKey,
		Engine:    engine,
	}
	router := app.NewRouter(routerParent)

	// Run Server
	if port == "" {
		port = ":3000" // Default fallback
	}
	err = router.Engine.Run(port)
	helper.FatalError(err)
}
