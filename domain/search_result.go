package domain

type SearchResult struct {
	Results          []*Item
	Filters          []*SearchFilter
	AvailableFilters []*SearchFilter
}

type SearchFilterValue struct {
	Id   string
	Name string
}

type SearchFilter struct {
	Id     string
	Name   string
	Values []*SearchFilterValue
}
