package router

import (
	"log"
	"github.com/joho/godotenv"
	"os"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/e-capture/ECMVinculacion/controllers/ImportconfigController"
	"github.com/e-capture/ECMVinculacion/controllers/VinculacionController"
)

var _port string
var _serviceName string

func init()  {
	if err := godotenv.Load(); err != nil {
		panic(err)
		log.Fatal("Error loading .env file")
	}
	_port = os.Getenv("APP_PORT")
	_serviceName = os.Getenv("APP_SERVICE_NAME")
}

func StartService()  {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.POST("/api/v1/import", ImportconfigController.Import)
	e.POST("/api/v1/importconfig", ImportconfigController.ImportConfig)
	e.GET("/api/v1/vinculacion/index", VinculacionController.Index)
	e.GET("/api/v1/vinculacion/ShowTypeDoc", VinculacionController.ShowTypeDoc)
	e.GET("/api/v1/vinculacion/ShowSex", VinculacionController.ShowSex)
	e.GET("/api/v1/vinculacion/ShowEstadoCivil", VinculacionController.ShowEstadoCivil)
	e.POST("api/v1/vinculacion/create", VinculacionController.Create)
	e.Logger.Fatal(e.Start(":"+_port))
}






