package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func getHello(r *gin.RouterGroup) {
	r.GET("/hello/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.JSON(http.StatusOK, gin.H{"hello": name})
	})
}

func init() {
	app = gin.New()
	r := app.Group("/api")

	getHello(r)
}

// vercel server less injection
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
