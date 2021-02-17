package marshaller

import (
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/barrerajuanjose/usefulsearch/config"
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
	SiteConfiguration *SiteConfigurationDto `json:"site_configuration,omitempty"`
	Items             []*ItemDto            `json:"items,omitempty"`
	BrandFilterValues []*FilterValueDto     `json:"brand_filter_values,omitempty"`
	StateFilterValues []*FilterValueDto     `json:"state_filter_values,omitempty"`
	AvailableSites    []*SiteDto            `json:"available_sites,omitempty"`
}

type SiteConfigurationDto struct {
	Canonical       string `json:"canonical,omitempty"`
	PageTitle       string `json:"page_title,omitempty"`
	PageDescription string `json:"page_description,omitempty"`
	Title           string `json:"title,omitempty"`
}

type SiteDto struct {
	Canonical string `json:"canonical,omitempty"`
	Name      string `json:"name,omitempty"`
}

type Page interface {
	GetIndex(siteConfig *config.SiteConfiguration, avaiblableSites []*config.SiteConfiguration) *ModelDto
	GetUsefulPage(siteConfig *config.SiteConfiguration, avaiblableSites []*config.SiteConfiguration, searchResult *domain.SearchResult) *ModelDto
}

type page struct {
}

func NewPage() Page {
	return &page{}
}
func (m page) GetIndex(siteConfig *config.SiteConfiguration, avaiblableConfigurationSites []*config.SiteConfiguration) *ModelDto {
	return &ModelDto{
		SiteConfiguration: buildSiteConfiguration(siteConfig),
		AvailableSites:    buildAllAvailableSites(avaiblableConfigurationSites),
	}
}

func (m page) GetUsefulPage(siteConfig *config.SiteConfiguration, avaiblableConfigurationSites []*config.SiteConfiguration, searchResult *domain.SearchResult) *ModelDto {
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
		SiteConfiguration: buildSiteConfiguration(siteConfig),
		Items:             itemsDto,
		BrandFilterValues: brandFilterValues,
		StateFilterValues: stateFilterValues,
		AvailableSites:    buildAvailableSites(siteConfig, avaiblableConfigurationSites),
	}
}

func buildSiteConfiguration(siteConfig *config.SiteConfiguration) *SiteConfigurationDto {
	return &SiteConfigurationDto{
		PageTitle:       siteConfig.PageTitle,
		PageDescription: siteConfig.PageDescription,
		Title:           siteConfig.Title,
		Canonical:       siteConfig.BaseUrl + siteConfig.URI,
	}
}

func buildAvailableSites(currentSiteConfig *config.SiteConfiguration, avaiblableConfigurationSites []*config.SiteConfiguration) []*SiteDto {
	var sitesDto []*SiteDto

	for _, availableConfigurationSite := range avaiblableConfigurationSites {
		if currentSiteConfig.SiteId != availableConfigurationSite.SiteId && !availableConfigurationSite.IsDefault {
			sitesDto = append(sitesDto, &SiteDto{
				Canonical: availableConfigurationSite.BaseUrl + availableConfigurationSite.URI,
				Name:      availableConfigurationSite.Name,
			})
		}
	}

	return sitesDto
}

func buildAllAvailableSites(avaiblableConfigurationSites []*config.SiteConfiguration) []*SiteDto {
	var sitesDto []*SiteDto

	for _, availableConfigurationSite := range avaiblableConfigurationSites {
		sitesDto = append(sitesDto, &SiteDto{
			Canonical: availableConfigurationSite.BaseUrl + availableConfigurationSite.URI,
			Name:      availableConfigurationSite.Name,
		})
	}

	return sitesDto
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
