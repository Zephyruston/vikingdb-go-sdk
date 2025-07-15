package utils

import (
	"math"
	"math/rand"
	"time"

	"github.com/volcengine/vikingdb-go-sdk/vector/model"
)

// Retry 重试函数
func Retry(maxRetries int, fn func() error, shouldRetry func(error) bool) error {
	var err error
	
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())
	
	// 尝试执行函数
	for i := 0; i <= maxRetries; i++ {
		// 执行函数
		err = fn()
		
		// 如果没有错误或者不应该重试，直接返回
		if err == nil || (i == maxRetries) || (shouldRetry != nil && !shouldRetry(err)) {
			return err
		}
		
		// 计算退避时间
		backoffTime := calculateBackoffTime(i)
		
		// 等待退避时间
		time.Sleep(backoffTime)
	}
	
	return err
}

// IsRetryableError 判断错误是否可重试
func IsRetryableError(err error) bool {
	if err == nil {
		return false
	}
	
	// 尝试将错误转换为 SDK 错误
	sdkErr, ok := err.(*model.Error)
	if !ok {
		return false
	}
	
	// 根据状态码判断是否可重试
	switch sdkErr.StatusCode {
	case 429, 500, 502, 503, 504:
		return true
	}
	
	// 根据错误码判断是否可重试
	switch sdkErr.Code {
	case model.ErrCodeServiceUnavailable, model.ErrCodeTimeout, model.ErrCodeRequestLimitExceeded:
		return true
	}
	
	return false
}

// calculateBackoffTime 计算退避时间
func calculateBackoffTime(retryCount int) time.Duration {
	// 基础退避时间（毫秒）
	baseBackoff := 100
	
	// 计算退避时间
	backoff := baseBackoff * int(math.Pow(2, float64(retryCount)))
	
	// 添加随机抖动
	jitter := rand.Intn(backoff / 2)
	backoff = backoff + jitter
	
	// 最大退避时间（10 秒）
	maxBackoff := 10000
	if backoff > maxBackoff {
		backoff = maxBackoff
	}
	
	return time.Duration(backoff) * time.Millisecond
}