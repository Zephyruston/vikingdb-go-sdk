// Copyright (c) 2025 Beijing Volcano Engine Technology Co., Ltd.
// SPDX-License-Identifier: Apache-2.0

package model

// EmbeddingModelOpt 表示嵌入模型选项
type EmbeddingModelOpt struct {
	// 模型名称
	ModelName *string `json:"name"`

	// 模型版本
	ModelVersion *string `json:"version,omitempty"`

	// 维度
	Dim *int `json:"dim,omitempty"`
}

// EmbeddingData 表示嵌入数据
type EmbeddingData struct {
	// 文本
	Text *string `json:"text,omitempty"`

	// 图像
	Image *string `json:"image,omitempty"`
}

// EmbeddingRequest 表示嵌入请求
type EmbeddingRequest struct {
	// 项目名称
	ProjectName *string `json:"project_name,omitempty"`

	// 密集模型
	DenseModel *EmbeddingModelOpt `json:"dense_model,omitempty"`

	// 稀疏模型
	SparseModel *EmbeddingModelOpt `json:"sparse_model,omitempty"`

	// 数据
	Data []*EmbeddingData `json:"data"`
}

// EmbeddingResponse 表示嵌入响应
type EmbeddingResponse struct {
	// 通用响应
	CommonResponse

	Result *EmbeddingResult `json:"result,omitempty"`
}

type EmbeddingResult struct {
	Data []*Embedding `json:"data"`

	// Token使用情况
	TokenUsage interface{} `json:"token_usage,omitempty"`
}

// Embedding 表示嵌入结果
type Embedding struct {
	// 密集向量
	DenseVectors [][]float32 `json:"dense_vectors,omitempty"`

	// 稀疏向量
	SparseVectors []map[string]float32 `json:"sparse_vectors,omitempty"`
}
