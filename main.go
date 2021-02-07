package main

import (
	"os"

	"github.com/barrerajuanjose/usefulsearch/config"
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

	configurations := buildSiteConfiguration()

	getUsedCars := controller.NewGetUsedCars(configurations, marshaller.NewItem(), service.NewSearch())

	router.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})

	for _, config := range configurations {
		router.GET(config.URI, getUsedCars.Get)
	}

	router.Run(":" + port)
}

func buildSiteConfiguration() map[string]*config.SiteConfiguration {
	return map[string]*config.SiteConfiguration{
		"MLA": {
			PageTitle:       "Autos Usados Mercado Libre Argentina Útima Oportunidad",
			PageDescription: "Publicaciones de autos usados en Argentina que finalizan hoy, ideales para hacer una oferta!.",
			Title:           "Autos usados en Mercado Libre! Ultima oportunidad para comprarlos",
			Canonical:       "https://usefulsearch.herokuapp.com/autos-usados-mercadolibre-ultima-oportunidad-argentina/",
			BaseUrl:         "https://usefulsearch.herokuapp.com",
			URI:             "/autos-usados-mercadolibre-ultima-oportunidad-argentina/",
			CategoryId:      "MLA1744",
			StateId:         "TUxBUENBUGw3M2E1",
		},
		"MLB": {
			PageTitle:       "Carros Usados Mercado Livre Brasil",
			PageDescription: "Publicaciones de autos usados en Brasil que finalizan hoy, ideales para hacer una oferta!.",
			Title:           "Autos usados en Mercado Libre! Ultima oportunidad para comprarlos",
			Canonical:       "https://usefulsearch.herokuapp.com/carros-mercadolibre-ultima-oportunidad-brasil/",
			BaseUrl:         "https://usefulsearch.herokuapp.com",
			URI:             "/carros-mercadolibre-ultima-oportunidad-brasil/",
			CategoryId:      "MLB1744",
		},
		"MLM": {
			PageTitle:       "Carros Usados Mercado Libre México Útima Oportunidad",
			PageDescription: "Publicaciones de autos usados en México que finalizan hoy, ideales para hacer una oferta!.",
			Title:           "Autos usados en Mercado Libre! Ultima oportunidad para comprarlos",
			Canonical:       "https://usefulsearch.herokuapp.com/autos-usados-mercadolibre-ultima-oportunidad-mexico/",
			BaseUrl:         "https://usefulsearch.herokuapp.com",
			URI:             "/autos-usados-mercadolibre-ultima-oportunidad-mexico/",
			CategoryId:      "MLM1744",
		},
		"MLC": {
			PageTitle:       "Autos Usados Mercado Libre Chile Útima Oportunidad",
			PageDescription: "Publicaciones de autos usados en Chile que finalizan hoy, ideales para hacer una oferta!.",
			Title:           "Autos usados en Mercado Libre! Ultima oportunidad para comprarlos",
			Canonical:       "https://usefulsearch.herokuapp.com/autos-usados-mercadolibre-ultima-oportunidad-chile/",
			BaseUrl:         "https://usefulsearch.herokuapp.com",
			URI:             "/autos-usados-mercadolibre-ultima-oportunidad-chile/",
			CategoryId:      "MLC1744",
		},
		"MLU": {
			PageTitle:       "Autos Usados Mercado Libre Uruguay Útima Oportunidad",
			PageDescription: "Publicaciones de autos usados en Uruguay que finalizan hoy, ideales para hacer una oferta!.",
			Title:           "Autos usados en Mercado Libre! Ultima oportunidad para comprarlos",
			Canonical:       "https://usefulsearch.herokuapp.com/autos-usados-mercadolibre-ultima-oportunidad-uruguay/",
			BaseUrl:         "https://usefulsearch.herokuapp.com",
			URI:             "/autos-usados-mercadolibre-ultima-oportunidad-uruguay/",
			CategoryId:      "MLU1744",
		},
		"MCO": {
			PageTitle:       "Autos Usados Mercado Libre Colombia Útima Oportunidad",
			PageDescription: "Publicaciones de autos usados en Colombia que finalizan hoy, ideales para hacer una oferta!.",
			Title:           "Autos usados en Mercado Libre! Ultima oportunidad para comprarlos",
			Canonical:       "https://usefulsearch.herokuapp.com/autos-usados-mercadolibre-ultima-oportunidad-colombia/",
			BaseUrl:         "https://usefulsearch.herokuapp.com",
			URI:             "/autos-usados-mercadolibre-ultima-oportunidad-colombia/",
			CategoryId:      "MCO1744",
		},
		"MLV": {
			PageTitle:       "Autos Usados Mercado Libre Venezuela Útima Oportunidad",
			PageDescription: "Publicaciones de autos usados en Venezuela que finalizan hoy, ideales para hacer una oferta!.",
			Title:           "Autos usados en Mercado Libre! Ultima oportunidad para comprarlos",
			Canonical:       "https://usefulsearch.herokuapp.com/autos-usados-mercadolibre-ultima-oportunidad-venezuela/",
			BaseUrl:         "https://usefulsearch.herokuapp.com",
			URI:             "/autos-usados-mercadolibre-ultima-oportunidad-venezuela/",
			CategoryId:      "MLV1744",
		},
		"default": {
			PageTitle:       "Autos Usados Mercado Libre Útima Oportunidad",
			PageDescription: "Publicaciones de autos usados que finalizan hoy, ideales para hacer una oferta!.",
			Title:           "Autos usados en Mercado Libre! Ultima oportunidad para comprarlos",
			Canonical:       "https://usefulsearch.herokuapp.com/autos-usados-mercadolibre-ultima-oportunidad/",
			BaseUrl:         "https://usefulsearch.herokuapp.com",
			URI:             "/autos-usados-mercadolibre-ultima-oportunidad/",
			CategoryId:      "MLA1744",
			StateId:         "TUxBUENBUGw3M2E1",
		},
	}
}
