package controller

import (
	"net/http"
	"sync"

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
}

func NewGetUsedCars(itemMarshaller marshaller.Item, searchService service.Search) GetUsedCars {
	return &getUsedCars{
		itemMarshaller: itemMarshaller,
		searchService:  searchService,
	}
}

func (c getUsedCars) Get(ctx *gin.Context) {
	siteId := ctx.Query("site_id")

	if siteId == "" {
		siteId = "MLA"
	}

	stateId := ctx.Query("state_id")
	if stateId == "" {
		stateId = "TUxBUENBUGw3M2E1"
	}

	category := ctx.Query("category")
	if category == "" {
		category = "MLA1744"
	}

	brand := ctx.Query("brand")
	model := ctx.Query("model")

	searchChan := make(chan *domain.SearchResult, 1)
	viewChan := make(chan *marshaller.ModelDto, 1)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer close(searchChan)
		searchResult := c.searchService.GetEndTodayItems(siteId, stateId, category, brand, model)
		searchChan <- searchResult
	}()

	go func() {
		defer close(viewChan)
		wg.Wait()

		itemsDomain := <-searchChan
		viewChan <- c.itemMarshaller.GetView(itemsDomain)
	}()

	ctx.HTML(http.StatusOK, "used_car.tmpl.html", <-viewChan)
}
