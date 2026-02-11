// Copyright 2026 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package redis

import (
	"context"

	rds "github.com/redis/go-redis/v9"
)

// PubSub 订阅发布

func (rc *client) Publish(ctx context.Context, channel string, message interface{}) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.Publish(ctx, channel, message)
	default:
		return rc.Client.Publish(ctx, channel, message)
	}
}

func (rc *client) Subscribe(ctx context.Context, channels ...string) *rds.PubSub {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.Subscribe(ctx, channels...)
	default:
		return rc.Client.Subscribe(ctx, channels...)
	}
}

func (rc *client) PSubscribe(ctx context.Context, channels ...string) *rds.PubSub {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.PSubscribe(ctx, channels...)
	default:
		return rc.Client.PSubscribe(ctx, channels...)
	}
}

// ShardedSubscribe Redis 7.0+ Sharded Pub/Sub
func (rc *client) ShardedSubscribe(ctx context.Context, channels ...string) *rds.PubSub {
	// go-redis v9 supports ShardedSubscribe but it might be method on client.
	// Let's check standard interface.
	// Assuming it's available as SSubscribe (Sharded Subscribe) or similar.
	// Actually go-redis v9 has SSubscribe.
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.SSubscribe(ctx, channels...)
	default:
		return rc.Client.SSubscribe(ctx, channels...)
	}
}

func (rc *client) PubSubChannels(ctx context.Context, pattern string) *rds.StringSliceCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.PubSubChannels(ctx, pattern)
	default:
		return rc.Client.PubSubChannels(ctx, pattern)
	}
}

func (rc *client) PubSubNumSub(ctx context.Context, channels ...string) *rds.MapStringIntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.PubSubNumSub(ctx, channels...)
	default:
		return rc.Client.PubSubNumSub(ctx, channels...)
	}
}

func (rc *client) PubSubNumPat(ctx context.Context) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.PubSubNumPat(ctx)
	default:
		return rc.Client.PubSubNumPat(ctx)
	}
}
