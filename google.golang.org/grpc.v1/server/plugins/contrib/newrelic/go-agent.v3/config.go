package newrelic

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/google.golang.org/grpc.v1/server"
)

const (
	root    = server.PluginsRoot + ".newrelic"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable newrelic")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
