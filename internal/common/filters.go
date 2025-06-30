package common

import (
	"errors"
	"math"
	"ngMarketplace/pkg/validator"
	"strings"
)

var (
	ErrFilterValidationFailed = errors.New("filters validation failed")
)

// Filters holds information for pagination
type Filters struct {
	Page         int    `form:"page"`
	PageSize     int    `form:"page_size"`
	Sort         string `form:"sort"`
	SortSafeList []string
}

// Metadata holds information about current pagination
type Metadata struct {
	CurrentPage  int `json:"current_page,omitempty"`
	PageSize     int `json:"page_size,omitempty"`
	FirstPage    int `json:"first_page,omitempty"`
	LastPage     int `json:"last_page,omitempty"`
	TotalRecords int `json:"total_records,omitempty"`
}

// CalculateMetadata calculates metadata for response
func CalculateMetadata(totalRecords, page, pageSize int) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}

	return Metadata{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     int(math.Ceil(float64(totalRecords) / float64(pageSize))),
		TotalRecords: totalRecords,
	}
}

// ValidateFilters checks the fields of Filters
func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Page > 0, "page", "must be greater than zero")
	v.Check(f.Page < 10_000_000, "page", "must be a maximum of 10 million")
	v.Check(f.PageSize > 0, "page_size", "must be greater then zero")
	v.Check(f.PageSize <= 100, "page_size", "must be a maximum of 100")
	v.Check(validator.In(f.Sort, f.SortSafeList...), "sort", "invalid sort value")
}

// SortColumn checks that the client-provided Sort field matches one of the entries in our safeList
// and if it does, extract the column name from the Sort field by stripping the leading
// hyphen character (if one exists)
func (f Filters) SortColumn() string {
	for _, safeValue := range f.SortSafeList {
		if f.Sort == safeValue {
			return strings.TrimPrefix(f.Sort, "-")
		}
	}

	panic("unsafe sort parameter: " + f.Sort)
}

// SortDirection returns the sort direction ("ASC" or "DESC") depending on the prefix character of the
// Sort field
func (f Filters) SortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}

	return "ASC"
}

// Limit returns the PageSize
func (f Filters) Limit() int {
	return f.PageSize
}

// Offset calculates and returns the offset
func (f Filters) Offset() int {
	return (f.Page - 1) * f.PageSize
}
