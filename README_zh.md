# Volc-VikingDB Golang SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/volcengine/vikingdb-go-sdk.svg)](https://pkg.go.dev/github.com/volcengine/vikingdb-go-sdk)

Volc-VikingDB Golang SDK 为与 Volc-VikingDB 服务交互提供了一套全面的工具。它旨在实现直观、灵活和高效，使开发人员能够轻松地将其应用程序与 Volc-VikingDB 集成。

## 特性

- **全面的 API 覆盖**：SDK 为 Volc-VikingDB 的所有功能提供了完整的接口，包括集合、索引和数据操作。
- **专用客户端**：SDK 为不同的功能模块（如 `CollectionClient`、`IndexClient` 和 `EmbeddingClient`）提供专用客户端，以提供更专注、更直观的开发体验。
- **灵活的配置**：SDK 支持客户端的灵活配置，包括端点、区域和凭据的自定义设置。
- **自动重试**：SDK 会自动重试因网络问题或服务器端瞬时错误而失败的请求，从而提高了应用程序的可靠性。
- **上下文感知**：SDK 具有上下文感知能力，可以更好地控制请求取消和超时。

## 安装

要安装 Volc-VikingDB Golang SDK，您可以使用 `go get` 命令：

```bash
go get github.com/volcengine/vikingdb-go-sdk
```

## 快速入门

以下是使用 Volc-VikingDB Golang SDK 的快速入门指南。

### 客户端初始化

首先，您需要使用您的凭据和所需的配置来初始化客户端。SDK 支持使用 Access Key 和 Secret Key 进行初始化。

```go
import (
    "github.com/volcengine/vikingdb-go-sdk/vector"
    "github.com/volcengine/vikingdb-go-sdk/vector/model"
)

// 使用 Access Key 和 Secret Key 初始化客户端
accessKey := "YOUR_AK"
secretKey := "YOUR_SK"
endpoint := "http://YOUR_ENDPOINT" // 例如 http://10.1.2.3:8080
collectionName := "your_collection_name"

collectionConfig := model.DataAPICollectionBase{
    CollectionName: collectionName,
}

opts := []vector.ClientOption{
    vector.WithEndpoint(endpoint),
}

client, err := vector.NewCollectionClientWithAkSk(accessKey, secretKey, collectionConfig, opts...)
if err != nil {
    // 处理错误
}
```

### 专用客户端

SDK 为不同的功能模块（如 `CollectionClient`、`IndexClient` 和 `EmbeddingClient`）提供专用客户端。您可以以类似的方式初始化这些客户端。

#### 集合客户端

`CollectionClient` 用于管理和操作特定集合中的数据。

```go
import (
    "github.com/volcengine/vikingdb-go-sdk/vector"
    "github.com/volcengine/vikingdb-go-sdk/vector/model"
)

accessKey := "YOUR_AK"
secretKey := "YOUR_SK"
endpoint := "http://YOUR_ENDPOINT"
collectionName := "your_collection_name"

collectionConfig := model.DataAPICollectionBase{
    CollectionName: collectionName,
}

opts := []vector.ClientOption{
    vector.WithEndpoint(endpoint),
}

collectionClient, err := vector.NewCollectionClientWithAkSk(accessKey, secretKey, collectionConfig, opts...)
if err != nil {
    // 处理错误
}
```

#### 索引客户端

`IndexClient` 用于管理和操作特定集合中的索引。

```go
import (
    "github.com/volcengine/vikingdb-go-sdk/vector"
    "github.com/volcengine/vikingdb-go-sdk/vector/model"
)

accessKey := "YOUR_AK"
secretKey := "YOUR_SK"
endpoint := "http://YOUR_ENDPOINT"
collectionName := "your_collection_name"
indexName := "your_index_name"

indexConfig := model.DataAPIIndexBase{
    DataAPICollectionBase: model.DataAPICollectionBase{
        CollectionName: collectionName,
    },
    IndexName:      indexName,
}

opts := []vector.ClientOption{
    vector.WithEndpoint(endpoint),
}

indexClient, err := vector.NewIndexClientWithAkSk(accessKey, secretKey, indexConfig, opts...)
if err != nil {
    // 处理错误
}
```

#### 嵌入客户端

`EmbeddingClient` 用于执行文本嵌入操作。

```go
import (
    "github.com/volcengine/vikingdb-go-sdk/vector"
)

accessKey := "YOUR_AK"
secretKey := "YOUR_SK"
endpoint := "http://YOUR_ENDPOINT"

opts := []vector.ClientOption{
    vector.WithEndpoint(endpoint),
}

embeddingClient, err := vector.NewEmbeddingClientWithAkSk(accessKey, secretKey, opts...)
if err != nil {
    // 处理错误
}
```

### 数据操作

SDK 提供了丰富的数据操作接口，包括 `Upsert`、`Update`、`Delete`、`Fetch` 和 `Query`。

#### Upsert 数据

`Upsert` 操作用于插入新数据或更新现有数据。

```go
import (
    "context"
    "github.com/volcengine/vikingdb-go-sdk/vector/model"
)

req := &model.UpsertDataRequest{
    WriteDataBase: model.WriteDataBase{
        Data: []model.MapStr{
            {
                "ID": 1,
                "vector": []float64{1.1,2.2,3.4,4.2},
            },
        },
    },
}

resp, err := collectionClient.Upsert(context.Background(), req)
if err != nil {
    // 处理错误
}
// 处理响应
```

#### 向量检索

`Query` 操作用于执行向量检索。（注意：测试文件中不包含 `Query` 示例。以下为通用结构。）

```go
import (
    "context"
    "github.com/volcengine/vikingdb-go-sdk/vector/model"
)

req := &model.QueryDataRequest{
    // ... 指定查询参数
}

resp, err := collectionClient.Query(context.Background(), req)
if err != nil {
    // 处理错误
}
// 处理响应
```

### 文本嵌入

`Embedding` 操作用于将文本转换为向量。

```go
import (
    "context"
    "github.com/volcengine/vikingdb-go-sdk/vector/model"
)

text := "hello world"
model_name := "doubao-embedding"
model_version := "240715"
req := &model.EmbeddingRequest{
    DenseModel: &model.EmbeddingModelOpt{
        ModelName: &model_name,
        ModelVersion: &model_version,
    },
    Data: []*model.EmbeddingData{
        {
            Text: &text,
        },
    },
}

resp, err := embeddingClient.Embedding(context.Background(), req)
if err != nil {
    // 处理错误
}
// 处理响应
```

## API 参考

有关详细的 API 参考，请访问 [Go Reference](https://pkg.go.dev/github.com/volcengine/vikingdb-go-sdk)。