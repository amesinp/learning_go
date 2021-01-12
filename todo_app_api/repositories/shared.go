package repositories

import (
	"fmt"
)

// ResponseParams defines the parameters for a get by function
type ResponseParams struct {
	Skip          int
	Limit         int
	SortAscending bool
	IncludeCount  bool
	OrderBy       string
}

func (p *ResponseParams) getCountQuery() string {
	if p.IncludeCount {
		return ", COUNT(*) OVER() AS total_count"
	}
	return ""
}

func (p *ResponseParams) getQuerySuffix() string {
	if p.OrderBy == "" {
		p.OrderBy = "ID"
	}

	var sortDirection string
	if p.SortAscending {
		sortDirection = "asc"
	} else {
		sortDirection = "desc"
	}

	return fmt.Sprintf(" ORDER BY %s %s LIMIT %d OFFSET %d", p.OrderBy, sortDirection, p.Limit, p.Skip)
}
