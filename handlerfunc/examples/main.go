package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dijkvy/geocentric/handlerfunc"
)

func main() {
	engine := gin.Default()
	engine.Use(handlerfunc.InjectTraceHandler)
	engine.Use(handlerfunc.InjectTimeOutHandler(1_000))
	engine.GET("/index", func(ctx *gin.Context) {
		ctx.JSON(200,
			gin.H{
				"status": "ok",
				"date":   time.Now().Format("2006-01-02 15:04:05.9999999"),
			},
		)
	})
	if err := engine.Run(); err != nil {
		log.Fatalln(err)
	}
}
