package hgapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"hugegraph/internal/model"
	"io/ioutil"
	"net/http"
	"strings"
)

type PropertyKeyAction string

const (
	PropertyKeyActionAppend    PropertyKeyAction = "append"
	PropertyKeyActionEliminate PropertyKeyAction = "eliminate"
)

// ----- API Definition -------------------------------------------------------

// 为已存在的 PropertyKey 添加或移除 userdata
//
// See full documentation https://hugegraph.apache.org/cn/docs/clients/restful-api/propertykey/#122-%E4%B8%BA%E5%B7%B2%E5%AD%98%E5%9C%A8%E7%9A%84-propertykey-%E6%B7%BB%E5%8A%A0%E6%88%96%E7%A7%BB%E9%99%A4-userdata
//
func newPropertyKeysUpdateUserdataFunc(t Transport) PropertyKeysUpdateUserdata {
	return func(o ...func(*PropertyKeysUpdateUserdataRequest)) (*PropertyKeysUpdateUserdataResponse, error) {
		var r = PropertyKeysUpdateUserdataRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

type PropertyKeysUpdateUserdata func(o ...func(*PropertyKeysUpdateUserdataRequest)) (*PropertyKeysUpdateUserdataResponse, error)

type PropertyKeysUpdateUserdataRequest struct {
	ctx      context.Context             `json:"-"`
	Action   PropertyKeyAction           `json:"-"`
	Name     string                      `json:"name"`
	UserData *PropertyKeysUpdateUserData `json:"user_data"`
}

type PropertyKeysUpdateUserData struct {
	Min int `json:"min,omitempty"`
	Max int `json:"max,omitempty"`
}

type PropertyKeysUpdateUserdataResponse struct {
	StatusCode int         `json:"-"`
	Header     http.Header `json:"-"`
	Data       *PropertyKeysUpdateUserdataData
}

type PropertyKeysUpdateUserdataData struct {
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
			Min        int    `json:"min"`
			Max        int    `json:"max"`
			CreateTime string `json:"~create_time"`
		} `json:"user_data"`
	} `json:"property_key"`
	TaskID int `json:"task_id"`
}

func (r PropertyKeysUpdateUserdataRequest) Do(ctx context.Context, transport Transport) (*PropertyKeysUpdateUserdataResponse, error) {

	if len(r.Name) < 1 {
		return nil, errors.New("PropertyKeysUpdateUserdataRequest Param error, name is empty")
	}

	if len(r.Action) < 1 {
		return nil, errors.New("PropertyKeysUpdateUserdataRequest Param error, action is empty")
	}

	bytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	byteBody, _ := json.Marshal(&r)               // 序列化
	reader := strings.NewReader(string(byteBody)) // 转化为reader

	req, _ := newRequest("PUT", fmt.Sprintf(model.UrlPrefix+"/graphs/${GRAPH_NAME}/schema/propertykeys/%s?action=%s", r.Name, r.Action), reader)

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	PropertyKeysUpdateUserdataResp := &PropertyKeysUpdateUserdataResponse{}
	bytes, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	respData := &PropertyKeysUpdateUserdataData{}
	err = json.Unmarshal(bytes, respData)
	if err != nil {
		return nil, err
	}
	PropertyKeysUpdateUserdataResp.StatusCode = res.StatusCode
	PropertyKeysUpdateUserdataResp.Header = res.Header
	PropertyKeysUpdateUserdataResp.Data = respData
	return PropertyKeysUpdateUserdataResp, nil
}

func (v PropertyKeysUpdateUserdata) WithName(name string) func(*PropertyKeysUpdateUserdataRequest) {
	return func(r *PropertyKeysUpdateUserdataRequest) {
		r.Name = name
	}
}

func (v PropertyKeysUpdateUserdata) WithAction(action PropertyKeyAction) func(*PropertyKeysUpdateUserdataRequest) {
	return func(r *PropertyKeysUpdateUserdataRequest) {
		r.Action = action
	}
}

func (v PropertyKeysUpdateUserdata) WithUserdata(userdata PropertyKeysUpdateUserData) func(*PropertyKeysUpdateUserdataRequest) {
	return func(r *PropertyKeysUpdateUserdataRequest) {
		r.UserData = &userdata
	}
}
