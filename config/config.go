package config

import (
	"os"
	"github.com/joho/godotenv"
	"github.com/kunxl-gg/lfx-lezgooo/helpers"
)

var (
	Username string
	Password string
)

func Init() {

	// fetch the env variables
	err := godotenv.Load()
	helpers.CheckError(err)

	// set the global variables
	Username = os.Getenv("OBS_USERNAME")
	Password = os.Getenv("OBS_PASSWORD")

}
