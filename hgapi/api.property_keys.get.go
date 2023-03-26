package hgapi

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// ----- API Definition -------------------------------------------------------

// CatSegments provides low-level information about the segments in the shards of an index.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cat-segments.html.
//
func newPropertyKeysGetFunc(t Transport) PropertyKeysGet {
	return func(o ...func(*PropertyKeysGetRequest)) (*PropertyKeysGetResponse, error) {
		var r = PropertyKeysGetRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

type PropertyKeysGet func(o ...func(*PropertyKeysGetRequest)) (*PropertyKeysGetResponse, error)

type PropertyKeysGetRequest struct {
	Body io.Reader
	ctx  context.Context
}

type PropertyKeysGetResponse struct {
	StatusCode   int           `json:"-"`
	Header       http.Header   `json:"-"`
	Body         io.ReadCloser `json:"-"`
	Propertykeys []struct {
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
	} `json:"propertykeys"`
}

func (r PropertyKeysGetRequest) Do(ctx context.Context, transport Transport) (*PropertyKeysGetResponse, error) {

	req, _ := newRequest("GET", "/graphs/${GRAPH_NAME}/schema/propertykeys", r.Body)

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	PropertyKeysGetResp := &PropertyKeysGetResponse{}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, PropertyKeysGetResp)
	if err != nil {
		return nil, err
	}
	PropertyKeysGetResp.StatusCode = res.StatusCode
	PropertyKeysGetResp.Header = res.Header
	PropertyKeysGetResp.Body = res.Body
	return PropertyKeysGetResp, nil
}
