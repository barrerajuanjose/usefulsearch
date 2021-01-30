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

type ModelDto struct {
	PageTitle       string     `json:"page_title,omitempty"`
	PageDescription string     `json:"page_description,omitempty"`
	Title           string     `json:"title,omitempty"`
	Items           []*ItemDto `json:"items,omitempty"`
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

	return &ModelDto{
		PageTitle:       "Autos Usados Mercado Libre Útima Oportunidad",
		PageDescription: "Publicaciones de autos usados que finalizan hoy, ideales para hacer una oferta!.",
		Title:           "Autos usados en Mercado Libre! Ultima oportunidad para comprarlos",
		Items:           itemsDto,
	}
}
