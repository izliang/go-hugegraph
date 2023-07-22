package hugegraph_test

import (
	"fmt"
	"github.com/izliang/go-hugegraph"
	"github.com/izliang/go-hugegraph/hgapi/v1"
	"io"
	"io/ioutil"
	"log"
	"testing"
)

func initClient() *hugegraph.CommonClient {

	client, err := hugegraph.NewDefaultCommonClient()

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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	fmt.Println(res.Versions)

	fmt.Println(res.Versions.Version)
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error getting the response: %s\n", err)
	}

	fmt.Println("我是结果=>" + string(bytes))
}

func TestSchemaGet(t *testing.T) {

	client := initClient()

	res, err := client.SchemaGet(
		client.SchemaGet.WithFormat("groovy"),
	)
	if err != nil {
		log.Fatalf("Error getting the response: %s\n", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

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
		client.PropertyKeys.Create.WithDataType(v1.PropertyDataTypeInt),
		client.PropertyKeys.Create.WithCardinality(v1.PropertyCardinalityTypeSingle),
	)
	if err != nil {
		log.Fatalf("Error getting the response: %s\n", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error getting the response: %s\n", err)
	}

	fmt.Println("我是结果=>" + string(bytes))
}

func TestPropertyKeysDeleteByName(t *testing.T) {
	client := initClient()

	res, err := client.PropertyKeys.DeleteByName(
		client.PropertyKeys.DeleteByName.WithName("title"),
	)
	if err != nil {
		log.Fatalf("Error getting the response: %s\n", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

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
		client.PropertyKeys.UpdateUserdata.WithAction(v1.PropertyKeyActionAppend),
		client.PropertyKeys.UpdateUserdata.WithUserdata(v1.PropertyKeysUpdateUserData{
			Min: 1,
			Max: 255,
		}),
	)
	if err != nil {
		log.Fatalf("Error getting the response: %s\n", err)
	}

	fmt.Printf("我是结果=>%+v\n", res.Data)
}

func TestVertexLabelCreate(t *testing.T) {
	client := initClient()

	res, err := client.VertexLabel.Create(
		client.VertexLabel.Create.WithData(v1.VertexLabelCreateRequestData{
			Name:             "vertex",
			IDStrategy:       v1.VertexLabelIDStrategyTypeAutomatic,
			Properties:       []string{"title"},
			PrimaryKeys:      nil,
			NullableKeys:     nil,
			EnableLabelIndex: true,
		}),
	)
	if err != nil {
		log.Fatalf("Error getting the response: %s\n", err)
	}

	fmt.Printf("我是结果=>%+v\n", res.Data)
}

func TestVertexLabelCreate1(t *testing.T) {
	client := initClient()

	res, err := client.VertexLabel.Create(
		client.VertexLabel.Create.WithData(v1.VertexLabelCreateRequestData{
			Name:             "vertex",
			IDStrategy:       v1.VertexLabelIDStrategyTypeCustomizeString,
			Properties:       []string{"title"},
			PrimaryKeys:      nil,
			NullableKeys:     nil,
			EnableLabelIndex: true,
		}),
	)
	if err != nil {
		log.Fatalf("Error getting the response: %s\n", err)
	}

	fmt.Printf("我是结果=>%+v\n", res.Data)
}

func TestGremlinGet(t *testing.T) {
	client := initClient()

	res, err := client.Gremlin.Get(
		client.Gremlin.Get.WithGremlinGetData(v1.GremlinGetRequestReqData{
			Gremlin: "lemma.traversal().V().limit(10)",
		}),
	)
	if err != nil {
		log.Fatalf("Error getting the response: %s\n", err)
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error getting the response: %s\n", err)
	}
	fmt.Println("我是结果=>" + string(bytes))
}

func TestGremlinPost(t *testing.T) {
	client := initClient()
	_, err := client.Gremlin.Post(
		client.Gremlin.Post.WithGremlinPostData(v1.GremlinPostRequestReqData{
			Gremlin: "g.V().limit(10)",
		}),
	)
	if err != nil {
		return
	}
}

type QueryCase struct {
	Word             string
	FeatureId        int64
	RelationCategory string
	Limit            int
}

func TestGremlinSuggest(t *testing.T) {
	client := initClient()

	queryCase := &QueryCase{
		Word:             "葡萄",
		FeatureId:        3724,
		RelationCategory: "treeItemTreeItem",
		Limit:            15,
	}

	_, err := client.Gremlin.Post(
		client.Gremlin.Post.WithGremlinPostData(v1.GremlinPostRequestReqData{
			Gremlin:  fmt.Sprintf("g.V().hasLabel('lemma').has('lemmaTitle',Text.contains('%s')).where(bothE().has('featureId',__.unfold().is(%d)).has('relationCategory','%s')).dedup().limit(%d)", queryCase.Word, queryCase.FeatureId, queryCase.RelationCategory, queryCase.Limit),
			Bindings: nil,
			Aliases: struct {
				Graph string `json:"graph"`
				G     string `json:"g"`
			}{
				Graph: "baike-lemma",
				G:     "__g_baike-lemma",
			},
		}),
	)
	if err != nil {
		return
	}
}
