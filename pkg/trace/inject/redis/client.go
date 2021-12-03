// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package redis

import (
	"strconv"
	"strings"

	"github.com/go-redis/redis"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"

	"github.com/erda-project/erda-infra/pkg/trace/inject/hook"
)

// Client .
type Client interface {
	WrapProcess(fn func(oldProcess func(redis.Cmder) error) func(redis.Cmder) error)
	WrapProcessPipeline(fn func(oldProcess func([]redis.Cmder) error) func([]redis.Cmder) error)
}

// Wrap .
func Wrap(client Client, opts ...Option) {
	if client == nil {
		return
	}
	cfg := newConfig("redis", opts...)
	client.WrapProcess(newProcessWrapper(cfg))
	client.WrapProcessPipeline(newProcessPipeline(cfg))
}

//go:noinline
func originalNewClient(opts *redis.Options) *redis.Client {
	return redis.NewClient(opts)
}

//go:noinline
func originalNewFailoverClient(opts *redis.FailoverOptions) *redis.Client {
	return redis.NewFailoverClient(opts)
}

//go:noinline
func originalNewClusterClient(opts *redis.ClusterOptions) *redis.ClusterClient {
	return redis.NewClusterClient(opts)
}

//go:noinline
func wrappedNewClient(opts *redis.Options) *redis.Client {
	client := originalNewClient(opts)
	Wrap(client, WithAttributes(
		attribute.Key("db.host").String(opts.Addr),
		attribute.Key("redis.client.type").String("client"),
		semconv.DBNameKey.String(strconv.FormatInt(int64(opts.DB), 10)),
	))
	return client
}

//go:noinline
func wrappedNewFailoverClient(opts *redis.FailoverOptions) *redis.Client {
	client := originalNewFailoverClient(opts)
	Wrap(client, WithAttributes(
		attribute.Key("db.host").String(strings.Join(opts.SentinelAddrs, ",")),
		attribute.Key("redis.master.name").String(opts.MasterName),
		attribute.Key("redis.client.type").String("sentinel"),
		semconv.DBNameKey.String(strconv.FormatInt(int64(opts.DB), 10)),
	))
	return client
}

//go:noinline
func wrappedNewClusterClient(opts *redis.ClusterOptions) *redis.ClusterClient {
	client := originalNewClusterClient(opts)
	Wrap(client, WithAttributes(
		attribute.Key("db.host").String(strings.Join(opts.Addrs, ",")),
		attribute.Key("redis.client.type").String("cluster"),
	))
	return client
}

func init() {
	hook.Hook(redis.NewClient, wrappedNewClient, originalNewClient)
	hook.Hook(redis.NewFailoverClient, wrappedNewFailoverClient, originalNewFailoverClient)
	hook.Hook(redis.NewClusterClient, wrappedNewClusterClient, originalNewClusterClient)
}
