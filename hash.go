// Copyright 2026 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package redis

import (
	"context"

	rds "github.com/redis/go-redis/v9"
)

// Hash 哈希操作

func (rc *client) HSet(ctx context.Context, key string, values ...interface{}) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.HSet(ctx, key, values...)
	default:
		return rc.Client.HSet(ctx, key, values...)
	}
}

func (rc *client) HSetNX(ctx context.Context, key, field string, value interface{}) *rds.BoolCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.HSetNX(ctx, key, field, value)
	default:
		return rc.Client.HSetNX(ctx, key, field, value)
	}
}

func (rc *client) HGet(ctx context.Context, key, field string) *rds.StringCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.HGet(ctx, key, field)
	default:
		return rc.Client.HGet(ctx, key, field)
	}
}

func (rc *client) HGetAll(ctx context.Context, key string) *rds.MapStringStringCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.HGetAll(ctx, key)
	default:
		return rc.Client.HGetAll(ctx, key)
	}
}

func (rc *client) HDel(ctx context.Context, key string, fields ...string) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.HDel(ctx, key, fields...)
	default:
		return rc.Client.HDel(ctx, key, fields...)
	}
}

func (rc *client) HExists(ctx context.Context, key, field string) *rds.BoolCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.HExists(ctx, key, field)
	default:
		return rc.Client.HExists(ctx, key, field)
	}
}

func (rc *client) HIncrBy(ctx context.Context, key, field string, incr int64) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.HIncrBy(ctx, key, field, incr)
	default:
		return rc.Client.HIncrBy(ctx, key, field, incr)
	}
}

func (rc *client) HIncrByFloat(ctx context.Context, key, field string, incr float64) *rds.FloatCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.HIncrByFloat(ctx, key, field, incr)
	default:
		return rc.Client.HIncrByFloat(ctx, key, field, incr)
	}
}

func (rc *client) HKeys(ctx context.Context, key string) *rds.StringSliceCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.HKeys(ctx, key)
	default:
		return rc.Client.HKeys(ctx, key)
	}
}

func (rc *client) HVals(ctx context.Context, key string) *rds.StringSliceCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.HVals(ctx, key)
	default:
		return rc.Client.HVals(ctx, key)
	}
}

func (rc *client) HLen(ctx context.Context, key string) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.HLen(ctx, key)
	default:
		return rc.Client.HLen(ctx, key)
	}
}

func (rc *client) HMGet(ctx context.Context, key string, fields ...string) *rds.SliceCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.HMGet(ctx, key, fields...)
	default:
		return rc.Client.HMGet(ctx, key, fields...)
	}
}

// HMSet is deprecated in modern Redis (use HSet), but kept for compatibility if needed.
// go-redis v9 supports HMSet via HSet or explicit HMSet if available, but HSet accepts multiple values now.
func (rc *client) HMSet(ctx context.Context, key string, values ...interface{}) *rds.BoolCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.HMSet(ctx, key, values...)
	default:
		return rc.Client.HMSet(ctx, key, values...)
	}
}
