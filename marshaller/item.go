package marshaller

import (
	"fmt"

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
	Title string     `json:"title,omitempty"`
	Items []*ItemDto `json:"items,omitempty"`
}

type Item interface {
	GetView(item []*domain.Item) *ModelDto
}

type item struct {
}

func NewItem() Item {
	return &item{}
}

func (m item) GetView(itemsDomain []*domain.Item) *ModelDto {
	var itemsDto []*ItemDto

	for _, itemDomain := range itemsDomain {
		itemsDto = append(itemsDto, &ItemDto{
			Id:         itemDomain.Id,
			Title:      itemDomain.Title,
			Price:      fmt.Sprintf("%.2f", itemDomain.Price),
			CurrencyId: itemDomain.CurrencyId,
			Thumbnail:  itemDomain.Thumbnail,
			Permalink:  itemDomain.Permalink,
			StopTime:   itemDomain.StopTime,
		})
	}

	return &ModelDto{
		Title: "Publicaciones de autos usados que finalizan hoy, ideales para hacer una oferta!.",
		Items: itemsDto,
	}
}
