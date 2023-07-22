package v1

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/izliang/go-hugegraph/hgapi"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// ----- API Definition -------------------------------------------------------

// 向HugeGraphServer发送gremlin语句（GET），同步执行
//
// See full documentation at https://hugegraph.apache.org/cn/docs/clients/restful-api/gremlin/#811-%E5%90%91hugegraphserver%E5%8F%91%E9%80%81gremlin%E8%AF%AD%E5%8F%A5get%E5%90%8C%E6%AD%A5%E6%89%A7%E8%A1%8C
//
func newGremlinGetFunc(t hgapi.Transport) GremlinGet {
	return func(o ...func(*GremlinGetRequest)) (*GremlinGetResponse, error) {
		var r = GremlinGetRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

type GremlinGet func(o ...func(*GremlinGetRequest)) (*GremlinGetResponse, error)

type GremlinGetRequest struct {
	ctx        context.Context
	Body       io.ReadCloser `json:"-"`
	GremlinGet *GremlinGetRequestReqData
}

type GremlinGetRequestReqData struct {
	Gremlin string `json:"gremlin"`
}

type GremlinGetResponse struct {
	StatusCode int           `json:"-"`
	Header     http.Header   `json:"-"`
	Body       io.ReadCloser `json:"-"`
}

func (r GremlinGetRequest) Do(ctx context.Context, transport hgapi.Transport) (*GremlinGetResponse, error) {

	if len(r.GremlinGet.Gremlin) < 1 {
		return nil, errors.New("GremlinGetRequest param error , gremlin is empty")
	}

	req, _ := hgapi.NewRequest("GET", "/gremlin", r.Body)

	params := url.Values{}
	if len(r.GremlinGet.Gremlin) > 0 {
		params.Set("gremlin", r.GremlinGet.Gremlin)
	}

	req.URL.RawQuery = params.Encode()
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	GremlinGetResp := &GremlinGetResponse{}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, GremlinGetResp)
	if err != nil {
		return nil, err
	}
	GremlinGetResp.StatusCode = res.StatusCode
	GremlinGetResp.Header = res.Header
	GremlinGetResp.Body = res.Body
	return GremlinGetResp, nil
}

func (g *GremlinGet) WithGremlinGetData(gremlin GremlinGetRequestReqData) func(*GremlinGetRequest) {
	return func(r *GremlinGetRequest) {
		r.GremlinGet = &gremlin
	}
}
