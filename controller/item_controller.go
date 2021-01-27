package controller

import (
	"strconv"
	"sync"

	"github.com/barrerajuanjose/usefulsearch/domain"
	"github.com/barrerajuanjose/usefulsearch/marshaller"
	"github.com/barrerajuanjose/usefulsearch/service"
	"github.com/gin-gonic/gin"
)

type ItemController interface {
	Get(ctx *gin.Context)
}

type itemController struct {
	itemMarshaller marshaller.ItemMarshaller
	itemService    service.ItemService
	userService    service.UserService
}

func NewItemController(itemMarshaller marshaller.ItemMarshaller, itemService service.ItemService, userService service.UserService) ItemController {
	return &itemController{
		itemMarshaller: itemMarshaller,
		itemService:    itemService,
		userService:    userService,
	}
}

func (c itemController) Get(ctx *gin.Context) {
	itemId := ctx.Param("item_id")
	buyerId := ctx.Query("buyer_id")

	itemChan := make(chan *domain.Item, 1)
	sellerChan := make(chan *domain.User, 1)
	buyerChan := make(chan *domain.User, 1)
	viewChan := make(chan *marshaller.ItemDto, 1)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		defer close(buyerChan)
		if buyerInt, err := strconv.ParseInt(buyerId, 10, 32); err == nil {
			buyerChan <- c.userService.GetUserById(int32(buyerInt))
		} else {
			buyerChan <- nil
		}
	}()

	go func() {
		defer wg.Done()
		defer close(itemChan)
		item := c.itemService.GetItemById(itemId)
		itemChan <- item

		go func() {
			defer wg.Done()
			defer close(sellerChan)
			sellerChan <- c.userService.GetUserById(item.SellerId)
		}()

	}()

	go func() {
		defer close(viewChan)
		wg.Wait()

		itemDomain := <-itemChan
		sellerDomain := <-sellerChan
		buyerDomain := <-buyerChan

		viewChan <- c.itemMarshaller.GetView(itemDomain, sellerDomain, buyerDomain)
	}()

	ctx.JSON(200, <-viewChan)
}
