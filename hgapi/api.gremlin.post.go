package hgapi

import (
	"context"
	"encoding/json"
	"errors"
	_ "fmt"
	"io"
	"io/ioutil"
	"net/http"
	_ "net/url"
	"strings"
)

// ----- API Definition -------------------------------------------------------

// 向HugeGraphServer发送gremlin语句（GET），同步执行
//
// See full documentation at https://hugegraph.apache.org/cn/docs/clients/restful-api/gremlin/#811-%E5%90%91hugegraphserver%E5%8F%91%E9%80%81gremlin%E8%AF%AD%E5%8F%A5get%E5%90%8C%E6%AD%A5%E6%89%A7%E8%A1%8C
//
func newGremlinPostFunc(t Transport) GremlinPost {
	return func(o ...func(*GremlinPostRequest)) (*GremlinPostResponse, error) {
		var r = GremlinPostRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

type GremlinPost func(o ...func(*GremlinPostRequest)) (*GremlinPostResponse, error)

type GremlinPostRequest struct {
	ctx         context.Context
	Body        io.ReadCloser `json:"-"`
	GremlinPost *GremlinPostRequestReqData
}

type GremlinPostRequestReqData struct {
	Gremlin  string            `json:"gremlin"`
	Bindings map[string]string `json:"bindings,omitempty"`
	Language string            `json:"language"`
	Aliases  struct {
		Graph string `json:"graph"`
		G     string `json:"g"`
	} `json:"aliases"`
}

type GremlinPostResponse struct {
	StatusCode int                      `json:"-"`
	Header     http.Header              `json:"-"`
	Body       io.ReadCloser            `json:"-"`
	Data       *GremlinPostResponseData `json:"data"`
}

type GremlinPostResponseData struct {
	RequestID string `json:"requestId,omitempty"`
	Status    struct {
		Message    string `json:"message"`
		Code       int    `json:"code"`
		Attributes struct {
		} `json:"attributes"`
	} `json:"status"`
	Result struct {
		Data interface{} `json:"data"`
		Meta struct {
		} `json:"meta"`
	} `json:"result,omitempty"`
	Exception string   `json:"exception,omitempty"`
	Message   string   `json:"message,omitempty"`
	Cause     string   `json:"cause,omitempty"`
	Trace     []string `json:"trace,omitempty"`
}

func (r GremlinPostRequest) Do(ctx context.Context, transport Transport) (*GremlinPostResponse, error) {

	if len(r.GremlinPost.Gremlin) < 1 {
		return nil, errors.New("GremlinPostRequest param error , gremlin is empty")
	}

	if len(r.GremlinPost.Language) < 1 {
		r.GremlinPost.Language = "gremlin-groovy"
	}

	// 重新修改参数
	r.GremlinPost.Aliases = struct {
		Graph string `json:"graph"`
		G     string `json:"g"`
	}(struct {
		Graph string
		G     string
	}{
		Graph: "${GRAPH_SPACE_NAME}-${GRAPH_NAME}",
		G:     "__g_${GRAPH_SPACE_NAME}-${GRAPH_NAME}",
	})

	byteBody, err := json.Marshal(&r.GremlinPost) // 序列化

	if err != nil {
		return nil, err
	}

	reader := strings.NewReader(string(byteBody)) // 转化为reader

	req, _ := newRequest("POST", "/gremlin", reader)

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	gremlinPostResp := &GremlinPostResponse{}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	respData := &GremlinPostResponseData{}
	err = json.Unmarshal(bytes, respData)
	if err != nil {
		return nil, err
	}
	gremlinPostResp.StatusCode = res.StatusCode
	gremlinPostResp.Header = res.Header
	gremlinPostResp.Body = res.Body
	gremlinPostResp.Data = respData
	return gremlinPostResp, nil
}

func (g *GremlinPost) WithGremlinPostData(gremlin GremlinPostRequestReqData) func(*GremlinPostRequest) {
	return func(r *GremlinPostRequest) {
		r.GremlinPost = &gremlin
	}
}
