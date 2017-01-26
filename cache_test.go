/*
* File Name:	cache_test.go
* Description:
* Author:	Chapman Ou <ochapman.cn@gmail.com>
* Created:	2017-01-18
 */

package cache_test

import (
	"testing"
	"time"

	"bitbucket.org/ochapman/cache"
)

func TestAdd(t *testing.T) {
	var addTests = []struct {
		name       string
		keyToAdd   interface{}
		keyToGet   interface{}
		expectedOk bool
	}{
		{"string_hit", "string_key", "string_key", true},
		{"int_hit", 1024, 1024, true},
		{"string_hit_error", "string_key_error", "string_key_error_2", false},
	}

	for _, tt := range addTests {
		c := cache.New(0, 0)
		c.Add(tt.keyToAdd, 1234)
		v, ok := c.Get(tt.keyToGet)
		if ok != tt.expectedOk {
			t.Fatalf("%s hit: %v, want %v ", tt.name, ok, !ok)
		} else if ok && v != 1234 {
			t.Fatalf("%s expected to return 1234, but got %v", tt.name, ok, v)
		}
		//t.Logf("%s value: %v", tt.name, 1234)
	}
}

func TestExpire(t *testing.T) {
	var expiredTests = []struct {
		name       string
		key        interface{}
		wait       time.Duration
		expired    time.Duration
		expectedOk bool
	}{
		{"notExpire", "key", time.Second, time.Second * 2, true},
		{"expire", "key", time.Second * 2, time.Second, false},
		{"neverExpire", "key", time.Second, 0, true},
	}
	for _, tt := range expiredTests {
		c := cache.New(0, tt.expired)
		c.Add(tt.key, 1234)
		time.Sleep(tt.wait)
		v, ok := c.Get(tt.key)
		if ok != tt.expectedOk {
			t.Fatalf("%s hit: %v, want: %v", tt.name, ok, tt.expectedOk)
		} else if ok && v != 1234 {
			t.Fatalf("%s expected to return 1234, but got %v", tt.name, v)
		}
		//t.Logf("%s hit: %v, value: %v expired: %v", tt.name, ok, v, tt.expired)
	}
}
