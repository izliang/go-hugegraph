package model

// Vertex 顶点的基础模型
type Vertex[I, V] struct {
	Id         I
	Label      string
	Properties V
}

type Edge[I, SID, TID, V] struct {
	ID         I
	Label      string
	SourceType string
	TargetType string
	SourceID   SID
	Target     TID
	Properties V
}
