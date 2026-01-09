// Copyright (c) 2025 Beijing Volcano Engine Technology Co., Ltd.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/volcengine/vikingdb-go-sdk/memory"
)

func main() {
	apiKey := os.Getenv("MEMORY_API_KEY")
	if apiKey == "" {
		apiKey = "your_key"
	}

	client, err := memory.New(
		memory.AuthAPIKey(apiKey),
		memory.WithEndpoint("http://api-knowledgebase.mlp.cn-beijing.volces.com"),
	)
	if err != nil {
		fmt.Printf("Failed to create client: %v\n", err)
		return
	}

	collection := client.Collection("my_first_memory_collection", "default")

	ctx := context.Background()
	nowTs := time.Now().UnixMilli()

	err = collection.AddSession(
		ctx,
		"session_001",
		[]memory.Message{
			{Role: "user", Content: "今天天气怎么样？"},
			{Role: "assistant", Content: "今天天气晴朗，气温22度，非常适合外出。"},
		},
		memory.WithMetadata(map[string]interface{}{
			"default_user_id":      "user_01",
			"default_user_name":    "XiaoMing",
			"default_assistant_id": "assistant_01",
			"default_assistant_name": "Robot",
			"time": nowTs,
		}),
	)
	if err != nil {
		fmt.Printf("Failed to add session: %v\n", err)
		return
	}

	fmt.Println("AddSession succeeded!")
}
