package marshaller

import (
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/barrerajuanjose/usefulsearch/domain"
)

type ItemDto struct {
	Id         string `json:"id,omitempty"`
	Title      string `json:"title,omitempty"`
	Price      string `json:"price,omitempty"`
	CurrencyId string `json:"currency_id,omitempty"`
	Thumbnail  string `json:"thumbnail,omitempty"`
	Permalink  string `json:"permalink,omitempty"`
	StopTime   string `json:"stop_time,omitempty"`
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
			Id:         itemDomain.Id,
			Title:      itemDomain.Title,
			Price:      p.Sprintf("%d", int(itemDomain.Price)),
			CurrencyId: itemDomain.CurrencyId,
			Thumbnail:  strings.Replace(itemDomain.Thumbnail, "-I", "-O", 1),
			Permalink:  itemDomain.Permalink,
			StopTime:   itemDomain.StopTime,
		})
	}

	var brandFilterValues []*FilterValueDto

	for _, filtersDomaind := range searchResult.Filters {
		if filtersDomaind.Id == "BRAND" {
			for _, filtersValuesDomaind := range filtersDomaind.Values {
				brandFilterValues = append(brandFilterValues, &FilterValueDto{
					Name:     filtersValuesDomaind.Name,
					Value:    filtersValuesDomaind.Id,
					Selected: true,
				})
			}
		}
	}

	if len(brandFilterValues) == 0 {
		brandFilterValues = append(brandFilterValues, &FilterValueDto{
			Name:     "Selecciona una marca",
			Value:    "",
			Selected: true,
		})
	} else {
		brandFilterValues = append(brandFilterValues, &FilterValueDto{
			Name:     "Limpar selección",
			Value:    "",
			Selected: false,
		})
	}

	for _, filtersDomaind := range searchResult.AvailableFilters {
		if filtersDomaind.Id == "BRAND" {
			for _, filtersValuesDomaind := range filtersDomaind.Values {
				brandFilterValues = append(brandFilterValues, &FilterValueDto{
					Name:     filtersValuesDomaind.Name,
					Value:    filtersValuesDomaind.Id,
					Selected: false,
				})
			}
		}
	}

	return &ModelDto{
		PageTitle:         "Autos Usados Mercado Libre Útima Oportunidad",
		PageDescription:   "Publicaciones de autos usados que finalizan hoy, ideales para hacer una oferta!.",
		Title:             "Autos usados en Mercado Libre! Ultima oportunidad para comprarlos",
		Items:             itemsDto,
		BrandFilterValues: brandFilterValues,
	}
}
