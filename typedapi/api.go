package typedapi

import (
	"github.com/kinshotomoya/go-vespa/typedapi/core/search"
	"github.com/kinshotomoya/go-vespa/typedclient"
)

type Api struct {
	Search *search.Search
}

func New(t typedclient.ClientInterface) *Api {
	return &Api{
		Search: search.New(t),
		// TODO: GETなど必要なAPIを追加
	}
}
