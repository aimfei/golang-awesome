package route

import (
	"github.com/gin-gonic/gin"
	controller2 "golang-awesome/web/gin/controller"
)

func AreaRoute(route *gin.Engine) {
	areaController := route.Group("/v1/area")
	{
		areaController.GET("/getArea", controller2.AreaList)
		areaController.GET("/getArea/:code", controller2.AreaList)
	}
}
