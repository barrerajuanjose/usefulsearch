package controller

import (
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
	siteId := ctx.Param("site_id")

	if siteId == "" {
		siteId = "MLA"
	}

	stateId := ctx.Param("state_id")
	if stateId == "" {
		stateId = "TUxBUENBUGw3M2E1"
	}

	query := ctx.Param("query")
	if query == "" {
		query = "MLA1744"
	}

	searchChan := make(chan []*domain.Item, 1)
	viewChan := make(chan []*marshaller.ItemDto, 1)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer close(searchChan)
		items := c.searchService.GetEndTodayItems(siteId, stateId, query)
		searchChan <- items
	}()

	go func() {
		defer close(viewChan)
		wg.Wait()

		itemsDomain := <-searchChan
		viewChan <- c.itemMarshaller.GetView(itemsDomain)
	}()

	ctx.JSON(200, <-viewChan)
	//ctx.HTML(http.StatusOK, "used_car.tmpl.html", nil)
}
