package helpers

import (
	"log"
	"net/http"
	"github.com/kunxl-gg/lfx-lezgooo/config"
)

// helper function to check for errors
func CheckError(err error) {
	if err != nil {
		log.Fatal("There was an error: ", err)
	}
}

// function to make an api call
func APICall(method string, url string) *http.Response {
	
	// get credentials
	username := config.Username
	password := config.Password

	// creating a new request
	req, err := http.NewRequest(method, url, nil)
	CheckError(err)

	// setting the basic auth headers
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/xml")

	// creating a new client
	client := &http.Client{}
	resp, err := client.Do(req)
	CheckError(err)

	// return resp
	return resp

}