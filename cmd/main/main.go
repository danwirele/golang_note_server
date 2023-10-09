package main

import (
	"C"
	"notes_server/internal/cmd"

	"github.com/spf13/viper"
)
import "github.com/gin-gonic/gin"

func main() {
	viper.SetConfigFile("assets/config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if mode := viper.GetString("GIN_MODE"); mode == "RELEASE" {
		gin.SetMode(gin.ReleaseMode)

	}

	if err := cmd.StartServer(); err != nil {
		panic(err)
	}

}
