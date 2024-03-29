package main

import (
	"github.com/gin-gonic/gin"
	"main.go/common/BaseController"
	"main.go/route"
)

func main() {

	mainroute := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = ioutil.Discard
	route.OnRoute(mainroute)
	mainroute.Use(BaseController.CommonController(), gin.Recovery())
	mainroute.Run(":80")

}
