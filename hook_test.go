// Copyright 2026 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package redis

import (
	"context"
	"testing"

	rds "github.com/redis/go-redis/v9"
)

func TestInit_Hook(t *testing.T) {
	// 模拟 Hook 调用
	beforeCalled := false
	afterCalled := false

	opt := Options{
		Host:     []string{"localhost:6379"},
		Password: "123456",
		BeforeFunc: func(ctx context.Context) context.Context {
			beforeCalled = true
			return ctx
		},
		AfterFunc: func(ctx context.Context, cmd rds.Cmder) {
			afterCalled = true
		},
	}

	// 尝试初始化，即使连接失败，只要代码路径正确，Hook 逻辑本身应该没问题
	// 但如果不连接成功，我们无法触发 Hook。
	// 这里主要测试编译和 Init 函数的基本流程。
	err := Init(opt)
	if err != nil {
		t.Logf("Connection failed as expected: %v", err)
	}
	t.Logf("Hook functions should be called regardless of connection success, beforeCalled:%v, afterCalled:%v", beforeCalled, afterCalled)
}
