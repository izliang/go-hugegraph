package v1

import (
	"context"
	"encoding/json"
	"hugegraph/hgapi"
	"hugegraph/internal/model"
	"io"
	"io/ioutil"
	"net/http"
)

// ----- API Definition -------------------------------------------------------

// 查看HugeGraph的版本信息
//
// See full documentation at https://hugegraph.apache.org/cn/docs/clients/restful-api/other/#1011-%E6%9F%A5%E7%9C%8Bhugegraph%E7%9A%84%E7%89%88%E6%9C%AC%E4%BF%A1%E6%81%AF
//
func newVertexCreateFunc(t hgapi.Transport) VertexCreate {
	return func(o ...func(*VertexCreateRequest)) (*VertexCreateResponse, error) {
		var r = VertexCreateRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

type VertexCreate func(o ...func(*VertexCreateRequest)) (*VertexCreateResponse, error)

type VertexCreateRequest struct {
	Body io.Reader
	ctx  context.Context
}

type VertexCreateResponse struct {
	StatusCode    int           `json:"-"`
	Header        http.Header   `json:"-"`
	Body          io.ReadCloser `json:"-"`
	VertexCreates struct {
		VertexCreate string `json:"VertexCreate"`
		Core         string `json:"core"`
		Gremlin      string `json:"gremlin"`
		API          string `json:"api"`
	} `json:"VertexCreates"`
}

func (r VertexCreateRequest) Do(ctx context.Context, transport hgapi.Transport) (*VertexCreateResponse, error) {

	req, _ := hgapi.NewRequest("GET", model.UrlPrefix+"/VertexCreates", r.Body)

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	VertexCreateResp := &VertexCreateResponse{}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, VertexCreateResp)
	if err != nil {
		return nil, err
	}
	VertexCreateResp.StatusCode = res.StatusCode
	VertexCreateResp.Header = res.Header
	VertexCreateResp.Body = res.Body
	return VertexCreateResp, nil
}
