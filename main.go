package main

import (
	"net/http"
	"os"

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

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "used_car.tmpl.html", nil)
	})
	router.GET("/autos-usados-mercadolibre-ultima-oportunidad", func(c *gin.Context) {
		c.HTML(http.StatusOK, "used_car.tmpl.html", nil)
	})

	router.Run(":" + port)
}
