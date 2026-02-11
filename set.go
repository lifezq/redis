// Copyright 2026 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package redis

import (
	"context"

	rds "github.com/redis/go-redis/v9"
)

// Set 集合操作

func (rc *client) SAdd(ctx context.Context, key string, members ...interface{}) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.SAdd(ctx, key, members...)
	default:
		return rc.Client.SAdd(ctx, key, members...)
	}
}

func (rc *client) SRem(ctx context.Context, key string, members ...interface{}) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.SRem(ctx, key, members...)
	default:
		return rc.Client.SRem(ctx, key, members...)
	}
}

func (rc *client) SCard(ctx context.Context, key string) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.SCard(ctx, key)
	default:
		return rc.Client.SCard(ctx, key)
	}
}

func (rc *client) SIsMember(ctx context.Context, key string, member interface{}) *rds.BoolCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.SIsMember(ctx, key, member)
	default:
		return rc.Client.SIsMember(ctx, key, member)
	}
}

func (rc *client) SMIsMember(ctx context.Context, key string, members ...interface{}) *rds.BoolSliceCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.SMIsMember(ctx, key, members...)
	default:
		return rc.Client.SMIsMember(ctx, key, members...)
	}
}

func (rc *client) SMembers(ctx context.Context, key string) *rds.StringSliceCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.SMembers(ctx, key)
	default:
		return rc.Client.SMembers(ctx, key)
	}
}

func (rc *client) SPop(ctx context.Context, key string) *rds.StringCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.SPop(ctx, key)
	default:
		return rc.Client.SPop(ctx, key)
	}
}

func (rc *client) SPopN(ctx context.Context, key string, count int64) *rds.StringSliceCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.SPopN(ctx, key, count)
	default:
		return rc.Client.SPopN(ctx, key, count)
	}
}

func (rc *client) SRandMember(ctx context.Context, key string) *rds.StringCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.SRandMember(ctx, key)
	default:
		return rc.Client.SRandMember(ctx, key)
	}
}

func (rc *client) SRandMemberN(ctx context.Context, key string, count int64) *rds.StringSliceCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.SRandMemberN(ctx, key, count)
	default:
		return rc.Client.SRandMemberN(ctx, key, count)
	}
}

func (rc *client) SDiff(ctx context.Context, keys ...string) *rds.StringSliceCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.SDiff(ctx, keys...)
	default:
		return rc.Client.SDiff(ctx, keys...)
	}
}

func (rc *client) SDiffStore(ctx context.Context, destination string, keys ...string) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.SDiffStore(ctx, destination, keys...)
	default:
		return rc.Client.SDiffStore(ctx, destination, keys...)
	}
}

func (rc *client) SInter(ctx context.Context, keys ...string) *rds.StringSliceCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.SInter(ctx, keys...)
	default:
		return rc.Client.SInter(ctx, keys...)
	}
}

func (rc *client) SInterStore(ctx context.Context, destination string, keys ...string) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.SInterStore(ctx, destination, keys...)
	default:
		return rc.Client.SInterStore(ctx, destination, keys...)
	}
}

func (rc *client) SUnion(ctx context.Context, keys ...string) *rds.StringSliceCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.SUnion(ctx, keys...)
	default:
		return rc.Client.SUnion(ctx, keys...)
	}
}

func (rc *client) SUnionStore(ctx context.Context, destination string, keys ...string) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.SUnionStore(ctx, destination, keys...)
	default:
		return rc.Client.SUnionStore(ctx, destination, keys...)
	}
}
