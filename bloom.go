// Copyright 2026 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package redis

import (
	"context"

	rds "github.com/redis/go-redis/v9"
)

// BloomFilter 布隆过滤器 (Requires RedisBloom module)

// BFAdd adds an item to the Bloom Filter
func (rc *client) BFAdd(ctx context.Context, key string, item interface{}) *rds.BoolCmd {
	cmd := rds.NewBoolCmd(ctx, "BF.ADD", key, item)
	_ = rc.Process(ctx, cmd)
	return cmd
}

// BFMAdd adds multiple items to the Bloom Filter
func (rc *client) BFMAdd(ctx context.Context, key string, items ...interface{}) *rds.BoolSliceCmd {
	args := make([]interface{}, 2+len(items))
	args[0] = "BF.MADD"
	args[1] = key
	copy(args[2:], items)
	cmd := rds.NewBoolSliceCmd(ctx, args...)
	_ = rc.Process(ctx, cmd)
	return cmd
}

// BFExists checks if an item exists in the Bloom Filter
func (rc *client) BFExists(ctx context.Context, key string, item interface{}) *rds.BoolCmd {
	cmd := rds.NewBoolCmd(ctx, "BF.EXISTS", key, item)
	_ = rc.Process(ctx, cmd)
	return cmd
}

// BFMExists checks if multiple items exist in the Bloom Filter
func (rc *client) BFMExists(ctx context.Context, key string, items ...interface{}) *rds.BoolSliceCmd {
	args := make([]interface{}, 2+len(items))
	args[0] = "BF.MEXISTS"
	args[1] = key
	copy(args[2:], items)
	cmd := rds.NewBoolSliceCmd(ctx, args...)
	_ = rc.Process(ctx, cmd)
	return cmd
}

// BFReserve creates a new Bloom Filter
func (rc *client) BFReserve(ctx context.Context, key string, errorRate float64, capacity int64) *rds.StatusCmd {
	cmd := rds.NewStatusCmd(ctx, "BF.RESERVE", key, errorRate, capacity)
	_ = rc.Process(ctx, cmd)
	return cmd
}

// BFInsert adds items to the Bloom Filter with options
// Options can be passed as strings, e.g. "CAPACITY", "1000", "ERROR", "0.01", "NOCREATE"
func (rc *client) BFInsert(ctx context.Context, key string, options []string, items ...interface{}) *rds.BoolSliceCmd {
	args := make([]interface{}, 2+len(options)+1+len(items)) // command + key + options + ITEMS keyword + items
	args[0] = "BF.INSERT"
	args[1] = key
	idx := 2
	for _, opt := range options {
		args[idx] = opt
		idx++
	}
	args[idx] = "ITEMS"
	idx++
	for _, item := range items {
		args[idx] = item
		idx++
	}

	cmd := rds.NewBoolSliceCmd(ctx, args...)
	_ = rc.Process(ctx, cmd)
	return cmd
}

// BFInfo returns information about the Bloom Filter
func (rc *client) BFInfo(ctx context.Context, key string) *rds.Cmd {
	cmd := rds.NewCmd(ctx, "BF.INFO", key)
	_ = rc.Process(ctx, cmd)
	return cmd
}
