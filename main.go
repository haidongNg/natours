package main

import (
	"fmt"
	"natours/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.DebugMode)
}

func main() {
	routerInit := routers.InitRouter()
	routerInit.SetTrustedProxies([]string{"localhost"})
	endPoint := fmt.Sprintf(":%d", 8080)

	routerInit.Run(endPoint)
}
