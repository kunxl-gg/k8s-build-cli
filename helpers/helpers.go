package helpers

import "log"

// helper function to check for errors
func CheckError(err error) {
	if err != nil {
		log.Fatal("There was an error: ", err)
	}
}