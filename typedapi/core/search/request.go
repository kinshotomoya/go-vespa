package search

import (
	"github.com/kinshotomoya/go-vespa/typedapi/types"
	"time"
)

type Request struct {
	Path    string
	Query   types.Query
	TimeOut time.Duration
}
