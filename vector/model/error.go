package model

import (
	"fmt"
	"net/http"
)

// ErrorCode 定义错误码
type ErrorCode string

// 错误码常量
const (
	// HTTP 报错
	ErrCodeHTTPRequestFailed ErrorCode = "HTTPRequestFailed"

	// 通用错误
	ErrCodeUnknown           ErrorCode = "Unknown"
	ErrCodeInvalidParameter  ErrorCode = "InvalidParameter"
	ErrCodeServiceUnavailable ErrorCode = "ServiceUnavailable"
	ErrCodeTimeout           ErrorCode = "Timeout"
	ErrCodeRequestLimitExceeded ErrorCode = "RequestLimitExceeded"
	ErrCodeUnauthorized      ErrorCode = "Unauthorized"
	ErrCodeForbidden         ErrorCode = "Forbidden"
	ErrCodeNotFound          ErrorCode = "NotFound"
	
	// 集合相关错误
	ErrCodeCollectionNotExists ErrorCode = "CollectionNotExists"
	ErrCodeCollectionAlreadyExists ErrorCode = "CollectionAlreadyExists"
	ErrCodeCollectionCreateFailed ErrorCode = "CollectionCreateFailed"
	ErrCodeCollectionUpdateFailed ErrorCode = "CollectionUpdateFailed"
	ErrCodeCollectionDeleteFailed ErrorCode = "CollectionDeleteFailed"
	
	// 数据相关错误
	ErrCodeDataInsertFailed   ErrorCode = "DataInsertFailed"
	ErrCodeDataUpdateFailed   ErrorCode = "DataUpdateFailed"
	ErrCodeDataDeleteFailed   ErrorCode = "DataDeleteFailed"
	ErrCodeDataNotFound       ErrorCode = "DataNotFound"
	
	// 检索相关错误
	ErrCodeSearchFailed       ErrorCode = "SearchFailed"
	ErrCodeIndexNotExists     ErrorCode = "IndexNotExists"
	
	// 嵌入相关错误
	ErrCodeEmbeddingFailed    ErrorCode = "EmbeddingFailed"
	ErrCodeModelNotFound      ErrorCode = "ModelNotFound"
)

// Error 表示 SDK 错误
type Error struct {
	// 错误码
	Code ErrorCode `json:"code"`
	
	// 错误消息
	Message string `json:"message"`
	
	// HTTP 状态码
	StatusCode int `json:"status_code,omitempty"`
	
	// 请求 ID
	RequestID string `json:"request_id,omitempty"`
	
	// 原始错误
	Err error `json:"-"`
}

// Error 实现 error 接口
func (e *Error) Error() string {
	if e.RequestID != "" {
		return fmt.Sprintf("vikingdb error: code=%s, message=%s, status_code=%d, err=%v, request_id=%s", e.Code, e.Message, e.StatusCode, e.Err, e.RequestID)
	}
	return fmt.Sprintf("vikingdb error: code=%s, message=%s, status_code=%d, err=%v", e.Code, e.Message, e.StatusCode, e.Err)
}

// Unwrap 返回原始错误
func (e *Error) Unwrap() error {
	return e.Err
}

// NewError 创建一个新的错误
func NewError(code ErrorCode, message string) *Error {
	return &Error{
		Code:       code,
		Message:    message,
		StatusCode: http.StatusInternalServerError,
	}
}

// NewErrorWithStatusCode 创建一个带有状态码的新错误
func NewErrorWithStatusCode(code ErrorCode, message string, statusCode int) *Error {
	return &Error{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
	}
}

// NewErrorWithRequestID 创建一个带有请求 ID 的新错误
func NewErrorWithRequestID(code ErrorCode, message string, requestID string, statusCode int) *Error {
	return &Error{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
		RequestID:  requestID,
	}
}

// NewErrorWithCause 创建一个带有原因的新错误
func NewErrorWithCause(code ErrorCode, message string, cause error, statusCode int) *Error {
	return &Error{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
		Err:        cause,
	}
}

// IsRetryableError 判断错误是否可重试
func IsRetryableError(err error) bool {
	if err == nil {
		return false
	}
	
	sdkErr, ok := err.(*Error)
	if !ok {
		return false
	}
	
	switch sdkErr.StatusCode {
	case http.StatusTooManyRequests, http.StatusServiceUnavailable, http.StatusGatewayTimeout:
		return true
	}
	
	switch sdkErr.Code {
	case ErrCodeServiceUnavailable, ErrCodeTimeout, ErrCodeRequestLimitExceeded:
		return true
	}
	
	return false
}

// NewInvalidParameterError 创建一个参数无效错误
func NewInvalidParameterError(message string) *Error {
	return NewErrorWithStatusCode(ErrCodeInvalidParameter, message, http.StatusBadRequest)
}

// NewUnauthorizedError 创建一个未授权错误
func NewUnauthorizedError(message string) *Error {
	return NewErrorWithStatusCode(ErrCodeUnauthorized, message, http.StatusUnauthorized)
}

// NewForbiddenError 创建一个禁止访问错误
func NewForbiddenError(message string) *Error {
	return NewErrorWithStatusCode(ErrCodeForbidden, message, http.StatusForbidden)
}

// NewNotFoundError 创建一个资源不存在错误
func NewNotFoundError(message string) *Error {
	return NewErrorWithStatusCode(ErrCodeNotFound, message, http.StatusNotFound)
}

// NewServiceUnavailableError 创建一个服务不可用错误
func NewServiceUnavailableError(message string) *Error {
	return NewErrorWithStatusCode(ErrCodeServiceUnavailable, message, http.StatusServiceUnavailable)
}

// NewTimeoutError 创建一个超时错误
func NewTimeoutError(message string) *Error {
	return NewErrorWithStatusCode(ErrCodeTimeout, message, http.StatusGatewayTimeout)
}

// NewRequestLimitExceededError 创建一个请求限制超出错误
func NewRequestLimitExceededError(message string) *Error {
	return NewErrorWithStatusCode(ErrCodeRequestLimitExceeded, message, http.StatusTooManyRequests)
}