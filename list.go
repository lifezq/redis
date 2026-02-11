// Copyright 2026 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package redis

import (
	"context"
	"time"

	rds "github.com/redis/go-redis/v9"
)

// List 列表操作

func (rc *client) LPush(ctx context.Context, key string, values ...interface{}) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.LPush(ctx, key, values...)
	default:
		return rc.Client.LPush(ctx, key, values...)
	}
}

func (rc *client) LPushX(ctx context.Context, key string, values ...interface{}) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.LPushX(ctx, key, values...)
	default:
		return rc.Client.LPushX(ctx, key, values...)
	}
}

func (rc *client) RPush(ctx context.Context, key string, values ...interface{}) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.RPush(ctx, key, values...)
	default:
		return rc.Client.RPush(ctx, key, values...)
	}
}

func (rc *client) RPushX(ctx context.Context, key string, values ...interface{}) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.RPushX(ctx, key, values...)
	default:
		return rc.Client.RPushX(ctx, key, values...)
	}
}

func (rc *client) LPop(ctx context.Context, key string) *rds.StringCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.LPop(ctx, key)
	default:
		return rc.Client.LPop(ctx, key)
	}
}

func (rc *client) RPop(ctx context.Context, key string) *rds.StringCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.RPop(ctx, key)
	default:
		return rc.Client.RPop(ctx, key)
	}
}

func (rc *client) RPopLPush(ctx context.Context, source, destination string) *rds.StringCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.RPopLPush(ctx, source, destination)
	default:
		return rc.Client.RPopLPush(ctx, source, destination)
	}
}

func (rc *client) LLen(ctx context.Context, key string) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.LLen(ctx, key)
	default:
		return rc.Client.LLen(ctx, key)
	}
}

func (rc *client) LRange(ctx context.Context, key string, start, stop int64) *rds.StringSliceCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.LRange(ctx, key, start, stop)
	default:
		return rc.Client.LRange(ctx, key, start, stop)
	}
}

func (rc *client) LRem(ctx context.Context, key string, count int64, value interface{}) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.LRem(ctx, key, count, value)
	default:
		return rc.Client.LRem(ctx, key, count, value)
	}
}

func (rc *client) LSet(ctx context.Context, key string, index int64, value interface{}) *rds.StatusCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.LSet(ctx, key, index, value)
	default:
		return rc.Client.LSet(ctx, key, index, value)
	}
}

func (rc *client) LTrim(ctx context.Context, key string, start, stop int64) *rds.StatusCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.LTrim(ctx, key, start, stop)
	default:
		return rc.Client.LTrim(ctx, key, start, stop)
	}
}

func (rc *client) LIndex(ctx context.Context, key string, index int64) *rds.StringCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.LIndex(ctx, key, index)
	default:
		return rc.Client.LIndex(ctx, key, index)
	}
}

func (rc *client) BLPop(ctx context.Context, timeout time.Duration, keys ...string) *rds.StringSliceCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.BLPop(ctx, timeout, keys...)
	default:
		return rc.Client.BLPop(ctx, timeout, keys...)
	}
}

func (rc *client) BRPop(ctx context.Context, timeout time.Duration, keys ...string) *rds.StringSliceCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.BRPop(ctx, timeout, keys...)
	default:
		return rc.Client.BRPop(ctx, timeout, keys...)
	}
}
