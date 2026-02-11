// Copyright 2026 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package redis

import (
	"context"
	"time"

	rds "github.com/redis/go-redis/v9"
)

// String 字符串操作

func (rc *client) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *rds.StatusCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.Set(ctx, key, value, expiration)
	default:
		return rc.Client.Set(ctx, key, value, expiration)
	}
}

func (rc *client) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *rds.BoolCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.SetNX(ctx, key, value, expiration)
	default:
		return rc.Client.SetNX(ctx, key, value, expiration)
	}
}

func (rc *client) SetEx(ctx context.Context, key string, value interface{}, expiration time.Duration) *rds.StatusCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.SetEx(ctx, key, value, expiration)
	default:
		return rc.Client.SetEx(ctx, key, value, expiration)
	}
}

func (rc *client) Get(ctx context.Context, key string) *rds.StringCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.Get(ctx, key)
	default:
		return rc.Client.Get(ctx, key)
	}
}

func (rc *client) GetDel(ctx context.Context, key string) *rds.StringCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.GetDel(ctx, key)
	default:
		return rc.Client.GetDel(ctx, key)
	}
}

func (rc *client) GetSet(ctx context.Context, key string, value interface{}) *rds.StringCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.GetSet(ctx, key, value)
	default:
		return rc.Client.GetSet(ctx, key, value)
	}
}

func (rc *client) Incr(ctx context.Context, key string) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.Incr(ctx, key)
	default:
		return rc.Client.Incr(ctx, key)
	}
}

func (rc *client) IncrBy(ctx context.Context, key string, value int64) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.IncrBy(ctx, key, value)
	default:
		return rc.Client.IncrBy(ctx, key, value)
	}
}

func (rc *client) IncrByFloat(ctx context.Context, key string, value float64) *rds.FloatCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.IncrByFloat(ctx, key, value)
	default:
		return rc.Client.IncrByFloat(ctx, key, value)
	}
}

func (rc *client) Decr(ctx context.Context, key string) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.Decr(ctx, key)
	default:
		return rc.Client.Decr(ctx, key)
	}
}

func (rc *client) DecrBy(ctx context.Context, key string, value int64) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.DecrBy(ctx, key, value)
	default:
		return rc.Client.DecrBy(ctx, key, value)
	}
}

func (rc *client) Append(ctx context.Context, key, value string) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.Append(ctx, key, value)
	default:
		return rc.Client.Append(ctx, key, value)
	}
}

func (rc *client) MGet(ctx context.Context, keys ...string) *rds.SliceCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.MGet(ctx, keys...)
	default:
		return rc.Client.MGet(ctx, keys...)
	}
}

func (rc *client) MSet(ctx context.Context, values ...interface{}) *rds.StatusCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.MSet(ctx, values...)
	default:
		return rc.Client.MSet(ctx, values...)
	}
}
