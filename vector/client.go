package vector

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/volcengine/vikingdb-go-sdk/vector/model"
	"github.com/volcengine/vikingdb-go-sdk/vector/utils"
)

// client 是 Client 接口的实现
type client struct {
	config   *Config
	httpCli  *http.Client
}

func NewIndexClientWithAkSk(accessKey, secretKey string, indexConfig model.DataAPIIndexBase, opts ...ClientOption) (IndexClient, error) {
	if accessKey == "" || secretKey == "" {
		return nil, model.NewInvalidParameterError("accessKey and secretKey cannot be empty")
	}

	// 创建默认配置
	cfg := DefaultConfig()

	// 应用选项
	for _, opt := range opts {
		opt(cfg)
	}

	// 设置认证信息
	cfg.AccessKey = accessKey
	cfg.SecretKey = secretKey
	cfg.AuthType = AuthTypeAkSk

	// 创建 HTTP 客户端
	httpCli := &http.Client{
		Timeout: time.Duration(cfg.Timeout) * time.Millisecond,
	}

	client := &client{
		config:  cfg,
		httpCli: httpCli,
	}

	iClient := &indexClient{
		client: client,
		indexBase: indexConfig,
	}

	return iClient, nil
}

func NewCollectionClientWithAkSk(accessKey, secretKey string, collectionConfig model.DataAPICollectionBase, opts ...ClientOption) (CollectionClient, error) {
	if accessKey == "" || secretKey == "" {
		return nil, model.NewInvalidParameterError("accessKey and secretKey cannot be empty")
	}

	// 创建默认配置
	cfg := DefaultConfig()

	// 应用选项
	for _, opt := range opts {
		opt(cfg)
	}

	// 设置认证信息
	cfg.AccessKey = accessKey
	cfg.SecretKey = secretKey
	cfg.AuthType = AuthTypeAkSk

	// 创建 HTTP 客户端
	httpCli := &http.Client{
		Timeout: time.Duration(cfg.Timeout) * time.Millisecond,
	}
	client := &client{
		config:  cfg,
		httpCli: httpCli,
	}
	cClient := &collectionClient{	
		client: client,
		collectionBase: collectionConfig,
	}
	return cClient, nil
}

func NewEmbeddingClientWithAkSk(accessKey, secretKey string, opts ...ClientOption) (EmbeddingClient, error) {
	if accessKey == "" || secretKey == "" {
		return nil, model.NewInvalidParameterError("accessKey and secretKey cannot be empty")
	}
	
	// 创建默认配置
	cfg := DefaultConfig()

	// 应用选项
	for _, opt := range opts {
		opt(cfg)
	}

	// 设置认证信息
	cfg.AccessKey = accessKey
	cfg.SecretKey = secretKey
	cfg.AuthType = AuthTypeAkSk

	// 创建 HTTP 客户端
	httpCli := &http.Client{
		Timeout: time.Duration(cfg.Timeout) * time.Millisecond,
	}
	client := &client{
		config:  cfg,
		httpCli: httpCli,
	}
	return &embeddingClient{
		client: client,
	}, nil
}

// Close 关闭客户端
func (c *client) Close() error {
	// 目前没有需要关闭的资源，返回 nil
	return nil
}

// doRequest 执行 HTTP 请求
func (c *client) doRequest(ctx context.Context, method, path string, request, response interface{}, opts ...RequestOption) error {
	// 创建请求选项
	reqOpts := defaultRequestOptions()

	// 应用选项
	for _, opt := range opts {
		opt(reqOpts)
	}

	// 如果没有指定重试次数，使用配置中的重试次数
	if reqOpts.MaxRetries == 0 {
		reqOpts.MaxRetries = c.config.MaxRetries
	}

	// 执行请求，带重试
	return utils.Retry(reqOpts.MaxRetries, func() error {
		// 构建请求 URL
		url := fmt.Sprintf("%s%s", c.config.Endpoint, path)

		// 构建请求体
		var reqBody []byte
		var err error
		if request != nil {
			reqBody, err = utils.SerilizeToJsonBytesUseNumber(request)
			if err != nil {
				return model.NewErrorWithCause(model.ErrCodeInvalidParameter, "failed to marshal request", err, http.StatusBadRequest)
			}
		}

		// 构建请求
		req, err := http.NewRequestWithContext(ctx, method, url, nil)
		if err != nil {
			return model.NewErrorWithCause(model.ErrCodeUnknown, "failed to create request", err, http.StatusBadRequest)
		}

		// 设置请求头
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		req.Header.Set("User-Agent", fmt.Sprintf("vikingdb-go-sdk/%s", Version))

		// 如果有请求体，设置请求体
		if len(reqBody) > 0 {
			req.Body = io.NopCloser(bytes.NewReader(reqBody))
			req.ContentLength = int64(len(reqBody))
			req.Header.Set("Content-Length", strconv.Itoa(len(reqBody)))
		}

		// 添加认证信息
		switch c.config.AuthType {
		case AuthTypeAkSk:
			req = utils.SignRequest(req, c.config.AccessKey, c.config.SecretKey)
		}

		// 执行请求
		resp, err := utils.DoHTTPRequest(c.httpCli, req)
		if err != nil {
			return model.NewErrorWithCause(model.ErrCodeHTTPRequestFailed, "failed to do http request", err, resp.StatusCode)
		}
		defer resp.Body.Close()

		// 解析响应
		return utils.ParseResponse(resp, response)
	}, utils.IsRetryableError)
}