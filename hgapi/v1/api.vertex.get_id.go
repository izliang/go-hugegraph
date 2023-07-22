package v1

import (
	"context"
	"fmt"
	"github.com/izliang/go-hugegraph/hgapi"
	"github.com/izliang/go-hugegraph/internal/model"
	"io"
	"net/http"
)

// ----- API Definition -------------------------------------------------------

// 根据Id获取顶点
//
// See full documentation at https://hugegraph.apache.org/cn/docs/clients/restful-api/vertex/#217-%E6%A0%B9%E6%8D%AEid%E8%8E%B7%E5%8F%96%E9%A1%B6%E7%82%B9
//
func newVertexGetIDFunc(t hgapi.Transport) VertexGetID {
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

func (r VertexGetIDRequest) Do(ctx context.Context, transport hgapi.Transport) (*VertexGetIDResponse, error) {

	req, _ := hgapi.NewRequest("GET", fmt.Sprintf(model.UrlPrefix+`/graphs/${GRAPH_NAME}/graph/vertices/"%s"`, r.ID), r.Body)

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
