package v1

import "github.com/izliang/go-hugegraph/hgapi"

// Code generated from specification version 5.6.15 (fe7575a32e2): DO NOT EDIT

// API contains the Elasticsearch APIs
//
type APIV1 struct {
	Version      Version
	VertexGetID  VertexGetID
	SchemaGet    SchemaGet
	PropertyKeys *PropertyKeys
	VertexLabel  *VertexLabel
	Gremlin      *Gremlin
}

type PropertyKeys struct {
	Create         PropertyKeysCreate
	Get            PropertyKeysGet
	DeleteByName   PropertyKeysDeleteByName
	GetByName      PropertyKeysGetByName
	UpdateUserdata PropertyKeysUpdateUserdata
}

type VertexLabel struct {
	Create VertexLabelCreate
}
type Gremlin struct {
	Get  GremlinGet
	Post GremlinPost
}

// New creates new API
//
func New(t hgapi.Transport) *APIV1 {
	return &APIV1{
		Version:     newVersionFunc(t),
		VertexGetID: newVertexGetIDFunc(t),
		SchemaGet:   newSchemaGetFunc(t),
		PropertyKeys: &PropertyKeys{
			Create:         newPropertyKeysCreateFunc(t),
			Get:            newPropertyKeysGetFunc(t),
			DeleteByName:   newPropertyKeysDeleteByNameFunc(t),
			GetByName:      newPropertyKeysGetByNameFunc(t),
			UpdateUserdata: newPropertyKeysUpdateUserdataFunc(t),
		},
		VertexLabel: &VertexLabel{
			Create: newVertexLabelCreateFunc(t),
		},
		Gremlin: &Gremlin{
			Get:  newGremlinGetFunc(t),
			Post: newGremlinPostFunc(t),
		},
	}
}
