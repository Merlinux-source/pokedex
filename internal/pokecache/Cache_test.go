package pokecache

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestCache_Add(t *testing.T) {
	type fields struct {
		mu       sync.RWMutex
		duration time.Duration
		entries  map[string]cacheEntry
	}
	type args struct {
		key   string
		value []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "add", fields: fields{mu: sync.RWMutex{}, duration: time.Second, entries: map[string]cacheEntry{}}, args: args{key: "key", value: []byte("value")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				mu:       tt.fields.mu,
				duration: tt.fields.duration,
				entries:  tt.fields.entries,
			}
			c.Add(tt.args.key, tt.args.value)
		})
	}
}

func TestCache_Get(t *testing.T) {
	type fields struct {
		mu       sync.RWMutex
		duration time.Duration
		entries  map[string]cacheEntry
	}
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue []byte
		wantOk    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				mu:       tt.fields.mu,
				duration: tt.fields.duration,
				entries:  tt.fields.entries,
			}
			gotValue, gotOk := c.Get(tt.args.key)
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("Get() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Get() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestCache_readLoop(t *testing.T) {
	type fields struct {
		mu       sync.RWMutex
		duration time.Duration
		entries  map[string]cacheEntry
	}
	type args struct {
		interval time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				mu:       tt.fields.mu,
				duration: tt.fields.duration,
				entries:  tt.fields.entries,
			}
			c.readLoop(tt.args.interval)
		})
	}
}

func TestNewCache(t *testing.T) {
	type args struct {
		duration time.Duration
	}
	tests := []struct {
		name            string
		args            args
		wantReturnCache *Cache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotReturnCache := NewCache(tt.args.duration); !reflect.DeepEqual(gotReturnCache, tt.wantReturnCache) {
				t.Errorf("NewCache() = %v, want %v", gotReturnCache, tt.wantReturnCache)
			}
		})
	}
}

func Test_cacheEntry_CreatedAt(t *testing.T) {
	type fields struct {
		createdAt time.Time
		value     []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &cacheEntry{
				createdAt: tt.fields.createdAt,
				value:     tt.fields.value,
			}
			if got := e.CreatedAt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreatedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cacheEntry_Value(t *testing.T) {
	type fields struct {
		createdAt time.Time
		value     []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &cacheEntry{
				createdAt: tt.fields.createdAt,
				value:     tt.fields.value,
			}
			if got := e.Value(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
