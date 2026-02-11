// Copyright 2026 The Goutils Author. All Rights Reserved.
//
// -------------------------------------------------------------------

package redis

import (
	"testing"
)

func TestInit_Sentinel_MissingMasterName(t *testing.T) {
	opt := Options{
		Host: []string{"sentinel:26379"},
		Mode: modeSentinel,
	}
	err := Init(opt)
	if err == nil {
		t.Error("Expected error for missing MasterName in sentinel mode, got nil")
	}
}

func TestInit_Compilation_Sentinel(t *testing.T) {
	opt := Options{
		Host:       []string{"localhost:26379"},
		MasterName: "mymaster",
		Mode:       modeSentinel,
	}
	err := Init(opt)
	if err == nil {
		// Connected (unlikely)
	} else {
		t.Logf("Expected connection error (no sentinel server): %v", err)
	}
}

func TestInstance_Panic(t *testing.T) {
	// Reset client for test
	defaultClient.Store(nil)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		} else {
			t.Logf("Recovered from panic as expected: %v", r)
		}
	}()

	Client()
}
