package vector

import (
	"context"

	"github.com/volcengine/vikingdb-go-sdk/vector/model"
)

// CollectionClient 定义了集合相关操作的runtime接口
type CollectionClient interface {
	// 插入或更新数据
	Upsert(ctx context.Context, request model.UpsertDataRequest, opts ...RequestOption) (*model.UpsertDataResponse, error)
	
	// 更新数据
	Update(ctx context.Context, request model.UpdateDataRequest, opts ...RequestOption) (*model.UpdateDataResponse, error)
	
	// 删除数据
	Delete(ctx context.Context, request model.DeleteDataRequest, opts ...RequestOption) (*model.DeleteDataResponse, error)
	
	// 获取数据
	Fetch(ctx context.Context, request model.FetchDataInCollectionRequest, opts ...RequestOption) (*model.FetchDataInCollectionResponse, error)
	
	// 获取集合名称
	CollectionName() string

	// 获取资源ID
	ResourceID() string

	// 获取项目名称
	ProjectName() string
}

// IndexClient 定义了索引相关操作的runtime接口
type IndexClient interface {
	// 获取索引中的数据
	Fetch(ctx context.Context, request model.FetchDataInIndexRequest, opts ...RequestOption) (*model.FetchDataInIndexResponse, error)
	
	// 向量检索
	SearchByVector(ctx context.Context, request model.SearchByVectorRequest, opts ...RequestOption) (*model.SearchResponse, error)
	
	// 多模态检索
	SearchByMultiModal(ctx context.Context, request model.SearchByMultiModalRequest, opts ...RequestOption) (*model.SearchResponse, error)
	
	// ID 检索
	SearchByID(ctx context.Context, request model.SearchByIDRequest, opts ...RequestOption) (*model.SearchResponse, error)
	
	// 标量检索
	SearchByScalar(ctx context.Context, request model.SearchByScalarRequest, opts ...RequestOption) (*model.SearchResponse, error)
	
	// 关键词检索
	SearchByKeywords(ctx context.Context, request model.SearchByKeywordsRequest, opts ...RequestOption) (*model.SearchResponse, error)
	
	// 随机检索
	SearchByRandom(ctx context.Context, request model.SearchByRandomRequest, opts ...RequestOption) (*model.SearchResponse, error)
	
	// 聚合
	Aggregate(ctx context.Context, request model.AggRequest, opts ...RequestOption) (*model.AggResponse, error)
	
	// 排序
	Sort(ctx context.Context, request model.SortRequest, opts ...RequestOption) (*model.SortResponse, error)
	
	// 获取集合名称
	CollectionName() string

	// 获取索引名称
	IndexName() string

	// 获取资源ID
	ResourceID() string

	// 获取项目名称
	ProjectName() string
}

// EmbeddingClient 定义了嵌入相关操作的runtime接口
type EmbeddingClient interface {
	Embedding(ctx context.Context, request model.EmbeddingRequest, opts ...RequestOption) (*model.EmbeddingResponse, error)
}