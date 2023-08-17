package types

import (
	"strconv"
	"strings"
)

type Query struct {
	Fields     []string
	Schema     string
	WhereQuery []WhereQuery
	Limit      *int
	OffSet     *int
}

func (q *Query) Build() string {
	queries := make([]string, len(q.WhereQuery))
	for i := range q.WhereQuery {
		queries[i] = q.WhereQuery[i].BuildQuery()
	}

	// NOTE: Need a slice that can store a minimum of 6 and a maximum of 10 elements, including limit and offset
	yqlSlice := make([]string, 6, 10)
	yqlSlice = append(yqlSlice, "select", strings.Join(q.Fields, ","), "from", q.Schema, "where", strings.Join(queries, " and "))
	if q.Limit != nil {
		yqlSlice = append(yqlSlice, "limit", strconv.Itoa(*q.Limit))
	}

	if q.OffSet != nil {
		yqlSlice = append(yqlSlice, "offset", strconv.Itoa(*q.OffSet))
	}

	return strings.Join(yqlSlice[:], " ")
}
