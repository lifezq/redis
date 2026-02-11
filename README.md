[![Build status](https://img.shields.io/appveyor/build/lifezq/redis.svg)](https://ci.appveyor.com/project/lifezq/redis)
[![Coverage Status](https://img.shields.io/coveralls/lifezq/redis.svg?style=flat-square)](https://coveralls.io/github/lifezq/redis?branch=master)
[![License](http://img.shields.io/badge/license-apache-blue.svg?style=flat-square)](https://raw.githubusercontent.com/lifezq/redis/master/LICENSE)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/lifezq/redis)

# Redis Client Wrapper

这是一个基于 `github.com/redis/go-redis/v9` 封装的 Redis 客户端库，旨在简化 Redis 的使用，提供统一的 API 接口支持单机、集群和哨兵模式，并内置了灵活的 Hook 机制。

## 功能特性

*   **多模式支持**: 统一的初始化接口，支持单机 (Standalone)、集群 (Cluster) 和哨兵 (Sentinel) 模式。
*   **统一 API**: 无论底层使用何种部署模式，上层调用接口保持一致。
*   **Hook 机制**: 支持 `BeforeFunc` 和 `AfterFunc`，可在 Redis 操作前后注入自定义逻辑（如日志记录、链路追踪）。
*   **全功能覆盖**:
    *   基础数据类型: String, Hash, List, Set, ZSet
    *   高级功能: PubSub, Geo, HyperLogLog, Bitmap
    *   扩展模块: Bloom Filter (需要 RedisBloom 模块支持)
    *   通用操作: Pipeline, Key 管理, 自定义命令 (Do/Process)

## 安装

```bash
go get github.com/lifezq/redis
```

## 快速开始

### 初始化

通过 `redis.Options` 配置连接参数。`Mode` 字段决定了连接模式：
*   `0`: 单机模式
*   `1`: 集群模式
*   `2`: 哨兵模式

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/lifezq/redis"
)

func main() {
	// 1. 初始化配置
	opt := redis.Options{
		// 地址列表
		// 单机: []string{"127.0.0.1:6379"}
		// 集群: []string{"127.0.0.1:7000", "127.0.0.1:7001", ...}
		// 哨兵: []string{"127.0.0.1:26379", "127.0.0.1:26380", ...}
		Host: []string{"127.0.0.1:6379"},
		
		Mode:         0, // 0:单机, 1:集群, 2:哨兵
		Password:     "",
		DB:           0,
		PoolSize:     10,
		MinIdleConns: 5,
		
		// 哨兵模式专用配置
		// MasterName: "mymaster",
		// SentinelPassword: "",
	}

	// 2. 初始化客户端
	if err := redis.Init(opt); err != nil {
		panic(err)
	}
	
	// 获取客户端实例
	rc := redis.Client()

	// 3. 使用客户端
	ctx := context.Background()
	
	// String 操作
	rc.Set(ctx, "key", "value", time.Minute)
	val := rc.Get(ctx, "key").Val()
	fmt.Println("Get:", val)
}
```

### 使用 Hooks

可以在初始化时传入 `BeforeFunc` 和 `AfterFunc` 来拦截所有 Redis 命令。

```go
opt := redis.Options{
    Host: []string{"127.0.0.1:6379"},
    
    // 操作前执行
    BeforeFunc: func(ctx context.Context) context.Context {
        fmt.Println("Preparing to execute redis command...")
        // 可以在这里添加 trace ID 等信息到 context
        return ctx
    },
    
    // 操作后执行
    AfterFunc: func(ctx context.Context, cmd rds.Cmder) {
        fmt.Printf("Command finished: %s, duration: %d\n", cmd.Name(), cmd.Err())
    },
}

_ = redis.Init(opt)
rc := redis.Client()
```

### 全局 Key 前缀 (Global Key Prefix)

可以通过配置 `Prefix` 字段来为所有 Redis 操作自动添加统一的 Key 前缀。这在多租户环境或需要隔离 Key 空间的场景下非常有用。

该功能通过 Hook 机制实现，支持：
*   所有基础命令 (Get, Set, HSet, ...)
*   批量操作 (MGet, MSet, Pipeline)
*   Key 扫描与匹配 (Keys, Scan, HScan, ...)

```go
opt := redis.Options{
    Host:   []string{"127.0.0.1:6379"},
    Prefix: "app_v1:", // 设置全局前缀
}

redis.Init(opt)
rc := redis.Client()

// 实际操作的 Key 为 "app_v1:mykey"
rc.Set(ctx, "mykey", "value", 0)

// Scan 时会自动处理 MATCH 模式，实际匹配 "app_v1:pattern*"
rc.Scan(ctx, 0, "pattern*", 10)
```

### 高级功能示例

#### 布隆过滤器 (Bloom Filter)

```go
// 添加元素
rc.BFAdd(ctx, "myBloom", "item1")

// 检查是否存在
exists := rc.BFExists(ctx, "myBloom", "item1").Val()
fmt.Println("Exists:", exists)
```

#### 发布订阅 (PubSub)

```go
// 订阅
pubsub := rc.Subscribe(ctx, "mychannel")
defer pubsub.Close()

// 处理消息
ch := pubsub.Channel()
for msg := range ch {
    fmt.Println(msg.Channel, msg.Payload)
}

// 发布 (在另一个 goroutine 或进程中)
rc.Publish(ctx, "mychannel", "hello world")
```

#### Pipeline 批量操作

```go
cmds, err := rc.Pipelined(ctx, func(pipe rds.Pipeliner) error {
    pipe.Set(ctx, "key1", "value1", 0)
    pipe.Set(ctx, "key2", "value2", 0)
    return nil
})
```

## 文件结构说明

*   `redis.go`: 核心定义、初始化逻辑、Hook 实现
*   `prefix_hook.go`: 全局 Key 前缀 Hook 实现
*   `common.go`: 通用 Key 操作、Pipeline、自定义命令支持
*   `string.go`: String 类型操作
*   `hash.go`: Hash 类型操作
*   `list.go`: List 类型操作
*   `set.go`: Set 类型操作
*   `zset.go`: Sorted Set 类型操作
*   `pubsub.go`: 发布订阅
*   `bloom.go`: 布隆过滤器
*   `geo.go`: 地理位置
*   `bitmap.go`: 位图操作
*   `hyperloglog.go`: 基数统计
