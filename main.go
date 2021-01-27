package main

import (
	"net/http"
	"os"

	"github.com/barrerajuanjose/usefulsearch/controller"
	"github.com/barrerajuanjose/usefulsearch/marshaller"
	"github.com/barrerajuanjose/usefulsearch/service"
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

	itemController := controller.NewItemController(marshaller.NewItemMarshaller(), service.NewItemService(), service.NewUserService())

	router.GET("/", itemController.Get)

	router.GET("/autos-usados-mercadolibre-ultima-oportunidad", func(c *gin.Context) {
		c.HTML(http.StatusOK, "used_car.tmpl.html", nil)
	})

	router.Run(":" + port)
}
