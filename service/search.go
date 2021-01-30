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
	GetEndTodayItems(siteId string, stateId string, query string, brand string, model string) *domain.SearchResult
}

type search struct {
}

type searchItemResponse struct {
	Id         string                         `json:"id,omitempty"`
	Title      string                         `json:"title,omitempty"`
	Price      float64                        `json:"price,omitempty"`
	CurrencyId string                         `json:"currency_id,omitempty"`
	Thumbnail  string                         `json:"thumbnail,omitempty"`
	Permalink  string                         `json:"permalink,omitempty"`
	StopTime   string                         `json:"stop_time,omitempty"`
	Attributes []*searchItemAttributeResponse `json:"attributes,omitempty"`
}

type searchItemAttributeResponse struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	ValueName string `json:"value_name,omitempty"`
}

type searchFilterValueResponse struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type searchFilterResponse struct {
	Id     string                       `json:"id,omitempty"`
	Name   string                       `json:"name,omitempty"`
	Values []*searchFilterValueResponse `json:"values,omitempty"`
}

type searchResponse struct {
	Results          []*searchItemResponse   `json:"results,omitempty"`
	Filters          []*searchFilterResponse `json:"filters,omitempty"`
	AvailableFilters []*searchFilterResponse `json:"available_filters,omitempty"`
}

func NewSearch() Search {
	return &search{}
}

func (*search) GetEndTodayItems(siteId string, stateId string, category string, brand string, model string) *domain.SearchResult {
	brandParam := ""
	if brand != "" {
		brandParam = fmt.Sprintf("&brand=%s", brand)
	}

	modelParam := ""
	if model != "" {
		modelParam = fmt.Sprintf("&model=%s", model)
	}

	stateParam := ""
	if stateId != "" {
		modelParam = fmt.Sprintf("&state=%s", stateId)
	}

	response, err := http.Get(fmt.Sprintf("https://api.mercadolibre.com/sites/%s/search?limit=50&until=today&category=%s%s%s%s", siteId, category, stateParam, brandParam, modelParam))
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
			Id:          itemResponse.Id,
			Title:       itemResponse.Title,
			Price:       itemResponse.Price,
			CurrencyId:  itemResponse.CurrencyId,
			Thumbnail:   itemResponse.Thumbnail,
			Permalink:   itemResponse.Permalink,
			StopTime:    itemResponse.StopTime,
			Description: buildItemDescription(itemResponse.Attributes),
		})
	}

	return &domain.SearchResult{
		Results:          items,
		Filters:          convertToDomainFilter(searchResponse.Filters),
		AvailableFilters: convertToDomainFilter(searchResponse.AvailableFilters),
	}
}

func convertToDomainFilter(searchFiltersResponse []*searchFilterResponse) []*domain.SearchFilter {
	var filters []*domain.SearchFilter

	for _, filterResponse := range searchFiltersResponse {
		var filterValues []*domain.SearchFilterValue

		for _, filterValueResponse := range filterResponse.Values {
			filterValues = append(filterValues, &domain.SearchFilterValue{
				Id:   filterValueResponse.Id,
				Name: filterValueResponse.Name,
			})
		}

		filters = append(filters, &domain.SearchFilter{
			Id:     filterResponse.Id,
			Name:   filterResponse.Name,
			Values: filterValues,
		})
	}

	return filters
}

func buildItemDescription(attributes []*searchItemAttributeResponse) string {
	kilometers := ""
	vehicleYear := ""

	for _, attribute := range attributes {
		if attribute.Id == "VEHICLE_YEAR" {
			vehicleYear = fmt.Sprintf("%s: %s", attribute.Name, attribute.ValueName)
		} else if attribute.Id == "KILOMETERS" {
			kilometers = fmt.Sprintf("%s: %s", attribute.Name, attribute.ValueName)
		}
	}

	return fmt.Sprintf("%s - %s", vehicleYear, kilometers)
}
