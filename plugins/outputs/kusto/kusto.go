package kusto

import (
	"fmt"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/outputs"
)

type Simple struct {
	Ok  bool            `toml:"ok"`
	Log telegraf.Logger `toml:"-"`
}

func (s *Simple) Description() string {
	return "a demo output"
}

func (s *Simple) SampleConfig() string {
	return `
  ok = true
`
}

// Init is for setup, and validating config.
func (s *Simple) Init() error {
	return nil
}

func (s *Simple) Connect() error {
	// Make any connection required here
	return nil
}

func (s *Simple) Close() error {
	// Close any connections here.
	// Write will not be called once Close is called, so there is no need to synchronize.
	return nil
}

// Write should write immediately to the output, and not buffer writes
// (Telegraf manages the buffer for you). Returning an error will fail this
// batch of writes and the entire batch will be retried automatically.
func (s *Simple) Write(metrics []telegraf.Metric) error {
	for _, metric := range metrics {
		// write `metric` to the output sink here
		fmt.Println(metric)
	}
	return nil
}

func init() {
	outputs.Add("kusto", func() telegraf.Output { return &Simple{} })
}
