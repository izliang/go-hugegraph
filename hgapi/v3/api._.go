package v3

import "hugegraph/hgapi"

// Code generated from specification version 5.6.15 (fe7575a32e2): DO NOT EDIT

// API contains the Elasticsearch APIs
//
type APIV3 struct {
	Version Version
	Gremlin *Gremlin
}

type Gremlin struct {
	Post GremlinPost
}

// New creates new API
//
func New(t hgapi.Transport) *APIV3 {
	return &APIV3{
		Version: newVersionFunc(t),
		Gremlin: &Gremlin{
			Post: newGremlinPostFunc(t),
		},
	}
}
