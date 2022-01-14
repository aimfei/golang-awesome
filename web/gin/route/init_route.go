package route

import "github.com/gin-gonic/gin"

func InitRoute(server *gin.Engine) {
	AreaRoute(server)
}
