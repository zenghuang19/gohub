// Copyright 2021 gotomicro
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"errors"
	"sync"
	"time"
)

var (
	errKeyNotFound = errors.New("toy-cache: key 不存在")
	errKeyExpired  = errors.New("toy-cache: key 已经过期")
)

type BuildinMapCache struct {
	lock      sync.RWMutex
	data      map[string]*item
	close     chan struct{}
	onEvicted func(key string, val any)
}

func NewBuildinMapCache() *BuildinMapCache {
	res := &BuildinMapCache{
		data: make(map[string]*item),
	}
	res.checkCycle()
	return res
}

func (b *BuildinMapCache) Get(ctx context.Context, key string) (any, error) {
	b.lock.RLock()
	val, ok := b.data[key]
	b.lock.RUnlock()
	if !ok {
		return nil, errKeyNotFound
	}

	now := time.Now()
	if val.deadlineBefore(now) {
		b.lock.Lock()
		defer b.lock.Unlock()
		val, ok = b.data[key]
		if !ok {
			return nil, errKeyNotFound
		}
		if val.deadlineBefore(now) {
			b.delete(key)
			return nil, errKeyExpired
		}
	}
	return val.val, nil
}

func (b *BuildinMapCache) Set(ctx context.Context, key string, val any, expiration time.Duration) error {
	b.lock.Lock()
	defer b.lock.Unlock()
	var dl time.Time
	if expiration > 0 {
		dl = time.Now().Add(expiration)
	}
	b.data[key] = &item{
		val:      val,
		deadline: dl,
	}
	return nil
}

// func (b *BuildinMapCache) Set(ctx context.Context, key string, val any, expiration time.Duration) error {
//  b.lock.Lock()
//  b.data[key] = val
//  b.lock.Unlock()
//  if expiration > 0 {
//      time.AfterFunc(expiration, func() {
//          delete(b.data, key)
//      })
//  }
//
//  return nil
// }

func (b *BuildinMapCache) Delete(ctx context.Context, key string) error {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.delete(key)
	return nil
}

func (b *BuildinMapCache) checkCycle() {
	// 这个 10s 可以做成配置的
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case now := <-ticker.C:
				b.lock.Lock()
				for key, val := range b.data {
					// 设置了过期时间，并且已经过期
					if !val.deadline.IsZero() &&
						val.deadline.Before(now) {
						b.delete(key)
					}
				}
				b.lock.Unlock()
			case <-b.close:
				close(b.close)
				return
			}
		}
	}()
}

func (b *BuildinMapCache) delete(key string) {
	val, ok := b.data[key]
	if ok {
		delete(b.data, key)
		b.onEvicted(key, val.val)
	}
}

func (b *BuildinMapCache) Close() error {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.close <- struct{}{}
	for key, val := range b.data {
		b.onEvicted(key, val.val)
	}
	b.data = nil
	return nil
}

type item struct {
	val      any
	deadline time.Time
}

func (i *item) deadlineBefore(t time.Time) bool {
	return !i.deadline.IsZero() && i.deadline.Before(t)
}
