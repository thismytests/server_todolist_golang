package appconfig

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

// init is invoked before main()
type AppConfig struct {
	APP_PORT int

	POSTGRES_DATABASE_NAME string

	POSTGRES_DATABASE_HOST     string
	POSTGRES_DATABASE_PORT     int
	POSTGRES_DATABASE_USERNAME string
	POSTGRES_DATABASE_PASSWORD string
}

func (t AppConfig) String() string {
	return fmt.Sprintf(
		"'APP port' is: %d ;\n "+
			"'DATABASE name' is: %s ;\n  "+
			"'DATABASE host' is: %s ;\n  "+
			"'DATABASE port' is: %d ;\n  "+
			"'DATABASE userName' is: %s ;\n  "+
			"'DATABASE password' is: %s ;\n  ",
		t.APP_PORT,
		t.POSTGRES_DATABASE_NAME,
		t.POSTGRES_DATABASE_HOST,
		t.POSTGRES_DATABASE_PORT,
		t.POSTGRES_DATABASE_NAME,
		t.POSTGRES_DATABASE_PASSWORD,
	)
}

// todo ... Mykolay Lytvyn ... should saving without saving into the global variables
func ReadEnvFile(file string) AppConfig {
	// loads values from .env into the system
	appConfig := AppConfig{}

	if err := godotenv.Load(file); err != nil {
		log.Print("No .env file found")
	}

	appPort, isExistsAppPort := os.LookupEnv("APP_PORT")
	dataBaseName, isExistsDataBaseName := os.LookupEnv("POSTGRES_DATABASE_NAME")

	dataBaseHost, isExistsDataBasehost := os.LookupEnv("POSTGRES_DATABASE_HOST")
	dataBasePort, isExistsDataBasePort := os.LookupEnv("POSTGRES_DATABASE_PORT")
	dataBaseUserName, isExistsDataBaseUserName := os.LookupEnv("POSTGRES_DATABASE_USERNAME")
	dataBasePassword, isExistsDataBasePassword := os.LookupEnv("POSTGRES_DATABASE_PASSWORD")

	if i, err := strconv.Atoi(appPort); err == nil && isExistsAppPort {
		appConfig.APP_PORT = i
	}

	if isExistsDataBaseName {
		appConfig.POSTGRES_DATABASE_NAME = dataBaseName
	}

	if isExistsDataBasehost {
		appConfig.POSTGRES_DATABASE_HOST = dataBaseHost
	}

	if i, err := strconv.Atoi(dataBasePort); err == nil && isExistsDataBasePort {
		appConfig.POSTGRES_DATABASE_PORT = i
	}
	if isExistsDataBaseUserName {
		appConfig.POSTGRES_DATABASE_USERNAME = dataBaseUserName
	}
	if isExistsDataBasePassword {
		appConfig.POSTGRES_DATABASE_PASSWORD = dataBasePassword
	}

	return appConfig
}

func ParseEnvFile() {

}
