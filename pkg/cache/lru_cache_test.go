package cache

import (
	"image-previewer/pkg/cache/lru"
	"reflect"
	"sync"
	"testing"
)

func TestLruCache_Get(t *testing.T) {
	type fields struct {
		capacity int
		queue    lru.List
		items    map[string]*lru.ListItem
		mu       *sync.Mutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantVal []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LruCache{
				capacity: tt.fields.capacity,
				queue:    tt.fields.queue,
				items:    tt.fields.items,
				mu:       tt.fields.mu,
			}
			gotVal, err := c.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotVal, tt.wantVal) {
				t.Errorf("Get() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
		})
	}
}

func TestLruCache_Set(t *testing.T) {
	type fields struct {
		capacity int
		queue    lru.List
		items    map[string]*lru.ListItem
		mu       *sync.Mutex
	}
	type args struct {
		key string
		val []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LruCache{
				capacity: tt.fields.capacity,
				queue:    tt.fields.queue,
				items:    tt.fields.items,
				mu:       tt.fields.mu,
			}
			if err := c.Set(tt.args.key, tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
