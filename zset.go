// Copyright 2026 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package redis

import (
	"context"

	rds "github.com/redis/go-redis/v9"
)

// ZSet 有序集合操作

func (rc *client) ZAdd(ctx context.Context, key string, members ...rds.Z) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.ZAdd(ctx, key, members...)
	default:
		return rc.Client.ZAdd(ctx, key, members...)
	}
}

func (rc *client) ZAddNX(ctx context.Context, key string, members ...rds.Z) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.ZAddNX(ctx, key, members...)
	default:
		return rc.Client.ZAddNX(ctx, key, members...)
	}
}

func (rc *client) ZAddXX(ctx context.Context, key string, members ...rds.Z) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.ZAddXX(ctx, key, members...)
	default:
		return rc.Client.ZAddXX(ctx, key, members...)
	}
}

func (rc *client) ZRem(ctx context.Context, key string, members ...interface{}) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.ZRem(ctx, key, members...)
	default:
		return rc.Client.ZRem(ctx, key, members...)
	}
}

func (rc *client) ZIncrBy(ctx context.Context, key string, increment float64, member string) *rds.FloatCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.ZIncrBy(ctx, key, increment, member)
	default:
		return rc.Client.ZIncrBy(ctx, key, increment, member)
	}
}

func (rc *client) ZCard(ctx context.Context, key string) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.ZCard(ctx, key)
	default:
		return rc.Client.ZCard(ctx, key)
	}
}

func (rc *client) ZScore(ctx context.Context, key, member string) *rds.FloatCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.ZScore(ctx, key, member)
	default:
		return rc.Client.ZScore(ctx, key, member)
	}
}

func (rc *client) ZCount(ctx context.Context, key, min, max string) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.ZCount(ctx, key, min, max)
	default:
		return rc.Client.ZCount(ctx, key, min, max)
	}
}

func (rc *client) ZRank(ctx context.Context, key, member string) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.ZRank(ctx, key, member)
	default:
		return rc.Client.ZRank(ctx, key, member)
	}
}

func (rc *client) ZRevRank(ctx context.Context, key, member string) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.ZRevRank(ctx, key, member)
	default:
		return rc.Client.ZRevRank(ctx, key, member)
	}
}

func (rc *client) ZRange(ctx context.Context, key string, start, stop int64) *rds.StringSliceCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.ZRange(ctx, key, start, stop)
	default:
		return rc.Client.ZRange(ctx, key, start, stop)
	}
}

func (rc *client) ZRangeWithScores(ctx context.Context, key string, start, stop int64) *rds.ZSliceCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.ZRangeWithScores(ctx, key, start, stop)
	default:
		return rc.Client.ZRangeWithScores(ctx, key, start, stop)
	}
}

func (rc *client) ZRevRange(ctx context.Context, key string, start, stop int64) *rds.StringSliceCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.ZRevRange(ctx, key, start, stop)
	default:
		return rc.Client.ZRevRange(ctx, key, start, stop)
	}
}

func (rc *client) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *rds.ZSliceCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.ZRevRangeWithScores(ctx, key, start, stop)
	default:
		return rc.Client.ZRevRangeWithScores(ctx, key, start, stop)
	}
}

func (rc *client) ZRangeByScore(ctx context.Context, key string, opt *rds.ZRangeBy) *rds.StringSliceCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.ZRangeByScore(ctx, key, opt)
	default:
		return rc.Client.ZRangeByScore(ctx, key, opt)
	}
}

func (rc *client) ZRangeByScoreWithScores(ctx context.Context, key string, opt *rds.ZRangeBy) *rds.ZSliceCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.ZRangeByScoreWithScores(ctx, key, opt)
	default:
		return rc.Client.ZRangeByScoreWithScores(ctx, key, opt)
	}
}

func (rc *client) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.ZRemRangeByRank(ctx, key, start, stop)
	default:
		return rc.Client.ZRemRangeByRank(ctx, key, start, stop)
	}
}

func (rc *client) ZRemRangeByScore(ctx context.Context, key, min, max string) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.ZRemRangeByScore(ctx, key, min, max)
	default:
		return rc.Client.ZRemRangeByScore(ctx, key, min, max)
	}
}
