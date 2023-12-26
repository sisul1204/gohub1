package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
	btsConfig "gohub/config"
	"gohub/pkg/config"
)

func init() {
	btsConfig.Initialize()
}

func main() {
	var env string
	flag.StringVar(&env, "env", "", "加载.env文件，如 --env=testing 加载的是.env.testing文件")
	flag.Parse()
	config.InitConfig(env)

	bootstrap.SetupLogger()

	router := gin.New()

	bootstrap.SetupDB()

	bootstrap.SetupRoute(router)

	err := router.Run(":" + config.GetString("app.port"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
