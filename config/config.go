package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {

	viper.SetConfigName("config")
	viper.AddConfigPath("../")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
}
