package hgapi

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// ----- API Definition -------------------------------------------------------

// CatSegments provides low-level information about the segments in the shards of an index.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cat-segments.html.
//
func newVertexGetIDFunc(t Transport) VertexGetID {
	return func(o ...func(req *VertexGetIDRequest)) (*VertexGetIDResponse, error) {
		var r = VertexGetIDRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

type VertexGetID func(o ...func(*VertexGetIDRequest)) (*VertexGetIDResponse, error)

type VertexGetIDRequest struct {
	Body io.Reader
	ctx  context.Context

	Label string
	ID    string
}

type VertexGetIDResponse struct {
	StatusCode int           `json:"-"`
	Header     http.Header   `json:"-"`
	Body       io.ReadCloser `json:"-"`
}

func (r VertexGetIDRequest) Do(ctx context.Context, transport Transport) (*VertexGetIDResponse, error) {

	req, _ := newRequest("GET", fmt.Sprintf(`/graphs/${GRAPH_NAME}/graph/vertices/"%s"`, r.ID), r.Body)

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	VertexGetIDResp := &VertexGetIDResponse{}

	VertexGetIDResp.StatusCode = res.StatusCode
	VertexGetIDResp.Header = res.Header
	VertexGetIDResp.Body = res.Body
	return VertexGetIDResp, nil
}

func (v VertexGetID) WithLabel(label string) func(*VertexGetIDRequest) {
	return func(r *VertexGetIDRequest) {
		r.Label = label
	}
}

func (v VertexGetID) WithID(id string) func(*VertexGetIDRequest) {
	return func(r *VertexGetIDRequest) {
		r.ID = id
	}
}

