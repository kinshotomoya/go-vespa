package types

import (
	"strings"
)

type Query struct {
	Fields     []string
	Schema     string
	WhereQuery []WhereQuery
	Limit      int
	OffSet     int
}

func (q *Query) Build() string {
	fields := strings.Join(q.Fields, ",")
	queries := make([]string, len(q.WhereQuery))
	for i := range q.WhereQuery {
		queries[i] = q.WhereQuery[i].BuildQuery()
	}
	yqlSlice := []string{"SELECT", fields, "FROM", q.Schema, "WHERE", strings.Join(queries, " AND ")}
	return strings.Join(yqlSlice, " ")
}
