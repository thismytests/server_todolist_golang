package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"main/appconfig"
	repository "main/db/postgressql"
	"main/routes"
	"net/http"
)

const envFilePath = "../env/.env.dev"

func main() {
	config := getAppConfig()

	// Log appconfig info
	fmt.Println("appconfig", config)

	// create tables
	repository.CreateTodoListTable(repository.NewRepository())

	// server init
	initServer(config)
}

func getAppConfig() appconfig.AppConfig {
	return appconfig.ReadEnvFile(envFilePath)
}

func initServer(config appconfig.AppConfig) {
	port := fmt.Sprintf("%s%d", ":", config.APP_PORT)

	// register mux routes
	routes.RegisterMuxRoutes{}.Init()

	fmt.Println("Server is listening...")

	// start server
	log.Fatal(http.ListenAndServe(port, nil))

}
