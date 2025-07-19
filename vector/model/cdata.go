package model

import "time"

// Collection 表示一个向量集合
type Collection struct {
	// 集合ID
	ID string `json:"id"`

	// 集合名称
	Name string `json:"name"`

	// 集合描述
	Description string `json:"description,omitempty"`

	// 向量维度
	Dimension int `json:"dimension"`

	// 向量类型
	VectorType string `json:"vector_type,omitempty"`

	// 索引类型
	IndexType string `json:"index_type,omitempty"`

	// 索引参数
	IndexParams map[string]interface{} `json:"index_params,omitempty"`

	// 元数据字段
	Fields []Field `json:"fields,omitempty"`

	// 创建时间
	CreatedAt time.Time `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// Field 表示集合中的一个字段
type Field struct {
	// 字段名称
	Name string `json:"name"`

	// 字段类型
	Type string `json:"type"`

	// 是否为主键
	PrimaryKey bool `json:"primary_key,omitempty"`

	// 是否可以为空
	Nullable bool `json:"nullable,omitempty"`

	// 是否可以索引
	Indexable bool `json:"indexable,omitempty"`

	// 是否存储
	Store bool `json:"store,omitempty"`
}

// DataItem 表示一个数据项
type DataItem struct {
	// 数据ID
	ID interface{} `json:"id"`
	// 数据字段
	Fields MapStr `json:"fields"`
}

// WriteDataBase 表示写入数据的基础请求
type WriteDataBase struct {
	// 数据字段
	Data []MapStr `json:"data"`

	// 过期时间（秒）
	TTL *int32 `json:"ttl,omitempty"`

	// 是否忽略未知字段
	IgnoreUnknownFields bool `json:"ignore_unknown_fields,omitempty"`
}

// UpsertDataRequest 表示插入或更新数据的请求
type UpsertDataRequest struct {
	// 写入数据基础请求
	WriteDataBase

	// 是否异步处理
	Async bool `json:"async"`
}

// UpsertDataResponse 表示插入或更新数据的响应
type UpsertDataResponse struct {
	// 通用响应
	CommonResponse

	// Token使用情况
	TokenUsage interface{} `json:"token_usage,omitempty"`
}

// UpdateDataRequest 表示更新数据的请求
type UpdateDataRequest struct {
	// 写入数据基础请求
	WriteDataBase
}

// UpdateDataResponse 表示更新数据的响应
type UpdateDataResponse struct {
	// 通用响应
	CommonResponse
}

// DeleteDataRequest 表示删除数据的请求
type DeleteDataRequest struct {
	// 数据ID列表
	IDs []interface{} `json:"ids"`

	// 是否删除所有数据
	DelAll bool `json:"del_all,omitempty"`
}

// DeleteDataResponse 表示删除数据的响应
type DeleteDataResponse struct {
	// 通用响应
	CommonResponse
}

// FetchDataInCollectionRequest 表示在集合中获取数据的请求
type FetchDataInCollectionRequest struct {
	// 数据ID列表
	IDs []interface{} `json:"ids"`

	// 是否返回下载URL
	ReturnDownloadURL bool `json:"return_download_url,omitempty"`
}

// FetchDataInCollectionResponse 表示在集合中获取数据的响应
type FetchDataInCollectionResponse struct {
	// 通用响应
	CommonResponse

	// 数据内容
	Result *FetchDataInCollectionResult `json:"result,omitempty"`
}

type FetchDataInCollectionResult struct {
	// 数据列表
	Datas []DataItem `json:"fetch,omitempty"`

	// 不存在的主键列表
	NotFoundIDs []interface{} `json:"ids_not_exist,omitempty"`
}
