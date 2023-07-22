package hgapi

import (
	"context"
	"encoding/json"
	"hugegraph/internal/model"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// ----- API Definition -------------------------------------------------------

// Schema
//
// See full documentation at https://hugegraph.apache.org/cn/docs/clients/restful-api/schema/#11-schema
//
func newSchemaGetFunc(t Transport) SchemaGet {
	return func(o ...func(*SchemaGetRequest)) (*SchemaGetResponse, error) {
		var r = SchemaGetRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

type SchemaGet func(o ...func(*SchemaGetRequest)) (*SchemaGetResponse, error)

type SchemaGetRequest struct {
	ctx    context.Context
	Format string
}

type SchemaGetResponse struct {
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
	Vertexlabels []struct {
		ID               int      `json:"id"`
		Name             string   `json:"name"`
		IDStrategy       string   `json:"id_strategy"`
		PrimaryKeys      []string `json:"primary_keys"`
		NullableKeys     []string `json:"nullable_keys"`
		IndexLabels      []string `json:"index_labels"`
		Properties       []string `json:"properties"`
		Status           string   `json:"status"`
		TTL              int      `json:"ttl"`
		EnableLabelIndex bool     `json:"enable_label_index"`
		UserData         struct {
			CreateTime string `json:"~create_time"`
		} `json:"user_data"`
	} `json:"vertexlabels"`
	Edgelabels []struct {
		ID               int      `json:"id"`
		Name             string   `json:"name"`
		SourceLabel      string   `json:"source_label"`
		TargetLabel      string   `json:"target_label"`
		Frequency        string   `json:"frequency"`
		SortKeys         []string `json:"sort_keys"`
		NullableKeys     []string `json:"nullable_keys"`
		IndexLabels      []string `json:"index_labels"`
		Properties       []string `json:"properties"`
		Status           string   `json:"status"`
		TTL              int      `json:"ttl"`
		EnableLabelIndex bool     `json:"enable_label_index"`
		UserData         struct {
			CreateTime string `json:"~create_time"`
		} `json:"user_data"`
	} `json:"edgelabels"`
	Indexlabels []struct {
		ID        int      `json:"id"`
		Name      string   `json:"name"`
		BaseType  string   `json:"base_type"`
		BaseValue string   `json:"base_value"`
		IndexType string   `json:"index_type"`
		Fields    []string `json:"fields"`
		Status    string   `json:"status"`
		UserData  struct {
			CreateTime string `json:"~create_time"`
		} `json:"user_data"`
	} `json:"indexlabels"`
}

func (r SchemaGetRequest) Do(ctx context.Context, transport Transport) (*SchemaGetResponse, error) {

	req, _ := newRequest("GET", model.UrlPrefix+"/graphs/${GRAPH_NAME}/schema", nil)

	if len(r.Format) > 0 {
		params := url.Values{}
		params.Set("format", r.Format)
		req.URL.RawQuery = params.Encode()
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	SchemaGetResp := &SchemaGetResponse{}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, SchemaGetResp)
	if err != nil {
		return nil, err
	}
	SchemaGetResp.StatusCode = res.StatusCode
	SchemaGetResp.Header = res.Header
	SchemaGetResp.Body = res.Body
	return SchemaGetResp, nil
}

func (v SchemaGet) WithFormat(format string) func(*SchemaGetRequest) {
	return func(r *SchemaGetRequest) {
		r.Format = format
	}
}
