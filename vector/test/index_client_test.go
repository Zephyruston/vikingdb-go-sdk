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
func newTestIndexClient(t *testing.T) vector.IndexClient {
	// Replace with your actual test configuration
	accessKey := ""
	secretKey := ""
	endpoint := ""
	collectionName := "golang_test"
	indexName := "golang_test"

	indexConfig := model.DataAPIIndexBase{
		DataAPICollectionBase: model.DataAPICollectionBase{
			CollectionName: collectionName,
		},
		IndexName: indexName,
	}

	opts := []vector.ClientOption{
		vector.WithEndpoint(endpoint),
	}

	client, err := vector.NewIndexClientWithAkSk(accessKey, secretKey, indexConfig, opts...)
	assert.NoError(t, err)
	return client
}

func TestIndexClient_Fetch(t *testing.T) {
	client := newTestIndexClient(t)
	req := model.FetchDataInIndexRequest{
		IDs: []interface{}{"1"},
	}
	resp, err := client.Fetch(context.Background(), req)
	fmt.Println(resp.Result)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestIndexClient_SearchByVector(t *testing.T) {
	client := newTestIndexClient(t)
	limit := 10
	req := model.SearchByVectorRequest{
		SearchBase: model.SearchBase{
			Limit: &limit,
		},
		Dense: []float64{0.1, 0.2, 0.3, 0.4},
	}
	resp, err := client.SearchByVector(context.Background(), req)
	fmt.Println(resp.Result)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestIndexClient_SearchByID(t *testing.T) {
	client := newTestIndexClient(t)
	req := model.SearchByIDRequest{
		ID: "1",
	}
	resp, err := client.SearchByID(context.Background(), req)
	fmt.Println(resp.Result)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}