package controller

import (
	"context"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/barrerajuanjose/usefulsearch/config"
	"github.com/barrerajuanjose/usefulsearch/domain"
	"github.com/barrerajuanjose/usefulsearch/marshaller"
	"github.com/barrerajuanjose/usefulsearch/service"
	"github.com/gin-gonic/gin"
)

type GetUsedCars interface {
	Get(ctx *gin.Context)
}

type getUsedCars struct {
	itemMarshaller marshaller.Item
	searchService  service.Search
	configurations map[string]*config.SiteConfiguration
}

func NewGetUsedCars(configurations map[string]*config.SiteConfiguration, itemMarshaller marshaller.Item, searchService service.Search) GetUsedCars {
	return &getUsedCars{
		itemMarshaller: itemMarshaller,
		searchService:  searchService,
		configurations: configurations,
	}
}

func (s getUsedCars) Get(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	siteId := c.Query("site_id")
	var siteConfig *config.SiteConfiguration

	if siteId == "" {
		for site, config := range s.configurations {
			if strings.Contains(c.Request.RequestURI, config.URI) {
				if site == "default" {
					siteId = "MLA"
				} else {
					siteId = site
				}
				siteConfig = config
				break
			}
		}
	}

	category := c.Query("category")
	if category == "" {
		category = siteConfig.CategoryId
	}

	brand := c.Query("brand")
	if brand == "clean" {
		brand = ""
	}

	stateId := c.Query("state_id")
	if stateId == "" {
		stateId = siteConfig.StateId
	} else if stateId == "clean" {
		stateId = ""
	}

	model := c.Query("model")

	searchChan := make(chan *domain.SearchResult, 1)
	viewChan := make(chan *marshaller.ModelDto, 1)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer close(searchChan)
		select {
		case <-ctx.Done():
			searchChan <- &domain.SearchResult{}
		default:
			searchResult := s.searchService.GetEndTodayItems(siteId, stateId, category, brand, model)
			searchChan <- searchResult
		}
	}()

	go func() {
		defer close(viewChan)
		wg.Wait()

		select {
		case <-ctx.Done():
			viewChan <- &marshaller.ModelDto{}
		default:
			itemsDomain := <-searchChan
			viewChan <- s.itemMarshaller.GetView(siteId, itemsDomain)
		}
	}()

	if c.GetHeader("accept") == "application/json" {
		c.JSON(http.StatusOK, <-viewChan)
	} else {
		c.HTML(http.StatusOK, "used_car.tmpl.html", <-viewChan)
	}
}
