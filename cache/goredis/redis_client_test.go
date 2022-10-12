package goredis

import "testing"

func TestRedisClient(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"set key"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RedisClient()
		})
	}
}
