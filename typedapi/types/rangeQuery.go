package types

import "fmt"

type RangeQuery struct {
	Field string
	Lower string // Infinity or -Infinity
	Upper string // Infinity or -Infinity
}

func NewRangeQuery(field string, lower *string, upper *string) *RangeQuery {
	var lowerBound string
	var upperBound string
	if lower == nil {
		lowerBound = "0"
	}
	if upper == nil {
		upperBound = "0"
	}
	return &RangeQuery{
		Field: field,
		Lower: lowerBound,
		Upper: upperBound,
	}
}

func (r *RangeQuery) BuildQuery() string {
	return fmt.Sprintf("range(%s, %s, %s)", r.Field, r.Lower, r.Upper)
}
