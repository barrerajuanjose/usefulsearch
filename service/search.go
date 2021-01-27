package service

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/barrerajuanjose/usefulsearch/domain"
)

type Search interface {
	GetEndTodayItems(siteId string, stateId string, query string) []*domain.Item
}

type search struct {
}

type searchItemResponse struct {
	Id         string  `json:"id,omitempty"`
	Title      string  `json:"title,omitempty"`
	Price      float32 `json:"price,omitempty"`
	CurrencyId string  `json:"currency_id,omitempty"`
	Thumbnail  string  `json:"thumbnail,omitempty"`
	Permalink  string  `json:"permalink,omitempty"`
	StopTime   string  `json:"stop_time,omitempty"`
}

type searchResponse struct {
	Results []searchItemResponse `json:"results,omitempty"`
}

func NewSearch() Search {
	return &search{}
}

func (*search) GetEndTodayItems(siteId string, stateId string, query string) []*domain.Item {
	response, err := http.Get(fmt.Sprintf("https://api.mercadolibre.com/sites/%s/search?limit=50&until=today&state=%s&category=%s", siteId, stateId, query))
	if err != nil {
		return nil
	}

	defer response.Body.Close()

	respBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil
	}

	var searchResponse searchResponse
	_ = json.Unmarshal(respBody, &searchResponse)

	var items []*domain.Item

	for _, itemResponse := range searchResponse.Results {
		items = append(items, &domain.Item{
			Id:         itemResponse.Id,
			Title:      itemResponse.Title,
			Price:      itemResponse.Price,
			CurrencyId: itemResponse.CurrencyId,
			Thumbnail:  itemResponse.Thumbnail,
			Permalink:  itemResponse.Permalink,
			StopTime:   itemResponse.StopTime,
		})
	}

	return items
}
