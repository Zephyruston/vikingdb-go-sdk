// Copyright (c) 2025 Beijing Volcano Engine Technology Co., Ltd.
// SPDX-License-Identifier: Apache-2.0

package model

import "time"

// Index 表示一个索引
type Index struct {
	// 索引ID
	ID string `json:"id"`

	// 索引名称
	Name string `json:"name"`

	// 索引描述
	Description string `json:"description,omitempty"`

	// 索引类型
	IndexType string `json:"index_type"`

	// 索引参数
	IndexParams map[string]interface{} `json:"index_params,omitempty"`

	// 所属集合ID
	CollectionID string `json:"collection_id"`

	// 所属集合名称
	CollectionName string `json:"collection_name"`

	// 创建时间
	CreatedAt time.Time `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// CreateIndexRequest 表示创建索引的请求
type CreateIndexRequest struct {
	// 集合名称
	CollectionName *string `json:"collection_name"`

	// 索引名称
	Name string `json:"name"`

	// 索引描述
	Description string `json:"description,omitempty"`

	// 索引类型
	IndexType string `json:"index_type"`

	// 索引参数
	IndexParams map[string]interface{} `json:"index_params,omitempty"`
}

// CreateIndexResponse 表示创建索引的响应
type CreateIndexResponse struct {
	// 通用响应
	CommonResponse

	// 索引信息
	Index *Index `json:"index,omitempty"`
}

// GetIndexRequest 表示获取索引的请求
type GetIndexRequest struct {
	// 集合名称
	CollectionName *string `json:"collection_name"`

	// 索引名称
	Name string `json:"name"`
}

// GetIndexResponse 表示获取索引的响应
type GetIndexResponse struct {
	// 通用响应
	CommonResponse

	// 索引信息
	Index *Index `json:"index,omitempty"`
}

// ListIndexesRequest 表示列出索引的请求
type ListIndexesRequest struct {
	// 集合名称
	CollectionName *string `json:"collection_name"`

	// 分页请求
	PaginationRequest

	// 名称前缀
	NamePrefix string `json:"name_prefix,omitempty"`
}

// ListIndexesResponse 表示列出索引的响应
type ListIndexesResponse struct {
	// 通用响应
	CommonResponse

	// 分页响应
	PaginationResponse

	// 索引列表
	Indexes []*Index `json:"indexes,omitempty"`
}

// DeleteIndexRequest 表示删除索引的请求
type DeleteIndexRequest struct {
	// 集合名称
	CollectionName *string `json:"collection_name"`

	// 索引名称
	Name string `json:"name"`
}

// DeleteIndexResponse 表示删除索引的响应
type DeleteIndexResponse struct {
	// 通用响应
	CommonResponse
}

// UpdateIndexRequest 表示更新索引的请求
type UpdateIndexRequest struct {
	// 集合名称
	CollectionName *string `json:"collection_name"`

	// 索引名称
	Name string `json:"name"`

	// 索引描述
	Description string `json:"description,omitempty"`

	// 索引参数
	IndexParams map[string]interface{} `json:"index_params,omitempty"`
}

// UpdateIndexResponse 表示更新索引的响应
type UpdateIndexResponse struct {
	// 通用响应
	CommonResponse

	// 索引信息
	Index *Index `json:"index,omitempty"`
}
