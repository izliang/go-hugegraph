package v1

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/izliang/go-hugegraph/hgapi"
	"github.com/izliang/go-hugegraph/internal/model"
	"io/ioutil"
	"net/http"
	"strings"
)

// ----- API Definition -------------------------------------------------------

// 创建一个VertexLabel
//
// See full documentation at https://hugegraph.apache.org/cn/docs/clients/restful-api/vertexlabel/#131-%E5%88%9B%E5%BB%BA%E4%B8%80%E4%B8%AAvertexlabel
//
func newVertexLabelCreateFunc(t hgapi.Transport) VertexLabelCreate {
	return func(o ...func(*VertexLabelCreateRequest)) (*VertexLabelCreateResponse, error) {
		var r = VertexLabelCreateRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

type VertexLabelCreate func(o ...func(*VertexLabelCreateRequest)) (*VertexLabelCreateResponse, error)

type VertexLabelCreateRequest struct {
	ctx  context.Context
	Data *VertexLabelCreateRequestData
}

type VertexLabelCreateRequestData struct {
	Name             string                    `json:"name"`
	IDStrategy       VertexLabelIDStrategyType `json:"id_strategy"`
	Properties       []string                  `json:"properties"`
	PrimaryKeys      []string                  `json:"primary_keys"`
	NullableKeys     []string                  `json:"nullable_keys"`
	CheckExist       bool                      `json:"check_exist"`
	Ttl              int                       `json:"ttl,omitempty"`
	TtlStartTime     string                    `json:"ttl_start_time,omitempty"`
	EnableLabelIndex bool                      `json:"enable_label_index"`
	UserData         struct {
		Super string `json:"super"`
	} `json:"user_data"`
}

type VertexLabelCreateResponse struct {
	StatusCode int         `json:"-"`
	Header     http.Header `json:"-"`
	Data       *VertexLabelCreateResponseData
}

type VertexLabelCreateResponseData struct {
	ID               int      `json:"id"`
	PrimaryKeys      []string `json:"primary_keys"`
	IDStrategy       string   `json:"id_strategy"`
	Name             string   `json:"name"`
	IndexNames       []string `json:"index_names"`
	Properties       []string `json:"properties"`
	NullableKeys     []string `json:"nullable_keys"`
	EnableLabelIndex bool     `json:"enable_label_index"`
	UserData         struct {
		Super string `json:"super,omitempty"`
	} `json:"user_data,omitempty"`
}

func (r VertexLabelCreateRequest) Do(ctx context.Context, transport hgapi.Transport) (*VertexLabelCreateResponse, error) {

	if r.Data == nil {
		return nil, errors.New("VertexLabelCreateRequest Param error, data is nil")
	}

	if len(r.Data.Name) < 1 {
		return nil, errors.New("VertexLabelCreateRequest Param error, vertex.name is empty")
	}

	if r.Data.Properties == nil || len(r.Data.Properties) < 1 {
		return nil, errors.New("VertexLabelCreateRequest Param error, vertex.properties is empty")
	}

	bytes, err := json.Marshal(r.Data) // 序列化
	if err != nil {
		return nil, err
	}
	reader := strings.NewReader(string(bytes)) // 转化为reader

	req, _ := hgapi.NewRequest("POST", model.UrlPrefix+"/graphs/${GRAPH_NAME}/schema/vertexlabels", reader)

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	vertexLabelCreateResponseData := &VertexLabelCreateResponseData{}
	bytes, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, vertexLabelCreateResponseData)
	if err != nil {
		return nil, err
	}

	vertexLabelCreateResponse := &VertexLabelCreateResponse{}

	vertexLabelCreateResponse.StatusCode = res.StatusCode
	vertexLabelCreateResponse.Header = res.Header
	vertexLabelCreateResponse.Data = vertexLabelCreateResponseData

	return vertexLabelCreateResponse, nil
}

func (v *VertexLabelCreate) WithData(data VertexLabelCreateRequestData) func(*VertexLabelCreateRequest) {
	return func(r *VertexLabelCreateRequest) {
		r.Data = &data
	}
}
