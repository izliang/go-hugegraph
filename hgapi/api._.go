package hgapi

// Code generated from specification version 5.6.15 (fe7575a32e2): DO NOT EDIT

// API contains the Elasticsearch APIs
//
type API struct {
	Version      Version
	VertexGetID  VertexGetID
	SchemaGet    SchemaGet
	PropertyKeys *PropertyKeys
}

type PropertyKeys struct {
	Create         PropertyKeysCreate
	Get            PropertyKeysGet
	DeleteByName   PropertyKeysDeleteByName
	GetByName      PropertyKeysGetByName
	UpdateUserdata PropertyKeysUpdateUserdata
}

// New creates new API
//
func New(t Transport) *API {
	return &API{
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
	}
}
