package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

var (
	Username string
	Password string
	BaseUrl string
)

func Init() {

	// fetch the env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("There is an error: ", err)
	}


	// set the global variables
	Username = os.Getenv("OBS_USERNAME")
	Password = os.Getenv("OBS_PASSWORD")
	BaseUrl = os.Getenv("BASE_URL")

}
