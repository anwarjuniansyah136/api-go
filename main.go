package main

import (
	"api/helper"
	"api/modules"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := helper.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load Config", err)
	}

	if config.LOG_FILE == "on" {
		helper.SetUpLogOutput()
	}

	gin.SetMode(config.GIN_MODE)

	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", config.ALLOW_ORIGIN)
		ctx.Writer.Header().Set("Access-Control-Allow-Headers","Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	})

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "APM API Service is running",
		})
	})

	db := helper.OpenDB(config.DB, config.SCHEMA, "v1")

	db = db.Debug()

	db = db.Set("gorm:table_options", "schema:aset")

	if err := db.Exec("CREATE SCHEMA IF NOT EXISTS aset").Error; err != nil {
		log.Fatal("failed create schema:", err)
	}

	if err := db.Exec("SET search_path TO aset").Error; err != nil {
		log.Fatal("failed set search path:", err)
	}

	versions := modules.NewVersions(config, r, db)
	versions.Run()

	if err := r.Run(":" + config.PORT); err != nil{
		log.Fatal(err)
	}
}