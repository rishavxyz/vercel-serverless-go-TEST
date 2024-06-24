package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func init() {
	app = gin.New()
	r := app.Group("/api")
	getHello(r)
}

// vercel handler
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}

// endpoints
func getHello(r *gin.RouterGroup) {
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"Hello": "World!"})
	})
}
