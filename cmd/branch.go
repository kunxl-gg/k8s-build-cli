/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
		var username string = os.Getenv("OBS_USERNAME");

		var password string = os.Getenv("OBS_PASSWORD");

		project_name := "home:kunxl.gg"
		package_name := "hello-world"
		var target_project string = "home:kunxl.gg:subproject"
		var target_package string  = "hello-world-again"
		var add_repositories int = 1

		base_url := "https://api.opensuse.org/source"

		url := base_url + "/" + project_name + "/" + package_name + "?cmd=branch&target_project=" + target_project + "&target_package=" + target_package + "&add_repositories=" + strconv.Itoa(add_repositories)

		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			panic(err)
		}

		req.Header.Set("Content-Type", "application/xml")
		req.SetBasicAuth(username, password)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		fmt.Println(resp)
		fmt.Printf("HTTP Response Status: %s\n", resp.Status)

	},
}

func init() {
	rootCmd.AddCommand(branchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// branchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// branchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
