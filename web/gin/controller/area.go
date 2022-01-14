package controller

import (
	"github.com/gin-gonic/gin"
	"golang-awesome/web/gin/base"
	"golang-awesome/web/gin/repos"
)

func AreaList(ctx *gin.Context) {
	code := ctx.Param("code")
	reaList := repos.FindList(code)
	ctx.JSON(200, base.Success(reaList))
}
