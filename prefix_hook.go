// Copyright 2026 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package redis

import (
	"context"
	"strings"

	rds "github.com/redis/go-redis/v9"
)

type prefixHook struct {
	prefix string
}

func (h *prefixHook) DialHook(next rds.DialHook) rds.DialHook {
	return next
}

func (h *prefixHook) ProcessHook(next rds.ProcessHook) rds.ProcessHook {
	return func(ctx context.Context, cmd rds.Cmder) error {
		h.prefixKey(cmd)
		return next(ctx, cmd)
	}
}

func (h *prefixHook) ProcessPipelineHook(next rds.ProcessPipelineHook) rds.ProcessPipelineHook {
	return func(ctx context.Context, cmds []rds.Cmder) error {
		for _, cmd := range cmds {
			h.prefixKey(cmd)
		}
		return next(ctx, cmds)
	}
}

func (h *prefixHook) prefixKey(cmd rds.Cmder) {
	args := cmd.Args()
	if len(args) < 2 {
		return
	}
	name := strings.ToLower(cmd.Name())

	switch name {
	// Simple Key (Arg 1)
	case "get", "set", "setnx", "setex", "expire", "expireat", "ttl", "pttl",
		"type", "exists", "incr", "decr", "incrby", "decrby", "incrbyfloat",
		"append", "getset", "getdel", "getex",
		"hget", "hset", "hmset", "hsetnx", "hmget", "hgetall", "hkeys", "hvals", "hlen", "hexists", "hdel", "hincrby", "hincrbyfloat",
		"lpush", "rpush", "lpushx", "rpushx", "lpop", "rpop", "llen", "lrange", "lindex", "lset", "lrem", "ltrim",
		"sadd", "srem", "smembers", "sismember", "scard", "spop", "srandmember", "spopn", "srandmembern",
		"zadd", "zrem", "zrange", "zrevrange", "zcard", "zscore", "zrank", "zrevrank", "zincrby", "zcount", "zrangebyscore", "zremrangebyrank", "zremrangebyscore",
		"pfadd", "pfcount", "pfmerge", "geoadd", "geopos", "geodist", "georadius", "georadiusbymember",
		"keys", "sscan", "hscan", "zscan":
		if key, ok := args[1].(string); ok {
			args[1] = h.prefix + key
		}

	// Multi Key (Arg 1...N)
	case "mget", "del", "unlink", "touch", "watch":
		for i := 1; i < len(args); i++ {
			if key, ok := args[i].(string); ok {
				args[i] = h.prefix + key
			}
		}

	// MSET (Key Value Key Value...)
	case "mset", "msetnx":
		for i := 1; i < len(args); i += 2 {
			if key, ok := args[i].(string); ok {
				args[i] = h.prefix + key
			}
		}

	// SMOVE source destination member
	case "smove":
		if key, ok := args[1].(string); ok {
			args[1] = h.prefix + key
		}
		if key, ok := args[2].(string); ok {
			args[2] = h.prefix + key
		}

	// RPOPLPUSH source destination
	case "rpoplpush", "brpoplpush":
		if key, ok := args[1].(string); ok {
			args[1] = h.prefix + key
		}
		if key, ok := args[2].(string); ok {
			args[2] = h.prefix + key
		}

	// BLPOP key... timeout
	case "blpop", "brpop":
		// Last arg is timeout
		for i := 1; i < len(args)-1; i++ {
			if key, ok := args[i].(string); ok {
				args[i] = h.prefix + key
			}
		}

	// SINTER, SUNION, SDIFF (Key...)
	// SINTERSTORE, SUNIONSTORE, SDIFFSTORE (Dest Key...)
	case "sinter", "sunion", "sdiff":
		for i := 1; i < len(args); i++ {
			if key, ok := args[i].(string); ok {
				args[i] = h.prefix + key
			}
		}
	case "sinterstore", "sunionstore", "sdiffstore":
		for i := 1; i < len(args); i++ {
			if key, ok := args[i].(string); ok {
				args[i] = h.prefix + key
			}
		}
	}

	// Handle MATCH for scan family
	if name == "scan" || name == "sscan" || name == "hscan" || name == "zscan" {
		for i := 0; i < len(args)-1; i++ {
			if s, ok := args[i].(string); ok && strings.ToLower(s) == "match" {
				if pattern, ok := args[i+1].(string); ok {
					args[i+1] = h.prefix + pattern
				}
			}
		}
	}
}
