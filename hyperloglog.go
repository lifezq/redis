// Copyright 2026 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package redis

import (
	"context"

	rds "github.com/redis/go-redis/v9"
)

// HyperLogLog 基数统计

func (rc *client) PFAdd(ctx context.Context, key string, els ...interface{}) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.PFAdd(ctx, key, els...)
	default:
		return rc.Client.PFAdd(ctx, key, els...)
	}
}

func (rc *client) PFCount(ctx context.Context, keys ...string) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.PFCount(ctx, keys...)
	default:
		return rc.Client.PFCount(ctx, keys...)
	}
}

func (rc *client) PFMerge(ctx context.Context, dest string, keys ...string) *rds.StatusCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.PFMerge(ctx, dest, keys...)
	default:
		return rc.Client.PFMerge(ctx, dest, keys...)
	}
}
