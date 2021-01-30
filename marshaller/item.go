package marshaller

import (
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/barrerajuanjose/usefulsearch/domain"
)

type ItemDto struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Price       string `json:"price,omitempty"`
	CurrencyId  string `json:"currency_id,omitempty"`
	Thumbnail   string `json:"thumbnail,omitempty"`
	Permalink   string `json:"permalink,omitempty"`
	StopTime    string `json:"stop_time,omitempty"`
	Description string `json:"description,omitempty"`
}

type FilterValueDto struct {
	Name     string `json:"name,omitempty"`
	Value    string `json:"value,omitempty"`
	Selected bool   `json:"selected,omitempty"`
}

type ModelDto struct {
	PageTitle         string            `json:"page_title,omitempty"`
	PageDescription   string            `json:"page_description,omitempty"`
	Title             string            `json:"title,omitempty"`
	Items             []*ItemDto        `json:"items,omitempty"`
	BrandFilterValues []*FilterValueDto `json:"brand_filter_values,omitempty"`
	StateFilterValues []*FilterValueDto `json:"state_filter_values,omitempty"`
}

type Item interface {
	GetView(searchResult *domain.SearchResult) *ModelDto
}

type item struct {
}

func NewItem() Item {
	return &item{}
}

func (m item) GetView(searchResult *domain.SearchResult) *ModelDto {
	var itemsDto []*ItemDto
	p := message.NewPrinter(language.BrazilianPortuguese)

	for _, itemDomain := range searchResult.Results {
		itemsDto = append(itemsDto, &ItemDto{
			Id:          itemDomain.Id,
			Title:       itemDomain.Title,
			Price:       p.Sprintf("%d", int(itemDomain.Price)),
			CurrencyId:  itemDomain.CurrencyId,
			Thumbnail:   strings.Replace(itemDomain.Thumbnail, "-I", "-C", 1),
			Permalink:   itemDomain.Permalink,
			StopTime:    itemDomain.StopTime,
			Description: itemDomain.Description,
		})
	}

	brandFilterValues := buildFilter("BRAND", searchResult.Filters, searchResult.AvailableFilters)
	stateFilterValues := buildFilter("state", searchResult.Filters, searchResult.AvailableFilters)

	return &ModelDto{
		PageTitle:         "Autos Usados Mercado Libre Útima Oportunidad",
		PageDescription:   "Publicaciones de autos usados que finalizan hoy, ideales para hacer una oferta!.",
		Title:             "Autos usados en Mercado Libre! Ultima oportunidad para comprarlos",
		Items:             itemsDto,
		BrandFilterValues: brandFilterValues,
		StateFilterValues: stateFilterValues,
	}
}

func buildFilter(filterId string, filters []*domain.SearchFilter, availavleFilters []*domain.SearchFilter) []*FilterValueDto {
	var filterValues []*FilterValueDto

	for _, filtersDomaind := range filters {
		if filtersDomaind.Id == filterId {
			for _, filtersValuesDomaind := range filtersDomaind.Values {
				filterValues = append(filterValues, &FilterValueDto{
					Name:     filtersValuesDomaind.Name,
					Value:    filtersValuesDomaind.Id,
					Selected: true,
				})
			}
		}
	}

	if len(filterValues) == 0 {
		filterValues = append(filterValues, &FilterValueDto{
			Name:     "Selecciona una opción",
			Value:    "",
			Selected: true,
		})
	} else {
		filterValues = append(filterValues, &FilterValueDto{
			Name:     "Limpar selección",
			Value:    "clean",
			Selected: false,
		})
	}

	for _, filtersDomaind := range availavleFilters {
		if filtersDomaind.Id == filterId {
			for _, filtersValuesDomaind := range filtersDomaind.Values {
				filterValues = append(filterValues, &FilterValueDto{
					Name:     filtersValuesDomaind.Name,
					Value:    filtersValuesDomaind.Id,
					Selected: false,
				})
			}
		}
	}

	return filterValues
}
