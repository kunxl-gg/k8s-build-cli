/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/kunxl-gg/lfx-lezgooo/cmd"
	"github.com/kunxl-gg/lfx-lezgooo/config"
)

func main() {
	config.Init()
	cmd.Execute()
}
