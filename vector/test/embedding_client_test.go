package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/volcengine/vikingdb-go-sdk/vector"
	"github.com/volcengine/vikingdb-go-sdk/vector/model"

	"github.com/stretchr/testify/assert"
)

// Mock a client for testing
func newTestEmbeddingClient(t *testing.T) vector.EmbeddingClient {
	// Replace with your actual test configuration
	accessKey := ""
	secretKey := ""
	endpoint := ""

	opts := []vector.ClientOption{
		vector.WithEndpoint(endpoint),
	}

	client, err := vector.NewEmbeddingClientWithAkSk(accessKey, secretKey, opts...)
	assert.NoError(t, err)
	return client
}

func TestEmbeddingClient_Embedding(t *testing.T) {
	client := newTestEmbeddingClient(t)
	text := "hello world"
	model_name := "doubao-embedding"
	model_version := "240715"
	// dim := 512
	req := model.EmbeddingRequest{
		DenseModel: &model.EmbeddingModelOpt{
			ModelName: &model_name,
			ModelVersion: &model_version,
			// Dim: &dim,
		},
		Data: []*model.EmbeddingData{
			{
				Text: &text,
			},
		},
	}
	resp, err := client.Embedding(context.Background(), req)
	fmt.Println(resp.RequestID)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}