// Copyright 2026 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package redis

import (
	"context"
	"testing"

	rds "github.com/redis/go-redis/v9"
)

func TestPrefixHook(t *testing.T) {
	// This test requires a running Redis instance at localhost:6379

	prefix := "test_prefix:"
	opt := Options{
		Host:   []string{"localhost:6379"},
		Prefix: prefix,
	}

	if err := Init(opt); err != nil {
		t.Logf("Failed to init redis: %v (skipping test if no redis)", err)
		return
	}

	ctx := context.Background()
	client := Client()

	key := "mykey"
	val := "myval"

	// Test SET
	if err := client.Set(ctx, key, val, 0).Err(); err != nil {
		t.Errorf("Set failed: %v", err)
	}

	// Verify directly with go-redis (without prefix) to see if key has prefix
	rawClient := rds.NewClient(&rds.Options{
		Addr: "localhost:6379",
	})
	defer rawClient.Close()

	gotVal, err := rawClient.Get(ctx, prefix+key).Result()
	if err != nil {
		t.Errorf("Raw Get failed: %v. Prefix '%s' was likely not applied to key '%s'", err, prefix, key)
	}
	if gotVal != val {
		t.Errorf("Expected %s, got %s", val, gotVal)
	}

	// Test GET
	gotVal2, err := client.Get(ctx, key).Result()
	if err != nil {
		t.Errorf("Client Get failed: %v", err)
	}
	if gotVal2 != val {
		t.Errorf("Client Expected %s, got %s", val, gotVal2)
	}

	// Test Pipeline
	pipe := client.Pipeline()
	pipe.Set(ctx, "pipe_key", "pipe_val", 0)
	pipe.Get(ctx, "pipe_key")
	cmds, err := pipe.Exec(ctx)
	if err != nil {
		t.Errorf("Pipeline failed: %v", err)
	}

	// Check pipeline results
	if cmds[0].Err() != nil {
		t.Errorf("Pipeline Set failed: %v", cmds[0].Err())
	}
	if val, _ := cmds[1].(*rds.StringCmd).Result(); val != "pipe_val" {
		t.Errorf("Pipeline Get expected pipe_val, got %s", val)
	}

	// Verify pipeline key
	gotPipeVal, err := rawClient.Get(ctx, prefix+"pipe_key").Result()
	if err != nil {
		t.Errorf("Raw Get pipe_key failed: %v", err)
	}
	if gotPipeVal != "pipe_val" {
		t.Errorf("Raw pipe_key expected pipe_val, got %s", gotPipeVal)
	}

	// Clean up
	rawClient.Del(ctx, prefix+key, prefix+"pipe_key")
}
