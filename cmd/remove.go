/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/kunxl-gg/lfx-lezgooo/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		// loading the dotenv file
		err := godotenv.Load()
		if err != nil {
			panic("Error loading .env file")
		}

		config.InitConfig()

		// getting the environment variables
		var username string = os.Getenv("OBS_USERNAME")
		var password string = viper.GetString("OBS_PASSWORD")

		var base_url string = "https://api.opensensemap.org"
		var project_name string = args[0]
		fmt.Println(project_name)

		// the complete url
		var url string = base_url + "/source/" + project_name 

		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			panic(err)
		}

		req.SetBasicAuth(username, password)
		req.Header.Set("Content-Type", "application/xml")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		fmt.Println(resp)
		fmt.Println(resp.Status)
		fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
