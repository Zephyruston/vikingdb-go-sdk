package vector

// 认证类型
const (
	// AuthTypeAkSk 使用 AK/SK 认证
	AuthTypeAkSk = "ak_sk"
	
	// AuthTypeApiKey 使用 API Key 认证
	AuthTypeApiKey = "api_key"
)

// SDK 版本
const (
	// Version SDK 版本号
	Version = "0.1.0"
)

// Config 表示客户端配置
type Config struct {
	// 服务端点
	Endpoint string
	
	// 区域
	Region string
	
	// 认证类型
	AuthType string
	
	// Access Key
	AccessKey string
	
	// Secret Key
	SecretKey string
	
	// API Key
	ApiKey string
	
	// 超时时间（毫秒）
	Timeout int
	
	// 最大重试次数
	MaxRetries int
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		Endpoint:   "https://api.vector.bytedance.com",
		Region:     "cn-beijing",
		AuthType:   AuthTypeAkSk,
		Timeout:    30000,  // 30 秒
		MaxRetries: 3,
	}
}

// ClientOption 表示客户端选项函数
type ClientOption func(*Config)

// WithEndpoint 设置服务端点
func WithEndpoint(endpoint string) ClientOption {
	return func(c *Config) {
		c.Endpoint = endpoint
	}
}

// WithRegion 设置区域
func WithRegion(region string) ClientOption {
	return func(c *Config) {
		c.Region = region
	}
}

// WithTimeout 设置超时时间（毫秒）
func WithTimeout(timeout int) ClientOption {
	return func(c *Config) {
		c.Timeout = timeout
	}
}

// WithMaxRetries 设置最大重试次数
func WithMaxRetries(maxRetries int) ClientOption {
	return func(c *Config) {
		c.MaxRetries = maxRetries
	}
}