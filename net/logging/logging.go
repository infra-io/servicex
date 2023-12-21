// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package logging

import "context"

type RequestResolver func(ctx context.Context, request any) []any

type Config struct {
	WithTraceID       bool
	WithServiceMethod bool
	Resolvers         []RequestResolver
}

func NewDefaultConfig() *Config {
	conf := &Config{
		WithTraceID:       false,
		WithServiceMethod: false,
	}

	return conf
}

type Option func(conf *Config)

func (o Option) ApplyTo(conf *Config) {
	o(conf)
}

func WithTraceID() Option {
	return func(conf *Config) {
		conf.WithTraceID = true
	}
}

func WithServiceMethod() Option {
	return func(conf *Config) {
		conf.WithServiceMethod = true
	}
}

func WithRequestResolver(resolver RequestResolver) Option {
	return func(conf *Config) {
		conf.Resolvers = append(conf.Resolvers, resolver)
	}
}
