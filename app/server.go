package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppPort string
	AppEnv  string
}

func (server *Server) Initialize(appConfig AppConfig) {
	fmt.Println("Selamat Datang " + appConfig.AppName)
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s \n", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func Run() {
	var server = Server{}
	var appConfig = AppConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error Loading Env File")
	}

	appConfig.AppName = os.Getenv("APP_NAME")
	appConfig.AppEnv = os.Getenv("APP_ENV")
	appConfig.AppPort = os.Getenv("APP_PORT")
	server.Initialize(appConfig)
	server.Run(":" + appConfig.AppPort)
}
