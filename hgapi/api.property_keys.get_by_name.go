package hgapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"hugegraph/internal/model"
	"io"
	"io/ioutil"
	"net/http"
)

// ----- API Definition -------------------------------------------------------

// 根据name获取PropertyKey
//
// See full documentation https://hugegraph.apache.org/cn/docs/clients/restful-api/propertykey/#124-%E6%A0%B9%E6%8D%AEname%E8%8E%B7%E5%8F%96propertykey
//
func newPropertyKeysGetByNameFunc(t Transport) PropertyKeysGetByName {
	return func(o ...func(*PropertyKeysGetByNameRequest)) (*PropertyKeysGetByNameResponse, error) {
		var r = PropertyKeysGetByNameRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

type PropertyKeysGetByName func(o ...func(*PropertyKeysGetByNameRequest)) (*PropertyKeysGetByNameResponse, error)

type PropertyKeysGetByNameRequest struct {
	Body io.Reader
	ctx  context.Context
	Name string
}

type PropertyKeysGetByNameResponse struct {
	StatusCode int         `json:"-"`
	Header     http.Header `json:"-"`
	Data       *PropertyKeysGetByNameData
}

type PropertyKeysGetByNameData struct {
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
}

func (r PropertyKeysGetByNameRequest) Do(ctx context.Context, transport Transport) (*PropertyKeysGetByNameResponse, error) {

	if len(r.Name) < 1 {
		return nil, errors.New("PropertyKeysGetByNameRequest Param error, name is not empty")
	}

	req, _ := newRequest("GET", fmt.Sprintf(model.UrlPrefix+"/graphs/${GRAPH_NAME}/schema/propertykeys/%s", r.Name), r.Body)

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	propertyKeysGetByNameResp := &PropertyKeysGetByNameResponse{}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	respData := &PropertyKeysGetByNameData{}
	err = json.Unmarshal(bytes, respData)
	if err != nil {
		return nil, err
	}
	propertyKeysGetByNameResp.StatusCode = res.StatusCode
	propertyKeysGetByNameResp.Header = res.Header
	propertyKeysGetByNameResp.Data = respData
	return propertyKeysGetByNameResp, nil
}

func (v PropertyKeysGetByName) WithName(name string) func(*PropertyKeysGetByNameRequest) {
	return func(r *PropertyKeysGetByNameRequest) {
		r.Name = name
	}
}
