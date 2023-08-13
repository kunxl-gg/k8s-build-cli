/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		err := godotenv.Load()
		if err != nil {
			fmt.Errorf("Error loading .env file")
		}

		// getting the environment variables
		var username string = os.Getenv("OBS_USERNAME")
		var password string = os.Getenv("OBS_PASSWORD")

		fmt.Println(username, password)

		var base_url string = "https://api.opensuse.org/source"
		var project_name string = "home:kunxl.gg"

		// the complete url
		var url string = base_url + "/" + project_name + "?deleted=0&expand=0"

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Errorf("Error creating request")
		}

		req.SetBasicAuth(username, password)
		req.Header.Set("Content-Type", "application/xml")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Errorf("Error sending request", err)
		}

		fmt.Println(resp.Status, http.StatusText(resp.StatusCode))
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Errorf("Error reading response body", err)
		}

		log.Default().Println(string(body))
		defer resp.Body.Close()

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}