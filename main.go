package main

import (
	"notes_server/internal/cmd"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("assets/config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	cmd.StartServer()
}
