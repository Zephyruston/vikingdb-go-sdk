// Copyright (c) 2025 Beijing Volcano Engine Technology Co., Ltd.
// SPDX-License-Identifier: Apache-2.0

package memory

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Version denotes the SDK version.
const Version = "0.1.0"

type authKind int

const (
	authKindNone authKind = iota
	authKindIAM
	authKindAPIKey
)

type Auth struct {
	kind      authKind
	accessKey string
	secretKey string
	apiKey    string
}

func AuthNone() Auth {
	return Auth{kind: authKindNone}
}

func AuthIAM(accessKey, secretKey string) Auth {
	return Auth{kind: authKindIAM, accessKey: accessKey, secretKey: secretKey}
}

func AuthAPIKey(apiKey string) Auth {
	return Auth{kind: authKindAPIKey, apiKey: apiKey}
}

type Config struct {
	Endpoint   string
	Region     string
	Timeout    time.Duration
	MaxRetries int
	HTTPClient *http.Client
	UserAgent  string
}

func DefaultConfig() Config {
	return Config{
		Endpoint:   "http://api-knowledgebase.mlp.cn-beijing.volces.com",
		Region:     "cn-beijing",
		Timeout:    30 * time.Second,
		MaxRetries: 3,
	}
}

type transport struct {
	config     Config
	httpClient *http.Client
	baseURL    *url.URL
	auth       authenticator
	userAgent  string
}

type authenticator interface {
	apply(req *http.Request) (*http.Request, error)
}

type noAuth struct{}

func (noAuth) apply(req *http.Request) (*http.Request, error) {
	return req, nil
}

type apiKeyAuth struct {
	token string
}

func (a apiKeyAuth) apply(req *http.Request) (*http.Request, error) {
	if a.token != "" {
		req.Header.Set("Authorization", "Bearer "+a.token)
	}
	return req, nil
}

type iamAuth struct {
	ak     string
	sk     string
	region string
}

func (a iamAuth) apply(req *http.Request) (*http.Request, error) {
	req.Header.Set("X-Date", time.Now().UTC().Format(http.TimeFormat))
	return req, nil
}

func newTransport(cfg Config, authConfig Auth) (*transport, error) {
	if cfg.Endpoint == "" {
		return nil, fmt.Errorf("endpoint cannot be empty")
	}

	baseURL, err := url.Parse(cfg.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid endpoint: %w", err)
	}
	if baseURL.Scheme == "" {
		baseURL.Scheme = "http"
	}

	if cfg.Timeout <= 0 {
		cfg.Timeout = DefaultConfig().Timeout
	}

	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{Timeout: cfg.Timeout}
	}

	userAgent := cfg.UserAgent
	if userAgent == "" {
		userAgent = fmt.Sprintf("vikingdb-go-sdk/memory/%s", Version)
	}

	var auth authenticator = noAuth{}
	switch authConfig.kind {
	case authKindIAM:
		auth = iamAuth{ak: authConfig.accessKey, sk: authConfig.secretKey, region: cfg.Region}
	case authKindAPIKey:
		auth = apiKeyAuth{token: authConfig.apiKey}
	}

	return &transport{
		config:     cfg,
		httpClient: httpClient,
		baseURL:    baseURL,
		auth:       auth,
		userAgent:  userAgent,
	}, nil
}

func (c *transport) doRequest(ctx context.Context, method, path string, request, response interface{}) error {
	if ctx == nil {
		ctx = context.Background()
	}

	var body []byte
	if request != nil {
		serialized, err := json.Marshal(request)
		if err != nil {
			return fmt.Errorf("failed to marshal request: %w", err)
		}
		body = serialized
	}

	targetURL := c.baseURL.ResolveReference(&url.URL{Path: path})
	var buf *bytes.Reader
	if len(body) > 0 {
		buf = bytes.NewReader(body)
	}
	req, err := http.NewRequestWithContext(ctx, method, targetURL.String(), buf)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	if len(body) > 0 {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("User-Agent", c.userAgent)

	signedReq, err := c.auth.apply(req)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(signedReq)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("request failed with status: %d", resp.StatusCode)
	}

	if response != nil {
		return json.NewDecoder(resp.Body).Decode(response)
	}
	return nil
}

// Client represents the Viking Memory client.
type Client struct {
	transport *transport
}

// New creates a new Viking Memory client.
func New(auth Auth, opts ...func(*Config)) (*Client, error) {
	cfg := DefaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	transport, err := newTransport(cfg, auth)
	if err != nil {
		return nil, err
	}

	return &Client{transport: transport}, nil
}

// Collection returns a Collection client for the given collection.
func (c *Client) Collection(collectionName, projectName string) *Collection {
	return &Collection{
		client:         c.transport,
		collectionName: collectionName,
		projectName:    projectName,
	}
}

// WithEndpoint sets the endpoint.
func WithEndpoint(endpoint string) func(*Config) {
	return func(c *Config) {
		c.Endpoint = endpoint
	}
}

// WithRegion sets the region.
func WithRegion(region string) func(*Config) {
	return func(c *Config) {
		c.Region = region
	}
}

// WithTimeout sets the timeout.
func WithTimeout(timeout time.Duration) func(*Config) {
	return func(c *Config) {
		c.Timeout = timeout
	}
}
