package hgapi

// Code generated from specification version 5.6.15 (fe7575a32e2): DO NOT EDIT

// API contains the Elasticsearch APIs
//
type API struct {
	Version            Version
	VertexGetID        VertexGetID
	SchemaGet          SchemaGet
	PropertyKeysCreate PropertyKeysCreate
	PropertyKeysGet    PropertyKeysGet
}

// New creates new API
//
func New(t Transport) *API {
	return &API{
		Version:            newVersionFunc(t),
		VertexGetID:        newVertexGetIDFunc(t),
		SchemaGet:          newSchemaGetFunc(t),
		PropertyKeysCreate: newPropertyKeysCreateFunc(t),
		PropertyKeysGet:    newPropertyKeysGetFunc(t),
	}
}
