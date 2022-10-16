package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRoute() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())

	engine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})

	engine.POST("/webhook", func(context *gin.Context) {

	})

	return engine
}
