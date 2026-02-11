// Copyright 2026 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package redis

import (
	"context"

	rds "github.com/redis/go-redis/v9"
)

// Geo 地理位置

func (rc *client) GeoAdd(ctx context.Context, key string, geoLocation ...*rds.GeoLocation) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.GeoAdd(ctx, key, geoLocation...)
	default:
		return rc.Client.GeoAdd(ctx, key, geoLocation...)
	}
}

func (rc *client) GeoPos(ctx context.Context, key string, members ...string) *rds.GeoPosCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.GeoPos(ctx, key, members...)
	default:
		return rc.Client.GeoPos(ctx, key, members...)
	}
}

func (rc *client) GeoDist(ctx context.Context, key string, member1, member2, unit string) *rds.FloatCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.GeoDist(ctx, key, member1, member2, unit)
	default:
		return rc.Client.GeoDist(ctx, key, member1, member2, unit)
	}
}

func (rc *client) GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *rds.GeoRadiusQuery) *rds.GeoLocationCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.GeoRadius(ctx, key, longitude, latitude, query)
	default:
		return rc.Client.GeoRadius(ctx, key, longitude, latitude, query)
	}
}

func (rc *client) GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query *rds.GeoRadiusQuery) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.GeoRadiusStore(ctx, key, longitude, latitude, query)
	default:
		return rc.Client.GeoRadiusStore(ctx, key, longitude, latitude, query)
	}
}

func (rc *client) GeoRadiusByMember(ctx context.Context, key, member string, query *rds.GeoRadiusQuery) *rds.GeoLocationCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.GeoRadiusByMember(ctx, key, member, query)
	default:
		return rc.Client.GeoRadiusByMember(ctx, key, member, query)
	}
}

func (rc *client) GeoRadiusByMemberStore(ctx context.Context, key, member string, query *rds.GeoRadiusQuery) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.GeoRadiusByMemberStore(ctx, key, member, query)
	default:
		return rc.Client.GeoRadiusByMemberStore(ctx, key, member, query)
	}
}

func (rc *client) GeoSearch(ctx context.Context, key string, q *rds.GeoSearchQuery) *rds.StringSliceCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.GeoSearch(ctx, key, q)
	default:
		return rc.Client.GeoSearch(ctx, key, q)
	}
}

func (rc *client) GeoSearchLocation(ctx context.Context, key string, q *rds.GeoSearchLocationQuery) *rds.GeoSearchLocationCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.GeoSearchLocation(ctx, key, q)
	default:
		return rc.Client.GeoSearchLocation(ctx, key, q)
	}
}

func (rc *client) GeoSearchStore(ctx context.Context, key, store string, q *rds.GeoSearchStoreQuery) *rds.IntCmd {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.GeoSearchStore(ctx, key, store, q)
	default:
		return rc.Client.GeoSearchStore(ctx, key, store, q)
	}
}
