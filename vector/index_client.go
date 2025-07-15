package vector

import (
	"context"
	"net/http"

	"github.com/volcengine/vikingdb-go-sdk/vector/model"
)

// indexClient 是 IndexClient 接口的实现
type indexClient struct {
	client    *client
	indexBase model.DataAPIIndexBase
}

// Fetch 获取索引中的数据
func (i *indexClient) Fetch(ctx context.Context, request model.FetchDataInIndexRequest, opts ...RequestOption) (*model.FetchDataInIndexResponse, error) {
	response := &model.FetchDataInIndexResponse{}
	req := struct {
		model.DataAPIIndexBase
		model.FetchDataInIndexRequest
	}{
		DataAPIIndexBase:        i.indexBase,
		FetchDataInIndexRequest: request,
	}
	err := i.client.doRequest(ctx, http.MethodPost, "/api/vikingdb/data/fetch_in_index", req, response, opts...)
	return response, err
}

// SearchByVector 向量检索
func (i *indexClient) SearchByVector(ctx context.Context, request model.SearchByVectorRequest, opts ...RequestOption) (*model.SearchResponse, error) {
	response := &model.SearchResponse{}
	req := struct {
		model.DataAPIIndexBase
		model.SearchByVectorRequest
	}{
		DataAPIIndexBase:      i.indexBase,
		SearchByVectorRequest: request,
	}
	err := i.client.doRequest(ctx, http.MethodPost, "/api/vikingdb/data/search/vector", req, response, opts...)
	return response, err
}

// SearchByMultiModal 多模态检索
func (i *indexClient) SearchByMultiModal(ctx context.Context, request model.SearchByMultiModalRequest, opts ...RequestOption) (*model.SearchResponse, error) {
	response := &model.SearchResponse{}
	req := struct {
		model.DataAPIIndexBase
		model.SearchByMultiModalRequest
	}{
		DataAPIIndexBase:          i.indexBase,
		SearchByMultiModalRequest: request,
	}
	err := i.client.doRequest(ctx, http.MethodPost, "/api/vikingdb/data/search/multi_modal", req, response, opts...)
	return response, err
}

// SearchByID ID 检索
func (i *indexClient) SearchByID(ctx context.Context, request model.SearchByIDRequest, opts ...RequestOption) (*model.SearchResponse, error) {
	response := &model.SearchResponse{}
	req := struct {
		model.DataAPIIndexBase
		model.SearchByIDRequest
	}{
		DataAPIIndexBase:    i.indexBase,
		SearchByIDRequest: request,
	}
	err := i.client.doRequest(ctx, http.MethodPost, "/api/vikingdb/data/search/id", req, response, opts...)
	return response, err
}

// SearchByScalar 标量检索
func (i *indexClient) SearchByScalar(ctx context.Context, request model.SearchByScalarRequest, opts ...RequestOption) (*model.SearchResponse, error) {
	response := &model.SearchResponse{}
	req := struct {
		model.DataAPIIndexBase
		model.SearchByScalarRequest
	}{
		DataAPIIndexBase:        i.indexBase,
		SearchByScalarRequest: request,
	}
	err := i.client.doRequest(ctx, http.MethodPost, "/api/vikingdb/data/search/scalar", req, response, opts...)
	return response, err
}

// SearchByKeywords 关键词检索
func (i *indexClient) SearchByKeywords(ctx context.Context, request model.SearchByKeywordsRequest, opts ...RequestOption) (*model.SearchResponse, error) {
	response := &model.SearchResponse{}
	req := struct {
		model.DataAPIIndexBase
		model.SearchByKeywordsRequest
	}{
		DataAPIIndexBase:          i.indexBase,
		SearchByKeywordsRequest: request,
	}
	err := i.client.doRequest(ctx, http.MethodPost, "/api/vikingdb/data/search/keywords", req, response, opts...)
	return response, err
}

// SearchByRandom 随机检索
func (i *indexClient) SearchByRandom(ctx context.Context, request model.SearchByRandomRequest, opts ...RequestOption) (*model.SearchResponse, error) {
	response := &model.SearchResponse{}
	req := struct {
		model.DataAPIIndexBase
		model.SearchByRandomRequest
	}{
		DataAPIIndexBase:        i.indexBase,
		SearchByRandomRequest: request,
	}
	err := i.client.doRequest(ctx, http.MethodPost, "/api/vikingdb/data/search/random", req, response, opts...)
	return response, err
}

// Aggregate 聚合
func (i *indexClient) Aggregate(ctx context.Context, request model.AggRequest, opts ...RequestOption) (*model.AggResponse, error) {
	response := &model.AggResponse{}
	req := struct {
		model.DataAPIIndexBase
		model.AggRequest
	}{
		DataAPIIndexBase: i.indexBase,
		AggRequest:     request,
	}
	err := i.client.doRequest(ctx, http.MethodPost, "/api/vikingdb/data/agg", req, response, opts...)
	return response, err
}

// Sort 排序
func (i *indexClient) Sort(ctx context.Context, request model.SortRequest, opts ...RequestOption) (*model.SortResponse, error) {
	response := &model.SortResponse{}
	req := struct {
		model.DataAPIIndexBase
		model.SortRequest
	}{
		DataAPIIndexBase: i.indexBase,
		SortRequest:    request,
	}
	err := i.client.doRequest(ctx, http.MethodPost, "/api/vikingdb/data/sort", req, response, opts...)
	return response, err
}

// CollectionName 返回集合名称
func (i *indexClient) CollectionName() string {
	return i.indexBase.CollectionName
}

// IndexName 返回索引名称
func (i *indexClient) IndexName() string {
	return i.indexBase.IndexName
}

func (i *indexClient) ResourceID() string {
	return i.indexBase.ResourceID
}

func (i *indexClient) ProjectName() string {
	return i.indexBase.ProjectName
}