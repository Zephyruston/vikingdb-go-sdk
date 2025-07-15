# Volc-VikingDB Golang SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/volcengine/vikingdb-go-sdk.svg)](https://pkg.go.dev/github.com/volcengine/vikingdb-go-sdk)

The Volc-VikingDB Golang SDK provides a comprehensive suite of tools for interacting with the Volc-VikingDB service. It is designed to be intuitive, flexible, and efficient, enabling developers to easily integrate their applications with Volc-VikingDB.

## Features

- **Comprehensive API Coverage**: The SDK provides a complete interface to all the features of Volc-VikingDB, including collections, indexes, and data manipulation.
- **Specialized Clients**: The SDK offers specialized clients for different functional modules, such as `CollectionClient`, `IndexClient`, and `EmbeddingClient`, to provide a more focused and intuitive development experience.
- **Flexible Configuration**: The SDK supports flexible configuration of clients, including custom settings for endpoints, regions, and credentials.
- **Automatic Retries**: The SDK automatically retries failed requests due to network issues or server-side transient errors, improving the reliability of applications.
- **Context-Awareness**: The SDK is context-aware, allowing for better control over request cancellation and timeouts.

## Installation

To install the Volc-VikingDB Golang SDK, you can use the `go get` command:

```bash
go get github.com/volcengine/vikingdb-go-sdk
```

## Quick Start

The following is a quick start guide to using the Volc-VikingDB Golang SDK.

### Client Initialization

To get started, you need to initialize a client with your credentials and the desired configuration. The SDK supports initialization with an Access Key and Secret Key.

```go
import (
    "github.com/volcengine/vikingdb-go-sdk/vector"
    "github.com/volcengine/vikingdb-go-sdk/vector/model"
)

// Initialize the client with an Access Key and Secret Key
accessKey := "YOUR_AK"
secretKey := "YOUR_SK"
endpoint := "http://YOUR_ENDPOINT" // e.g. http://10.1.2.3:8080
collectionName := "your_collection_name"

collectionConfig := model.DataAPICollectionBase{
    CollectionName: collectionName,
}

opts := []vector.ClientOption{
    vector.WithEndpoint(endpoint),
}

client, err := vector.NewCollectionClientWithAkSk(accessKey, secretKey, collectionConfig, opts...)
if err != nil {
    // Handle error
}
```

### Specialized Clients

The SDK provides specialized clients for different functional modules, such as `CollectionClient`, `IndexClient`, and `EmbeddingClient`. You can initialize these clients in a similar way.

#### Collection Client

The `CollectionClient` is used to manage and operate on data within a specific collection.

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
    // Handle error
}
```

#### Index Client

The `IndexClient` is used to manage and operate on indexes within a specific collection.

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
    // Handle error
}
```

#### Embedding Client

The `EmbeddingClient` is used to perform text embedding operations.

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
    // Handle error
}
```

### Data Operations

The SDK provides a rich set of data manipulation interfaces, including `Upsert`, `Update`, `Delete`, `Fetch`, and `Query`.

#### Upsert Data

The `Upsert` operation is used to insert new data or update existing data.

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
    // Handle error
}
// Process response
```

#### Vector Retrieval

The `Query` operation is used to perform vector retrieval. (Note: The test files do not contain a `Query` example. The following is a general structure.)

```go
import (
    "context"
    "github.com/volcengine/vikingdb-go-sdk/vector/model"
)

req := &model.QueryDataRequest{
    // ... specify query parameters
}

resp, err := collectionClient.Query(context.Background(), req)
if err != nil {
    // Handle error
}
// Process response
```

### Text Embedding

The `Embedding` operation is used to convert text into vectors.

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
    // Handle error
}
// Process response
```

## API Reference

For a detailed API reference, please visit the [Go Reference](https://pkg.go.dev/github.com/volcengine/vikingdb-go-sdk).