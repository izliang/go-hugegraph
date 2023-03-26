package hgapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// ----- API Definition -------------------------------------------------------

// 根据name删除PropertyKey
//
// See full documentation https://hugegraph.apache.org/cn/docs/clients/restful-api/propertykey/#125-%E6%A0%B9%E6%8D%AEname%E5%88%A0%E9%99%A4propertykey
//
func newPropertyKeysDeleteByNameFunc(t Transport) PropertyKeysDeleteByName {
	return func(o ...func(*PropertyKeysDeleteByNameRequest)) (*PropertyKeysDeleteByNameResponse, error) {
		var r = PropertyKeysDeleteByNameRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

type PropertyKeysDeleteByName func(o ...func(*PropertyKeysDeleteByNameRequest)) (*PropertyKeysDeleteByNameResponse, error)

type PropertyKeysDeleteByNameRequest struct {
	Body io.Reader
	ctx  context.Context
	Name string
}

type PropertyKeysDeleteByNameResponse struct {
	StatusCode int           `json:"-"`
	Header     http.Header   `json:"-"`
	Body       io.ReadCloser `json:"-"`
}

func (r PropertyKeysDeleteByNameRequest) Do(ctx context.Context, transport Transport) (*PropertyKeysDeleteByNameResponse, error) {

	if len(r.Name) < 1 {
		return nil, errors.New("PropertyKeysDeleteByNameRequest Param error, name is not empty")
	}

	req, _ := newRequest("DELETE", fmt.Sprintf("/graphs/${GRAPH_NAME}/schema/propertykeys/%s", r.Name), r.Body)

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	PropertyKeysDeleteByNameResp := &PropertyKeysDeleteByNameResponse{}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, PropertyKeysDeleteByNameResp)
	if err != nil {
		return nil, err
	}
	PropertyKeysDeleteByNameResp.StatusCode = res.StatusCode
	PropertyKeysDeleteByNameResp.Header = res.Header
	PropertyKeysDeleteByNameResp.Body = res.Body
	return PropertyKeysDeleteByNameResp, nil
}

func (v PropertyKeysDeleteByName) WithName(name string) func(*PropertyKeysDeleteByNameRequest) {
	return func(r *PropertyKeysDeleteByNameRequest) {
		r.Name = name
	}
}
