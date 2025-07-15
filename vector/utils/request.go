package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// CreateHTTPRequest 创建 HTTP 请求
func CreateHTTPRequest(method, url string, body interface{}) (*http.Request, []byte, error) {
	var reqBody []byte
	var err error
	
	// 如果有请求体，序列化为 JSON
	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, nil, err
		}
	}
	
	// 创建请求
	req, err := http.NewRequest(method, url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, nil, err
	}
	
	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	if len(reqBody) > 0 {
		req.Header.Set("Content-Length", strconv.Itoa(len(reqBody)))
	}
	
	return req, reqBody, nil
}

// ReadResponseBody 读取响应体
func ReadResponseBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	
	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	return body, nil
}