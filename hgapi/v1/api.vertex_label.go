package v1

type VertexLabelIDStrategyType string

const (
	// VertexLabelIDStrategyTypeDefault 默认策略,默认为主键ID
	VertexLabelIDStrategyTypeDefault VertexLabelIDStrategyType = "DEFAULT"
	// VertexLabelIDStrategyTypePrimaryKey 主键ID
	VertexLabelIDStrategyTypePrimaryKey VertexLabelIDStrategyType = "PRIMARY_KEY"
	// VertexLabelIDStrategyTypeCustomizeUUID 自定义UUID
	VertexLabelIDStrategyTypeCustomizeUUID VertexLabelIDStrategyType = "CUSTOMIZE_UUID"
	// VertexLabelIDStrategyTypeAutomatic 自动生成
	VertexLabelIDStrategyTypeAutomatic VertexLabelIDStrategyType = "AUTOMATIC"
	// VertexLabelIDStrategyTypeCustomizeNumber 自定义数字
	VertexLabelIDStrategyTypeCustomizeNumber VertexLabelIDStrategyType = "CUSTOMIZE_NUMBER"
	// VertexLabelIDStrategyTypeCustomizeString 自定义字符串
	VertexLabelIDStrategyTypeCustomizeString VertexLabelIDStrategyType = "CUSTOMIZE_STRING"
)
