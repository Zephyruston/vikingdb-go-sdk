// Copyright (c) 2025 Beijing Volcano Engine Technology Co., Ltd.
// SPDX-License-Identifier: Apache-2.0

package memory

import (
	"context"
	"net/http"
)

// Message represents a chat message.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// AddSessionRequest is the request for AddSession.
type AddSessionRequest struct {
	SessionID      string                   `json:"session_id"`
	Messages       []Message                `json:"messages"`
	Metadata       map[string]interface{}   `json:"metadata,omitempty"`
	Profiles       []interface{}            `json:"profiles,omitempty"`
	CollectionName string                   `json:"collection_name"`
	ProjectName    string                   `json:"project_name"`
	ResourceID     string                   `json:"resource_id,omitempty"`
}

// Collection represents a memory collection.
type Collection struct {
	client         *transport
	collectionName string
	projectName    string
	resourceID     string
}

// AddSession adds a session to the collection.
func (c *Collection) AddSession(ctx context.Context, sessionID string, messages []Message, opts ...CollectionOption) error {
	req := AddSessionRequest{
		SessionID:      sessionID,
		Messages:       messages,
		CollectionName: c.collectionName,
		ProjectName:    c.projectName,
		ResourceID:     c.resourceID,
	}

	for _, opt := range opts {
		opt(&req)
	}

	path := "/api/memory/session/add"
	return c.client.doRequest(ctx, http.MethodPost, path, req, nil)
}

// CollectionOption is an option for collection operations.
type CollectionOption func(*AddSessionRequest)

// WithMetadata sets the metadata.
func WithMetadata(metadata map[string]interface{}) CollectionOption {
	return func(r *AddSessionRequest) {
		r.Metadata = metadata
	}
}

// WithProfiles sets the profiles.
func WithProfiles(profiles []interface{}) CollectionOption {
	return func(r *AddSessionRequest) {
		r.Profiles = profiles
	}
}
