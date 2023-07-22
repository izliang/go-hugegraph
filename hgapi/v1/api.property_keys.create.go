package v1

import (
	"context"
	"encoding/json"
	"hugegraph/hgapi"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type PropertyKeyDataType string
type PropertyCardinalityType string

const (
	PropertyDataTypeDouble  PropertyKeyDataType = "DOUBLE"
	PropertyDataTypeByte    PropertyKeyDataType = "BYTE"
	PropertyDataTypeUnknown PropertyKeyDataType = "UNKNOWN"
	PropertyDataTypeUuid    PropertyKeyDataType = "UUID"
	PropertyDataTypeFloat   PropertyKeyDataType = "FLOAT"
	PropertyDataTypeBlob    PropertyKeyDataType = "BLOB"
	PropertyDataTypeDate    PropertyKeyDataType = "DATE"
	PropertyDataTypeObject  PropertyKeyDataType = "OBJECT"
	PropertyDataTypeBoolean PropertyKeyDataType = "BOOLEAN"
	PropertyDataTypeText    PropertyKeyDataType = "TEXT"
	PropertyDataTypeInt     PropertyKeyDataType = "INT"
	PropertyDataTypeLong    PropertyKeyDataType = "LONG"

	PropertyCardinalityTypeSingle PropertyCardinalityType = "SINGLE"
	PropertyCardinalityTypeSet    PropertyCardinalityType = "SET"
	PropertyCardinalityTypeList   PropertyCardinalityType = "LIST"
)

// ----- API Definition -------------------------------------------------------

// 创建一个 PropertyKey
//
// See full documentation https://hugegraph.apache.org/cn/docs/clients/restful-api/propertykey/#121-%E5%88%9B%E5%BB%BA%E4%B8%80%E4%B8%AA-propertykey
//
func newPropertyKeysCreateFunc(t hgapi.Transport) PropertyKeysCreate {
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
	ctx         context.Context         `json:"-"`
	Name        string                  `json:"name"`
	DataType    PropertyKeyDataType     `json:"data_type"`
	Cardinality PropertyCardinalityType `json:"cardinality"`
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

func (r PropertyKeysCreateRequest) Do(ctx context.Context, transport hgapi.Transport) (*PropertyKeysCreateResponse, error) {

	bytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	byteBody, _ := json.Marshal(&r)               // 序列化
	reader := strings.NewReader(string(byteBody)) // 转化为reader
	req, _ := hgapi.NewRequest("POST", "/graphs/${GRAPH_NAME}/schema/propertykeys", reader)

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

func (v PropertyKeysCreate) WithName(name string) func(*PropertyKeysCreateRequest) {
	return func(r *PropertyKeysCreateRequest) {
		r.Name = name
	}
}

func (v PropertyKeysCreate) WithDataType(dataType PropertyKeyDataType) func(*PropertyKeysCreateRequest) {
	return func(r *PropertyKeysCreateRequest) {
		r.DataType = dataType
	}
}

func (v PropertyKeysCreate) WithCardinality(cardinality PropertyCardinalityType) func(*PropertyKeysCreateRequest) {
	return func(r *PropertyKeysCreateRequest) {
		r.Cardinality = cardinality
	}
}
