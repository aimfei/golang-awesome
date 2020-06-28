package controller

import (
	"github.com/gin-gonic/gin"
	"golang-awesome/gin/model"
	"golang-awesome/gin/repos"
	"log"
)

func SaveAcc(ctx *gin.Context) {
	var acc model.Account
	if err := ctx.ShouldBindJSON(&acc); err != nil {
		log.Fatal(err.Error())
	}
	count := repos.SaveAcc(&acc)
	ctx.JSON(200, gin.H{
		"data": count,
	})
}
