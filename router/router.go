package router

import (
	"ProxyHttp/controller"
	"github.com/gin-gonic/gin"
)

var Handle *gin.Engine

func InitRouter() {
	gin.SetMode(gin.ReleaseMode)

	Handle = gin.New()
	Handle.Use(gin.Recovery())

	Handle.GET("/proxy", controller.Proxy)
}
