package model

// CommonResponse 表示通用响应结构
type CommonResponse struct {
	// API名称
	API string `json:"api,omitempty"`
	
	// 消息
	Message string `json:"message,omitempty"`

	// code 
	Code string `json:"string,omitempty"`

	// log ID
	RequestID string `json:"request_id,omitempty"`
}

// DataAPICollectionBase 表示数据API集合基础请求
type DataAPICollectionBase struct {
	// 集合名称
	CollectionName string `json:"collection_name"`
	
	// 项目名称
	ProjectName string `json:"project_name"`
	
	// 资源ID
	ResourceID string `json:"resource_id"`
}

// DataAPIIndexBase 表示数据API索引基础请求
type DataAPIIndexBase struct {
	// 集合基础请求
	DataAPICollectionBase
	
	// 索引名称
	IndexName string `json:"index_name"`
}

// Refer 表示引用信息
type Refer struct {
	// 账户ID
	AccountID string `json:"account_id"`
	
	// 实例编号
	InstanceNO string `json:"instance_no"`
	
	// 资源ID
	ResourceID string `json:"resource_id"`
}

// MapStr 表示字符串映射
type MapStr map[string]interface{}

// PaginationRequest 表示分页请求
type PaginationRequest struct {
	// 页码
	Page int `json:"page,omitempty"`
	
	// 每页大小
	PageSize int `json:"page_size,omitempty"`
}

// PaginationResponse 表示分页响应
type PaginationResponse struct {
	// 总数
	Total int `json:"total"`
	
	// 页码
	Page int `json:"page"`
	
	// 每页大小
	PageSize int `json:"page_size"`
}