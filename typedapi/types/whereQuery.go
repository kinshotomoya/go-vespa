package types

type WhereQuery interface {
	BuildQuery() string
}
