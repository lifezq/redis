// Copyright 2026 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package redis

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"

	rds "github.com/redis/go-redis/v9"
)

const (
	modeSingle   = iota // 单机模式
	modeCluster         // 集群模式
	modeSentinel        // 哨兵模式
)

type BeforeFunc func(context.Context) context.Context
type AfterFunc func(context.Context, rds.Cmder)

type Options struct {
	Host             []string
	DB               int
	Username         string // ACL username
	Password         string
	SentinelPassword string // Sentinel password
	MasterName       string // Sentinel master name
	Mode             int    `desc:"0单机,1集群,2哨兵"`
	WriteTimeout     time.Duration
	ReadTimeout      time.Duration
	PoolSize         int
	MinIdleConns     int
	Prefix           string // Key prefix

	BeforeFunc BeforeFunc
	AfterFunc  AfterFunc
}

type client struct {
	Client  *rds.Client
	Cluster *rds.ClusterClient
	mode    int
}

// universalHook 实现了 redis.Hook 接口，用于统一处理操作前后的注入逻辑
type universalHook struct {
	before func(context.Context) context.Context
	after  func(context.Context, rds.Cmder)
}

func (h *universalHook) DialHook(next rds.DialHook) rds.DialHook {
	return next
}

func (h *universalHook) ProcessHook(next rds.ProcessHook) rds.ProcessHook {
	return func(ctx context.Context, cmd rds.Cmder) error {
		if h.before != nil {
			ctx = h.before(ctx)
		}
		err := next(ctx, cmd)
		if h.after != nil {
			h.after(ctx, cmd)
		}
		return err
	}
}

func (h *universalHook) ProcessPipelineHook(next rds.ProcessPipelineHook) rds.ProcessPipelineHook {
	return func(ctx context.Context, cmds []rds.Cmder) error {
		if h.before != nil {
			ctx = h.before(ctx)
		}
		err := next(ctx, cmds)
		// 批量操作暂只支持 BeforeFunc，AfterFunc 逻辑较为复杂（针对单个命令还是整体），
		// 这里选择对每个命令调用一次 AfterFunc，以保持与 ProcessHook 一致的行为模式
		if h.after != nil {
			for _, cmd := range cmds {
				h.after(ctx, cmd)
			}
		}
		return err
	}
}

var (
	defaultClient atomic.Pointer[client]
	initMu        sync.Mutex
)

// Client returns the initialized redis instance.
func Client() *client {
	c := defaultClient.Load()
	if c == nil {
		panic("redis: client not initialized, please call redis.Init(...) first")
	}
	return c
}

func Init(opt Options) error {
	initMu.Lock()
	defer initMu.Unlock()

	if len(opt.Host) == 0 {
		return errors.New("redis host is empty")
	}

	rc := &client{
		mode: opt.Mode,
	}

	// 准备 Hook
	var uHook *universalHook
	if opt.BeforeFunc != nil || opt.AfterFunc != nil {
		uHook = &universalHook{
			before: opt.BeforeFunc,
			after:  opt.AfterFunc,
		}
	}

	var pHook *prefixHook
	if opt.Prefix != "" {
		pHook = &prefixHook{prefix: opt.Prefix}
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	switch opt.Mode {
	case modeCluster:
		rc.Cluster = rds.NewClusterClient(&rds.ClusterOptions{
			Addrs:        opt.Host,
			Username:     opt.Username,
			Password:     opt.Password,
			WriteTimeout: opt.WriteTimeout,
			ReadTimeout:  opt.ReadTimeout,
			PoolSize:     opt.PoolSize,
			MinIdleConns: opt.MinIdleConns,
		})

		if pHook != nil {
			rc.Cluster.AddHook(pHook)
		}
		if uHook != nil {
			rc.Cluster.AddHook(uHook)
		}

		// 检测
		if _, err := rc.Cluster.Ping(ctx).Result(); err != nil {
			return err
		}

	case modeSentinel:
		if opt.MasterName == "" {
			return errors.New("sentinel master name is required")
		}
		rc.Client = rds.NewFailoverClient(&rds.FailoverOptions{
			MasterName:       opt.MasterName,
			SentinelAddrs:    opt.Host,
			Username:         opt.Username,
			Password:         opt.Password,
			SentinelPassword: opt.SentinelPassword,
			DB:               opt.DB,
			WriteTimeout:     opt.WriteTimeout,
			ReadTimeout:      opt.ReadTimeout,
			PoolSize:         opt.PoolSize,
			MinIdleConns:     opt.MinIdleConns,
		})

		if pHook != nil {
			rc.Client.AddHook(pHook)
		}
		if uHook != nil {
			rc.Client.AddHook(uHook)
		}

		// 检测
		if _, err := rc.Client.Ping(ctx).Result(); err != nil {
			return err
		}

	default:
		// 默认单机
		rc.Client = rds.NewClient(&rds.Options{
			Addr:         opt.Host[0],
			Username:     opt.Username,
			Password:     opt.Password,
			DB:           opt.DB,
			WriteTimeout: opt.WriteTimeout,
			ReadTimeout:  opt.ReadTimeout,
			PoolSize:     opt.PoolSize,
			MinIdleConns: opt.MinIdleConns,
		})

		if pHook != nil {
			rc.Client.AddHook(pHook)
		}
		if uHook != nil {
			rc.Client.AddHook(uHook)
		}

		// 检测
		if _, err := rc.Client.Ping(ctx).Result(); err != nil {
			return err
		}
	}

	defaultClient.Store(rc)
	return nil
}

func (rc *client) TxPipelined(ctx context.Context, fn func(rds.Pipeliner) error) ([]rds.Cmder, error) {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.TxPipelined(ctx, fn)
	default:
		return rc.Client.TxPipelined(ctx, fn)
	}
}

func (rc *client) TxPipeline() rds.Pipeliner {
	switch rc.mode {
	case modeCluster:
		return rc.Cluster.TxPipeline()
	default:
		return rc.Client.TxPipeline()
	}
}
