// Copyright 2026 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package redis

import (
	"context"

	rds "github.com/redis/go-redis/v9"
)

// Bitmap 位图

func (rc *client) SetBit(ctx context.Context, key string, offset int64, value int) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.SetBit(ctx, key, offset, value)
	default:
		return rc.Client.SetBit(ctx, key, offset, value)
	}
}

func (rc *client) GetBit(ctx context.Context, key string, offset int64) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.GetBit(ctx, key, offset)
	default:
		return rc.Client.GetBit(ctx, key, offset)
	}
}

func (rc *client) BitCount(ctx context.Context, key string, bitCount *rds.BitCount) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.BitCount(ctx, key, bitCount)
	default:
		return rc.Client.BitCount(ctx, key, bitCount)
	}
}

func (rc *client) BitOpAnd(ctx context.Context, destKey string, keys ...string) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.BitOpAnd(ctx, destKey, keys...)
	default:
		return rc.Client.BitOpAnd(ctx, destKey, keys...)
	}
}

func (rc *client) BitOpOr(ctx context.Context, destKey string, keys ...string) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.BitOpOr(ctx, destKey, keys...)
	default:
		return rc.Client.BitOpOr(ctx, destKey, keys...)
	}
}

func (rc *client) BitOpXor(ctx context.Context, destKey string, keys ...string) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.BitOpXor(ctx, destKey, keys...)
	default:
		return rc.Client.BitOpXor(ctx, destKey, keys...)
	}
}

func (rc *client) BitOpNot(ctx context.Context, destKey string, key string) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.BitOpNot(ctx, destKey, key)
	default:
		return rc.Client.BitOpNot(ctx, destKey, key)
	}
}

func (rc *client) BitPos(ctx context.Context, key string, bit int64, pos ...int64) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.BitPos(ctx, key, bit, pos...)
	default:
		return rc.Client.BitPos(ctx, key, bit, pos...)
	}
}

func (rc *client) BitField(ctx context.Context, key string, args ...interface{}) *rds.IntSliceCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.BitField(ctx, key, args...)
	default:
		return rc.Client.BitField(ctx, key, args...)
	}
}
