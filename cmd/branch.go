/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"github.com/kunxl-gg/lfx-lezgooo/config"
	"github.com/kunxl-gg/lfx-lezgooo/helpers"
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

		// declaring parameters for the url
		project_name := "home:kunxl.sol"
		package_name := "hello-world"
		var target_project string = "home:kunxl.sol:subproject"
		var target_package string  = "hello-world-again"
		var add_repositories int = 1

		// final url
		url := config.BaseUrl + "source/?cmd=branch&project=" + project_name + "&package=" + package_name + "&target_project=" + target_project + "&target_package=" + target_package + "&add_repositories=" + strconv.Itoa(add_repositories)

		fmt.Println(url)
		// response of the request
		resp := helpers.APICall("POST", url)

		// printing the responseˀ
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
