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

package elasticsearch

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/erda-project/erda-infra/base/logs"
	"github.com/erda-project/erda-infra/base/servicehub"
	writer "github.com/erda-project/erda-infra/pkg/parallel-writer"
	"github.com/olivere/elastic"
)

// WriterConfig .
type WriterConfig struct {
	Type        string `file:"type" desc:"index type"`
	Parallelism uint64 `file:"parallelism" default:"4" desc:"parallelism"`
	Batch       struct {
		Size    uint64        `file:"size" default:"100" desc:"batch size"`
		Timeout time.Duration `file:"timeout" default:"30s" desc:"timeout to flush buffer for batch write"`
	} `file:"batch"`
	Retry int `file:"retry" desc:"retry if fail to write"`
}

// Interface .
type Interface interface {
	URL() string
	Client() *elastic.Client
	NewBatchWriter(*WriterConfig) writer.Writer
}

var clientType = reflect.TypeOf((*elastic.Client)(nil))

type config struct {
	URLs     string `file:"urls" default:"http://localhost:9200" desc:"servers urls"`
	Security bool   `file:"security" default:"false" desc:"enable http basic auth"`
	Username string `file:"username" default:"" desc:"username"`
	Password string `file:"password" default:"" desc:"password"`
}

// provider .
type provider struct {
	Cfg    *config
	Log    logs.Logger
	client *elastic.Client
}

// Init .
func (p *provider) Init(ctx servicehub.Context) error {
	options := []elastic.ClientOptionFunc{
		elastic.SetURL(strings.Split(p.Cfg.URLs, ",")...),
		elastic.SetSniff(false),
	}
	if p.Cfg.Security && (p.Cfg.Username != "" || p.Cfg.Password != "") {
		options = append(options, elastic.SetBasicAuth(p.Cfg.Username, p.Cfg.Password))
	}
	client, err := elastic.NewClient(options...)
	if err != nil {
		return fmt.Errorf("fail to create elasticsearch client: %s", err)
	}
	p.client = client
	return nil
}

// Provide .
func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	if ctx.Type() == clientType || ctx.Service() == "elasticsearch-client" || ctx.Service() == "elastic-client" {
		return p.client
	}
	return &service{
		p:   p,
		log: p.Log.Sub(ctx.Caller()),
	}
}

type service struct {
	p   *provider
	log logs.Logger
}

func (s *service) Client() *elastic.Client { return s.p.client }
func (s *service) URL() string {
	// TODO parse user
	return strings.Split(s.p.Cfg.URLs, ",")[0]
}

func (s *service) NewBatchWriter(c *WriterConfig) writer.Writer {
	return writer.ParallelBatch(func(uint64) writer.Writer {
		return &batchWriter{
			client:        s.p.client,
			log:           s.log,
			typ:           c.Type,
			retry:         c.Retry,
			retryDuration: 3 * time.Second,
			timeout:       fmt.Sprintf("%dms", c.Batch.Timeout.Milliseconds()),
		}
	}, c.Parallelism, c.Batch.Size, c.Batch.Timeout, s.batchWriteError)
}

func (s *service) batchWriteError(err error) error {
	s.log.Errorf("fail to write elasticsearch: %s", err)
	return nil // skip error
}

func init() {
	servicehub.Register("elasticsearch", &servicehub.Spec{
		Services: []string{"elasticsearch", "elasticsearch-client", "elastic-client"},
		Types: []reflect.Type{
			reflect.TypeOf((*Interface)(nil)).Elem(),
			clientType,
		},
		Description: "elasticsearch",
		ConfigFunc:  func() interface{} { return &config{} },
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
