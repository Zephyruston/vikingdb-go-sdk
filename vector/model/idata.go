// Copyright (c) 2025 Beijing Volcano Engine Technology Co., Ltd.
// SPDX-License-Identifier: Apache-2.0

package model

// FetchDataInIndexRequest 表示在索引中获取数据的请求
type FetchDataInIndexRequest struct {
	// 数据ID列表
	IDs []interface{} `json:"ids"`

	// 分区
	Partition interface{} `json:"partition,omitempty"`

	// 输出字段
	OutputFields []string `json:"output_fields,omitempty"`
}

type IndexDataItem struct {
	DataItem
	// 稠密向量维度
	DenseDim int `json:"dense_dim,omitempty"`
	// 稠密向量
	DenseVector []float32 `json:"dense_vector,omitempty"`
}

// FetchDataInIndexResponse 表示在索引中获取数据的响应
type FetchDataInIndexResponse struct {
	// 通用响应
	CommonResponse

	// 数据项列表
	Result *FetchDataInIndexResult `json:"result,omitempty"`
}

type FetchDataInIndexResult struct {
	// 数据项列表
	Datas []IndexDataItem `json:"fetch,omitempty"`

	// 不存在的主键列表
	NotFoundIDs []interface{} `json:"ids_not_exist,omitempty"`
}

// RecallBase 表示检索基础请求
type RecallBase struct {
	// 分区
	// Partition interface{} `json:"partition,omitempty"`

	// 过滤条件
	Filter MapStr `json:"filter,omitempty"`
}

// SearchBase 表示搜索基础请求
type SearchBase struct {
	// 检索基础请求
	RecallBase

	// 输出字段
	OutputFields []string `json:"output_fields,omitempty"`

	// 限制返回的数据条数
	Limit *int `json:"limit,omitempty"`

	// 偏移量
	Offset *int `json:"offset,omitempty"`

	// 高级搜索选项
	Advance *SearchAdvance `json:"advance,omitempty"`
}

// SearchAdvance 表示高级搜索选项
type SearchAdvance struct {
	// 密集向量权重
	DenseWeight *float64 `json:"dense_weight,omitempty"`

	// 包含的ID列表
	IDsIn []interface{} `json:"ids_in,omitempty"`

	// 排除的ID列表
	IDsNotIn []interface{} `json:"ids_not_in,omitempty"`

	// 后处理操作
	PostProcessOps []MapStr `json:"post_process_ops,omitempty"`

	// 后处理输入限制
	PostProcessInputLimit *int `json:"post_process_input_limit,omitempty"`

	// 比例系数K
	ScaleK *float64 `json:"scale_k,omitempty"`

	// 过滤前ANN限制
	FilterPreAnnLimit *int `json:"filter_pre_ann_limit,omitempty"`

	// 过滤前ANN比例
	FilterPreAnnRatio *float64 `json:"filter_pre_ann_ratio,omitempty"`
}

type SearchResponse struct {
	CommonResponse

	Result *SearchResult `json:"result,omitempty"`
}

type SearchResult struct {
	// 数据项列表
	Data []SearchItemResult `json:"data,omitempty"`

	// 过滤匹配数量
	FilterMatchCount int `json:"filter_match_count,omitempty"`

	// 总返回数量
	TotalReturnCount int `json:"total_return_count,omitempty"`

	// 真实查询文本
	RealTextQuery string `json:"real_text_query,omitempty"`

	// token 使用情况
	TokenUsage MapStr `json:"token_usage,omitempty"`
}

// SearchResult 表示搜索结果
type SearchItemResult struct {
	// 数据ID
	ID interface{} `json:"id"`

	// fields
	Fields MapStr `json:"fields,omitempty"`

	// ANN 分数
	ANNScore float32 `json:"ann_score,omitempty"`

	// 分数
	Score float32 `json:"score,omitempty"`
}

// SearchByVectorRequest 表示通过向量搜索的请求
type SearchByVectorRequest struct {
	// 搜索基础请求
	SearchBase

	// 密集向量
	Dense []float64 `json:"dense_vector"`

	// 稀疏向量
	Sparse map[string]float64 `json:"sparse_vector,omitempty"`
}

// SearchByMultiModalRequest 表示通过多模态搜索的请求
type SearchByMultiModalRequest struct {
	// 搜索基础请求
	SearchBase

	// 文本
	Text *string `json:"text,omitempty"`

	// 图像
	Image *string `json:"image,omitempty"`

	// 是否需要指令
	NeedInstruction *bool `json:"need_instruction,omitempty"`
}

// SearchByIDRequest 表示通过ID搜索的请求
type SearchByIDRequest struct {
	// 搜索基础请求
	SearchBase

	// ID
	ID interface{} `json:"id"`
}

// SearchByScalarRequest 表示通过标量搜索的请求
type SearchByScalarRequest struct {
	// 搜索基础请求
	SearchBase

	// 字段
	Field *string `json:"field,omitempty"`

	// 排序顺序
	Order *string `json:"order,omitempty"`
}

// SearchByKeywordsRequest 表示通过关键词搜索的请求
type SearchByKeywordsRequest struct {
	// 搜索基础请求
	SearchBase

	// 查询
	Query interface{} `json:"query"`

	// 是否区分大小写
	CaseSensitive bool `json:"case_sensitive,omitempty"`
}

// SearchByRandomRequest 表示随机搜索的请求
type SearchByRandomRequest struct {
	// 搜索基础请求
	SearchBase
}

// AggRequest 表示聚合请求
type AggRequest struct {
	// 检索基础请求
	RecallBase

	// 操作
	Op *string `json:"op"`

	// 字段
	Field *string `json:"field,omitempty"`

	// 条件
	Cond MapStr `json:"cond,omitempty"`

	// 排序顺序
	Order *string `json:"order,omitempty"`
}

// AggResponse 表示聚合响应
type AggResponse struct {
	// 通用响应
	CommonResponse

	// 聚合结果
	Result *AggResult `json:"result,omitempty"`
}

type AggResult struct {
	// 聚合结果
	Agg MapStr `json:"agg,omitempty"`

	// 算子
	Op string `json:"op,omitempty"`

	// 字段
	Field string `json:"field,omitempty"`
}

// SortRequest 表示排序请求
type SortRequest struct {
	// 查询向量
	QueryVector []float64 `json:"query_vector"`

	// ID列表
	IDs []interface{} `json:"ids"`
}

// SortResponse 表示排序响应
type SortResponse struct {
	// 通用响应
	CommonResponse

	// 排序结果
	Result *SortResult `json:"result,omitempty"`
}

type SortResult struct {
	// 数据项列表
	Datas []SortItem `json:"fetch,omitempty"`

	// 不存在的主键列表
	NotFoundIDs []interface{} `json:"ids_not_exist,omitempty"`
}

type SortItem struct {
	// 数据ID
	ID interface{} `json:"id"`

	// 分数
	Score float32 `json:"score,omitempty"`
}
