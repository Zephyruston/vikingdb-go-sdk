package utils

import (
	"fmt"
	"io"
	"net/http"

	"github.com/volcengine/vikingdb-go-sdk/vector/model"
)

// DoHTTPRequest 执行 HTTP 请求
func DoHTTPRequest(client *http.Client, req *http.Request) (*http.Response, error) {
	// 执行请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ParseResponse 解析 HTTP 响应
func ParseResponse(resp *http.Response, result interface{}) error {
	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.NewErrorWithCause(model.ErrCodeUnknown, "failed to read response body", err, http.StatusInternalServerError)
	}

	// 检查状态码
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		// 尝试解析错误响应
		var errResp struct {
			Code    string    `json:"code"`
			Message string `json:"message"`
		}
		if err = ParseJsonUseNumber2(body, &errResp); err == nil && errResp.Message != "" {
			ParseJsonUseNumber2(body, result)
			return model.NewErrorWithCause(model.ErrorCode(errResp.Code), errResp.Message, err, resp.StatusCode)
		}
		// 如果无法解析错误响应，返回通用错误
		return model.NewErrorWithCause(model.ErrCodeUnknown, fmt.Sprintf("unmarshal response failed: %s", string(body)), err, resp.StatusCode)
	}

	// 如果没有结果对象，直接返回
	if result == nil {
		return nil
	}

	// 解析响应体
	if err := ParseJsonUseNumber2(body, result); err != nil {
		return model.NewErrorWithCause(model.ErrCodeUnknown, "failed to unmarshal response body", err, http.StatusInternalServerError)
	}

	return nil
}