package search

import "github.com/kinshotomoya/go-vespa"

type Search struct {
	transport *vespa.TypedClient
	req       *Request
}

func New(t *vespa.TypedClient) *Search {
	return &Search{
		transport: t,
	}
}

func (s *Search) Request(r *Request) *Search {
	s.req = r
	return s
}

// TODO: httpリクエスト実行
func (r *Search) Do() {
	// TODO: performメソッドをTypedClientに実装してそれを実行する
	//r.transport
}
