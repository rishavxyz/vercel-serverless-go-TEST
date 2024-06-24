package handler

import (
	"github.com/gin-gonic/gin"
)

var app = gin.Default()

func GetHello(f func(ctx *gin.Context) (string, error)) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.JSON(200, gin.H{"Hello": name})
	}
}
