package types

import "strings"

type AndQuery struct {
	query []WhereQuery
}

func (a *AndQuery) BuildQuery() string {
	queries := make([]string, len(a.query))
	for i := range a.query {
		queries[i] = a.query[i].BuildQuery()
	}
	return strings.Join(queries, " AND ")
}
