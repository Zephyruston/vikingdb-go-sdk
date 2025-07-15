package vector

import (
	"context"
	"net/http"

	"github.com/volcengine/vikingdb-go-sdk/vector/model"
)

// collectionClient 是 CollectionClient 接口的实现
type collectionClient struct {
	client         *client
	collectionBase model.DataAPICollectionBase
}

// Upsert 插入或更新数据
func (c *collectionClient) Upsert(ctx context.Context, request model.UpsertDataRequest, opts ...RequestOption) (*model.UpsertDataResponse, error) {
	response := &model.UpsertDataResponse{}
	req := struct {
		model.DataAPICollectionBase
		model.UpsertDataRequest
	}{
		DataAPICollectionBase: c.collectionBase,
		UpsertDataRequest:   request,
	}
	err := c.client.doRequest(ctx, http.MethodPost, "/api/vikingdb/data/upsert", req, response, opts...)
	return response, err
}

// Update 更新数据
func (c *collectionClient) Update(ctx context.Context, request model.UpdateDataRequest, opts ...RequestOption) (*model.UpdateDataResponse, error) {
	response := &model.UpdateDataResponse{}
	req := struct {
		model.DataAPICollectionBase
		model.UpdateDataRequest
	}{
		DataAPICollectionBase: c.collectionBase,
		UpdateDataRequest:   request,
	}
	err := c.client.doRequest(ctx, http.MethodPost, "/api/vikingdb/data/update", req, response, opts...)
	return response, err
}

// Delete 删除数据
func (c *collectionClient) Delete(ctx context.Context, request model.DeleteDataRequest, opts ...RequestOption) (*model.DeleteDataResponse, error) {
	response := &model.DeleteDataResponse{}
	req := struct {
		model.DataAPICollectionBase
		model.DeleteDataRequest
	}{
		DataAPICollectionBase: c.collectionBase,
		DeleteDataRequest:   request,
	}
	err := c.client.doRequest(ctx, http.MethodPost, "/api/vikingdb/data/delete", req, response, opts...)
	return response, err
}

// Fetch 获取数据
func (c *collectionClient) Fetch(ctx context.Context, request model.FetchDataInCollectionRequest, opts ...RequestOption) (*model.FetchDataInCollectionResponse, error) {
	response := &model.FetchDataInCollectionResponse{}
	req := struct {
		model.DataAPICollectionBase
		model.FetchDataInCollectionRequest
	}{
		DataAPICollectionBase:      c.collectionBase,
		FetchDataInCollectionRequest: request,
	}
	err := c.client.doRequest(ctx, http.MethodPost, "/api/vikingdb/data/fetch_in_collection", req, response, opts...)
	return response, err
}

// CollectionName 返回集合名称
func (c *collectionClient) CollectionName() string {
	return c.collectionBase.CollectionName
}

func (c *collectionClient) ProjectName() string {
	return c.collectionBase.ProjectName
}

func (c *collectionClient) ResourceID() string {
	return c.collectionBase.ResourceID
}