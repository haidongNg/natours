package main

import (
	"fmt"
	"natours/db"
	"natours/pkg/logging"
	"natours/pkg/setting"
	"natours/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	db.Init()
	logging.Setup()
}

// @title Golang Gin API
// @version 1.0
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routerInit := routers.InitRouter()
	routerInit.SetTrustedProxies([]string{"localhost"})
	endPoint := fmt.Sprintf(":%d", 8080)
	routerInit.Run(endPoint)
}
