package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
)

func main() {
	//r := gin.Default()
	router := gin.New()

	bootstrap.SetupRoute(router)

	err := router.Run(":9000")

	if err != nil {
		fmt.Println(err.Error())
	}
}
