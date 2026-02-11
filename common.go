// Copyright 2026 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package redis

import (
	"context"
	"time"

	rds "github.com/redis/go-redis/v9"
)

// 通用 Key 操作

func (rc *client) Del(ctx context.Context, keys ...string) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.Del(ctx, keys...)
	default:
		return rc.Client.Del(ctx, keys...)
	}
}

func (rc *client) Unlink(ctx context.Context, keys ...string) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.Unlink(ctx, keys...)
	default:
		return rc.Client.Unlink(ctx, keys...)
	}
}

func (rc *client) Expire(ctx context.Context, key string, expiration time.Duration) *rds.BoolCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.Expire(ctx, key, expiration)
	default:
		return rc.Client.Expire(ctx, key, expiration)
	}
}

func (rc *client) ExpireAt(ctx context.Context, key string, tm time.Time) *rds.BoolCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.ExpireAt(ctx, key, tm)
	default:
		return rc.Client.ExpireAt(ctx, key, tm)
	}
}

func (rc *client) TTL(ctx context.Context, key string) *rds.DurationCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.TTL(ctx, key)
	default:
		return rc.Client.TTL(ctx, key)
	}
}

func (rc *client) Exists(ctx context.Context, keys ...string) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.Exists(ctx, keys...)
	default:
		return rc.Client.Exists(ctx, keys...)
	}
}

func (rc *client) Type(ctx context.Context, key string) *rds.StatusCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.Type(ctx, key)
	default:
		return rc.Client.Type(ctx, key)
	}
}

func (rc *client) Keys(ctx context.Context, pattern string) *rds.StringSliceCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.Keys(ctx, pattern)
	default:
		return rc.Client.Keys(ctx, pattern)
	}
}

func (rc *client) Scan(ctx context.Context, cursor uint64, match string, count int64) *rds.ScanCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.Scan(ctx, cursor, match, count)
	default:
		return rc.Client.Scan(ctx, cursor, match, count)
	}
}

func (rc *client) Eval(ctx context.Context, script string, keys []string, args ...interface{}) *rds.Cmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.Eval(ctx, script, keys, args...)
	default:
		return rc.Client.Eval(ctx, script, keys, args...)
	}
}

func (rc *client) EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *rds.Cmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.EvalSha(ctx, sha1, keys, args...)
	default:
		return rc.Client.EvalSha(ctx, sha1, keys, args...)
	}
}

// Do executes any Redis command
func (rc *client) Do(ctx context.Context, args ...interface{}) *rds.Cmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.Do(ctx, args...)
	default:
		return rc.Client.Do(ctx, args...)
	}
}

// Process executes the command
func (rc *client) Process(ctx context.Context, cmd rds.Cmder) error {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.Process(ctx, cmd)
	default:
		return rc.Client.Process(ctx, cmd)
	}
}

// Pipeline returns a new Pipeliner
func (rc *client) Pipeline() rds.Pipeliner {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.Pipeline()
	default:
		return rc.Client.Pipeline()
	}
}

// Pipelined executes a function with a Pipeliner
func (rc *client) Pipelined(ctx context.Context, fn func(rds.Pipeliner) error) ([]rds.Cmder, error) {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.Pipelined(ctx, fn)
	default:
		return rc.Client.Pipelined(ctx, fn)
	}
}
