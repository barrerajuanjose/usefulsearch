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
	pageMarshaller := marshaller.NewPage()

	getIndex := controller.NewGetIndex(configurations, pageMarshaller)
	getUsedCars := controller.NewGetUsedCars(configurations, pageMarshaller, service.NewSearch())

	router.GET("/", getIndex.Get)

	for _, config := range configurations {
		router.GET(config.URI, getUsedCars.Get)
	}

	router.Run(":" + port)
}

func buildSiteConfiguration() []*config.SiteConfiguration {
	return []*config.SiteConfiguration{
		{
			SiteId:          "MLA",
			Name:            "Argentina",
			PageTitle:       "Autos Usados Mercado Libre Argentina Útima Oportunidad",
			PageDescription: "Publicaciones de autos usados en Argentina que finalizan hoy, ideales para hacer una oferta!.",
			Title:           "Autos usados en Mercado Libre! Ultima oportunidad para comprarlos",
			BaseUrl:         getBaseUrl(),
			URI:             "/autos-usados-mercadolibre-ultima-oportunidad-argentina/",
			CategoryId:      "MLA1744",
			StateId:         "TUxBUENBUGw3M2E1",
			IsDefault:       true,
		},
		{
			SiteId:          "MLB",
			Name:            "Brasil",
			PageTitle:       "Carros Usados Mercado Livre Brasil",
			PageDescription: "Publicaciones de autos usados en Brasil que finalizan hoy, ideales para hacer una oferta!.",
			Title:           "Autos usados en Mercado Libre! Ultima oportunidad para comprarlos",
			BaseUrl:         getBaseUrl(),
			URI:             "/carros-mercadolibre-ultima-oportunidad-brasil/",
			CategoryId:      "MLB1744",
		},
		{
			SiteId:          "MLM",
			Name:            "México",
			PageTitle:       "Carros Usados Mercado Libre México Útima Oportunidad",
			PageDescription: "Publicaciones de autos usados en México que finalizan hoy, ideales para hacer una oferta!.",
			Title:           "Autos usados en Mercado Libre! Ultima oportunidad para comprarlos",
			BaseUrl:         getBaseUrl(),
			URI:             "/autos-usados-mercadolibre-ultima-oportunidad-mexico/",
			CategoryId:      "MLM1744",
		},
		{
			SiteId:          "MLC",
			Name:            "Chile",
			PageTitle:       "Autos Usados Mercado Libre Chile Útima Oportunidad",
			PageDescription: "Publicaciones de autos usados en Chile que finalizan hoy, ideales para hacer una oferta!.",
			Title:           "Autos usados en Mercado Libre! Ultima oportunidad para comprarlos",
			BaseUrl:         getBaseUrl(),
			URI:             "/autos-usados-mercadolibre-ultima-oportunidad-chile/",
			CategoryId:      "MLC1744",
		},
		{
			SiteId:          "MLU",
			Name:            "Uruguay",
			PageTitle:       "Autos Usados Mercado Libre Uruguay Útima Oportunidad",
			PageDescription: "Publicaciones de autos usados en Uruguay que finalizan hoy, ideales para hacer una oferta!.",
			Title:           "Autos usados en Mercado Libre! Ultima oportunidad para comprarlos",
			BaseUrl:         getBaseUrl(),
			URI:             "/autos-usados-mercadolibre-ultima-oportunidad-uruguay/",
			CategoryId:      "MLU1744",
		},
		{
			SiteId:          "MCO",
			Name:            "Colombia",
			PageTitle:       "Carros Usados Mercado Libre Colombia Útima Oportunidad",
			PageDescription: "Publicaciones de autos usados en Colombia que finalizan hoy, ideales para hacer una oferta!.",
			Title:           "Carros usados en Mercado Libre! Ultima oportunidad para comprarlos",
			BaseUrl:         getBaseUrl(),
			URI:             "/carros-usados-mercadolibre-ultima-oportunidad-colombia/",
			CategoryId:      "MCO1744",
		},
		{
			SiteId:          "MLV",
			Name:            "Venezuela",
			PageTitle:       "Carros Usados Mercado Libre Venezuela Útima Oportunidad",
			PageDescription: "Publicaciones de autos usados en Venezuela que finalizan hoy, ideales para hacer una oferta!.",
			Title:           "Carros usados en Mercado Libre! Ultima oportunidad para comprarlos",
			BaseUrl:         getBaseUrl(),
			URI:             "/carros-usados-mercadolibre-ultima-oportunidad-venezuela/",
			CategoryId:      "MLV1744",
		},
	}
}

func getBaseUrl() string {
	port := os.Getenv("PORT")

	if port == "" {
		return "http://localhost:8080"
	}

	return "https://usefulsearch.herokuapp.com"
}
