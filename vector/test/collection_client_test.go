// Copyright (c) 2025 Beijing Volcano Engine Technology Co., Ltd.
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/volcengine/vikingdb-go-sdk/vector"
	"github.com/volcengine/vikingdb-go-sdk/vector/model"
)

// Mock a client for testing
func newTestCollectionClient(t *testing.T) vector.CollectionClient {
	// Replace with your actual test configuration
	accessKey := ""
	secretKey := ""
	endpoint := ""
	collectionName := "golang_test"

	opts := []vector.ClientOption{
		vector.WithEndpoint(endpoint),
	}

	collectionConfig := model.DataAPICollectionBase{
		CollectionName: collectionName,
	}

	client, err := vector.NewCollectionClientWithAkSk(accessKey, secretKey, collectionConfig, opts...)
	assert.NoError(t, err)
	return client
}

func TestCollectionClient_Upsert(t *testing.T) {
	client := newTestCollectionClient(t)
	req := model.UpsertDataRequest{
		WriteDataBase: model.WriteDataBase{
			Data: []model.MapStr{
				{
					"ID":     1,
					"vector": []float64{1.1, 2.2, 3.4, 4.2},
				},
			},
		},
	}
	resp, err := client.Upsert(context.Background(), req)
	fmt.Println(resp.Message)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestCollectionClient_Update(t *testing.T) {
	client := newTestCollectionClient(t)
	req := model.UpdateDataRequest{
		WriteDataBase: model.WriteDataBase{
			Data: []model.MapStr{
				{
					"ID":     1,
					"vector": []float64{1.1, 2.2, 3.4, 4.2},
				},
			},
		},
	}
	resp, err := client.Update(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestCollectionClient_Delete(t *testing.T) {
	client := newTestCollectionClient(t)
	req := model.DeleteDataRequest{
		IDs: []interface{}{"1"},
	}
	resp, err := client.Delete(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestCollectionClient_Fetch(t *testing.T) {
	client := newTestCollectionClient(t)
	req := model.FetchDataInCollectionRequest{
		IDs: []interface{}{"1"},
	}
	resp, err := client.Fetch(context.Background(), req)
	fmt.Println(resp)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
