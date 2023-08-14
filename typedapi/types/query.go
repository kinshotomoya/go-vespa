package types

type Query struct {
	Fields     string
	Schema     string
	WhereQuery []WhereQuery
	Limit      int
	OffSet     int
}
