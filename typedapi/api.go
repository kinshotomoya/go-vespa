package typedapi

import (
	"github.com/kinshotomoya/go-vespa"
	"github.com/kinshotomoya/go-vespa/typedapi/core/search"
)

type Api struct {
	Search *search.Search
}

func New(t *vespa.TypedClient) *Api {
	return &Api{
		Search: search.New(t),
		// TODO: GETなど必要なAPIを追加
	}
}
