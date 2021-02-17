package controller

import (
	"net/http"

	"github.com/barrerajuanjose/usefulsearch/config"
	"github.com/barrerajuanjose/usefulsearch/marshaller"
	"github.com/gin-gonic/gin"
)

type GetIndex interface {
	Get(ctx *gin.Context)
}

type getIndex struct {
	siteConfigurations []*config.SiteConfiguration
	pageMarshaller     marshaller.Page
}

func NewGetIndex(siteConfigurations []*config.SiteConfiguration, pageMarshaller marshaller.Page) GetIndex {
	return &getIndex{
		siteConfigurations: siteConfigurations,
		pageMarshaller:     pageMarshaller,
	}
}

func (s getIndex) Get(c *gin.Context) {
	var defaultSiteConfiguration *config.SiteConfiguration

	for _, siteConfiguration := range s.siteConfigurations {
		if siteConfiguration.IsDefault {
			defaultSiteConfiguration = siteConfiguration
			break
		}
	}

	pageModel := s.pageMarshaller.GetIndex(defaultSiteConfiguration, s.siteConfigurations)

	if c.GetHeader("accept") == "application/json" {
		c.JSON(http.StatusOK, pageModel)
	} else {
		c.HTML(http.StatusOK, "index.tmpl.html", pageModel)
	}
}
