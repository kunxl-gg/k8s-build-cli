/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/kunxl-gg/lfx-lezgooo/helpers"
	"github.com/spf13/cobra"
)

// addFileCmd represents the addFile command
var addFileCmd = &cobra.Command{
	Use:   "addFile",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		
		base_url := "https://api.opensuse.org/source"
		project_name := "home:kunxl.gg"
		package_name := "hello-world"
		file_name := args[0]

		// data := []byte(`{
		// 	"project_name": "home:kunxl.gg",
		// 	"package_name": "hello-world",
		// 	"file_name": "root.go", 
		// }`)

		url := base_url + "/" + project_name + "/" + package_name + "/" + file_name

		// making an api call
		resp := helpers.APICall("PUT", url)
		fmt.Printf("HTTP Response Status: %s\n", resp.Status)
	
	},
}

func init() {
	rootCmd.AddCommand(addFileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addFileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addFileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
