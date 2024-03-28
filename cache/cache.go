// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cache

import (
	"time"

	"github.com/FishGoddess/cachego"
	"github.com/FishGoddess/logit"
)

func New(name string, opts ...cachego.Option) cachego.Cache {
	defaultOpts := []cachego.Option{
		cachego.WithCacheName(name),
		cachego.WithShardings(16),    // 16 个分片
		cachego.WithMaxEntries(5000), // 单个分片的最大数量，总共是 8w 个
		cachego.WithGC(10 * time.Minute),
		cachego.WithRecordLoad(false),
		cachego.WithReportGC(reportCacheGC),
	}

	opts = append(defaultOpts, opts...)
	cache, _ := cachego.NewCacheWithReport(opts...)

	return cache
}

func NewLRU(name string, maxEntries int, opts ...cachego.Option) cachego.Cache {
	defaultOpts := []cachego.Option{
		cachego.WithLRU(maxEntries / 2), cachego.WithShardings(2),
	}

	opts = append(defaultOpts, opts...)
	return New(name, opts...)
}

func NewLFU(name string, maxEntries int, opts ...cachego.Option) cachego.Cache {
	defaultOpts := []cachego.Option{
		cachego.WithLFU(maxEntries / 2), cachego.WithShardings(2),
	}

	opts = append(defaultOpts, opts...)
	return New(name, opts...)
}

func reportCacheGC(reporter *cachego.Reporter, cost time.Duration, cleans int) {
	size := reporter.CacheSize()
	logit.Info("report cache gc", "name", reporter.CacheName(), "type", reporter.CacheType(), "cost", cost, "cleans", cleans, "current_size", size, "hit_rate", reporter.HitRate())
}
