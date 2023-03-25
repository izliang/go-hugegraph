package hgapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// ----- API Definition -------------------------------------------------------

// CatSegments provides low-level information about the segments in the shards of an index.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cat-segments.html.
//
func newPropertyKeysCreateFunc(t Transport) PropertyKeysCreate {
	return func(o ...func(*PropertyKeysCreateRequest)) (*PropertyKeysCreateResponse, error) {
		var r = PropertyKeysCreateRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

type PropertyKeysCreate func(o ...func(*PropertyKeysCreateRequest)) (*PropertyKeysCreateResponse, error)

type PropertyKeysCreateRequest struct {
	ctx         context.Context `json:"-"`
	Graph       string          `json:"-"`
	Name        string          `json:"name"`
	DataType    string          `json:"data_type"`
	Cardinality string          `json:"cardinality"`
}

type PropertyKeysCreateResponse struct {
	StatusCode  int           `json:"-"`
	Header      http.Header   `json:"-"`
	Body        io.ReadCloser `json:"-"`
	PropertyKey struct {
		ID            int           `json:"id"`
		Name          string        `json:"name"`
		DataType      string        `json:"data_type"`
		Cardinality   string        `json:"cardinality"`
		AggregateType string        `json:"aggregate_type"`
		WriteType     string        `json:"write_type"`
		Properties    []interface{} `json:"properties"`
		Status        string        `json:"status"`
		UserData      struct {
			CreateTime string `json:"~create_time"`
		} `json:"user_data"`
	} `json:"property_key"`
	TaskID int `json:"task_id"`
}

func (r PropertyKeysCreateRequest) Do(ctx context.Context, transport Transport) (*PropertyKeysCreateResponse, error) {

	bytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	byteBody, _ := json.Marshal(&r)               // 序列化
	reader := strings.NewReader(string(byteBody)) // 转化为reader
	req, _ := newRequest("POST", fmt.Sprintf("/graphs/%s/schema/propertykeys", r.Graph), reader)

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	PropertyKeysCreateResp := &PropertyKeysCreateResponse{}
	bytes, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, PropertyKeysCreateResp)
	if err != nil {
		return nil, err
	}
	PropertyKeysCreateResp.StatusCode = res.StatusCode
	PropertyKeysCreateResp.Header = res.Header
	PropertyKeysCreateResp.Body = res.Body
	return PropertyKeysCreateResp, nil
}

func (v PropertyKeysCreate) WithGraph(graph string) func(*PropertyKeysCreateRequest) {
	return func(r *PropertyKeysCreateRequest) {
		r.Graph = graph
	}
}

func (v PropertyKeysCreate) WithName(name string) func(*PropertyKeysCreateRequest) {
	return func(r *PropertyKeysCreateRequest) {
		r.Name = name
	}
}

func (v PropertyKeysCreate) WithDataType(dataType string) func(*PropertyKeysCreateRequest) {
	//DOUBLE, BYTE, UNKNOWN, UUID, FLOAT, BLOB, DATE, OBJECT, BOOLEAN, TEXT, INT, LONG
	return func(r *PropertyKeysCreateRequest) {
		r.DataType = dataType
	}
}

func (v PropertyKeysCreate) WithCardinality(cardinality string) func(*PropertyKeysCreateRequest) {
	return func(r *PropertyKeysCreateRequest) {
		r.Cardinality = cardinality
	}
}
