// Copyright (c) 2025 Beijing Volcano Engine Technology Co., Ltd.
// SPDX-License-Identifier: Apache-2.0

package vector

// RequestOptions 表示请求选项
type RequestOptions struct {
	// 最大重试次数
	MaxRetries int

	// 请求头
	Headers map[string]string
}

// RequestOption 表示请求选项函数
type RequestOption func(*RequestOptions)

// defaultRequestOptions 返回默认请求选项
func defaultRequestOptions() *RequestOptions {
	return &RequestOptions{
		Headers: make(map[string]string),
	}
}

// WithRequestMaxRetries 设置请求的最大重试次数
func WithRequestMaxRetries(maxRetries int) RequestOption {
	return func(o *RequestOptions) {
		o.MaxRetries = maxRetries
	}
}

// WithRequestHeader 设置请求头
func WithRequestHeader(key, value string) RequestOption {
	return func(o *RequestOptions) {
		o.Headers[key] = value
	}
}
