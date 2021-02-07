package main

import (
	"os"

	"github.com/barrerajuanjose/usefulsearch/controller"
	"github.com/barrerajuanjose/usefulsearch/marshaller"
	"github.com/barrerajuanjose/usefulsearch/service"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	getUsedCars := controller.NewGetUsedCars(marshaller.NewItem(), service.NewSearch())

	router.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})

	router.GET("/autos-usados-mercadolibre-ultima-oportunidad/", getUsedCars.Get)

	router.Run(":" + port)
}
