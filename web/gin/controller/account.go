package controller

import (
	"github.com/gin-gonic/gin"
	"golang-awesome/web/gin/base"
	"golang-awesome/web/gin/model"
	"golang-awesome/web/gin/repos"
	"log"
)

func SaveAcc(ctx *gin.Context) {
	var acc model.Account
	if err := ctx.ShouldBindJSON(&acc); err != nil {
		log.Fatal(err.Error())
	}
	count := repos.SaveAcc(&acc)
	ctx.JSON(200, base.Success(count))
}
