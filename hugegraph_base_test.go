package hugegraph_test

import (
	"fmt"
	"hugegraph"
	"hugegraph/hgapi"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

var defaultResponse = http.Response{
	Status:        "200 OK",
	StatusCode:    200,
	ContentLength: 2,
	Header:        http.Header(map[string][]string{"Content-Type": {"application/json"}}),
	Body:          ioutil.NopCloser(strings.NewReader(`{}`)),
}

func initClient() *hugegraph.Client {

	client, err := hugegraph.NewDefaultClient()

	if err != nil {
		log.Fatalf("Error creating the client: %s\n", err)
	}

	return client
}

func TestClient(t *testing.T) {

	client := initClient()

	res, err := client.Version()
	if err != nil {
		log.Fatalf("Error getting the response: %s\n", err)
	}
	defer res.Body.Close()

	fmt.Println(res.Versions)

	fmt.Println(res.Versions.Version)
}

type VertexCase1 struct {
	Name string
}

func TestVertexById(t *testing.T) {

	client := initClient()

	res, err := client.VertexGetID(
		client.VertexGetID.WithID("1"),
		client.VertexGetID.WithLabel("vertex"),
	)
	if err != nil {
		log.Fatalf("Error getting the response: %s\n", err)
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error getting the response: %s\n", err)
	}

	fmt.Println("我是结果=>" + string(bytes))
}

func TestSchemaGet(t *testing.T) {

	client := initClient()

	res, err := client.SchemaGet()
	if err != nil {
		log.Fatalf("Error getting the response: %s\n", err)
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error getting the response: %s\n", err)
	}

	fmt.Println("我是结果=>" + string(bytes))
}

func TestPropertyKeysCreate(t *testing.T) {
	client := initClient()

	res, err := client.PropertyKeys.Create(
		client.PropertyKeys.Create.WithName("title"),
		client.PropertyKeys.Create.WithDataType(hgapi.PropertyDataTypeInt),
		client.PropertyKeys.Create.WithCardinality(hgapi.PropertyCardinalityTypeSingle),
	)
	if err != nil {
		log.Fatalf("Error getting the response: %s\n", err)
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error getting the response: %s\n", err)
	}

	fmt.Println("我是结果=>" + string(bytes))
}

func TestPropertyKeysGet(t *testing.T) {
	client := initClient()

	res, err := client.PropertyKeys.Get()
	if err != nil {
		log.Fatalf("Error getting the response: %s\n", err)
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error getting the response: %s\n", err)
	}

	fmt.Println("我是结果=>" + string(bytes))
}

func TestPropertyKeysDeleteByName(t *testing.T) {
	client := initClient()

	res, err := client.PropertyKeys.DeleteByName(
		client.PropertyKeys.DeleteByName.WithName("age"),
	)
	if err != nil {
		log.Fatalf("Error getting the response: %s\n", err)
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error getting the response: %s\n", err)
	}

	fmt.Println("我是结果=>" + string(bytes))
}

func TestPropertyKeysGetByName(t *testing.T) {
	client := initClient()

	res, err := client.PropertyKeys.GetByName(
		client.PropertyKeys.GetByName.WithName("title"),
	)
	if err != nil {
		log.Fatalf("Error getting the response: %s\n", err)
	}

	fmt.Printf("我是结果=>%+v\n", res.Data)
}

func TestProperKeysUpdateUserdata(t *testing.T) {

	client := initClient()

	res, err := client.PropertyKeys.UpdateUserdata(
		client.PropertyKeys.UpdateUserdata.WithName("title"),
		client.PropertyKeys.UpdateUserdata.WithUserdata(hgapi.PropertyKeysUpdateUserData{
			Min: 1,
			Max: 255,
		}),
	)
	if err != nil {
		log.Fatalf("Error getting the response: %s\n", err)
	}

	fmt.Printf("我是结果=>%+v\n", res.Data)
}
