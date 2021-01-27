package marshaller

import "github.com/barrerajuanjose/usefulsearch/domain"

type ItemDto struct {
	Id         string  `json:"id,omitempty"`
	Title      string  `json:"title,omitempty"`
	Price      float32 `json:"price,omitempty"`
	CurrencyId string  `json:"currency_id,omitempty"`
	Thumbnail  string  `json:"thumbnail,omitempty"`
	Permalink  string  `json:"permalink,omitempty"`
	StopTime   string  `json:"stop_time,omitempty"`
}

type Item interface {
	GetView(item []*domain.Item) []*ItemDto
}

type item struct {
}

func NewItem() Item {
	return &item{}
}

func (m item) GetView(itemsDomain []*domain.Item) []*ItemDto {
	var itemsDto []*ItemDto

	for _, itemDomain := range itemsDomain {
		itemsDto = append(itemsDto, &ItemDto{
			Id:         itemDomain.Id,
			Title:      itemDomain.Title,
			Price:      itemDomain.Price,
			CurrencyId: itemDomain.CurrencyId,
			Thumbnail:  itemDomain.Thumbnail,
			Permalink:  itemDomain.Permalink,
			StopTime:   itemDomain.StopTime,
		})
	}

	return itemsDto
}
