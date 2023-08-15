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
	Long: `Branch out from the main project to a subproject`,

	Run: func(cmd *cobra.Command, args []string) {

		// declaring parameters for the url
		project_name := "home:kunxl.sol"
		package_name := "hello-world"
		var target_project string = "home:kunxl.sol:subproject"
		var target_package string  = "hello-world-again"
		var add_repositories int = 1

		// final url
		url := config.BaseUrl + "source/?cmd=branch&project=" + project_name + "&package=" + package_name + "&target_project=" + target_project + "&target_package=" + target_package + "&add_repositories=" + strconv.Itoa(add_repositories)

		// response of the request
		resp := helpers.APICall("POST", url)

		// printing the responseˀ
		fmt.Printf("HTTP Response Status: %s\n", resp.Status)

	},
}

func init() {
	rootCmd.AddCommand(branchCmd)
}
