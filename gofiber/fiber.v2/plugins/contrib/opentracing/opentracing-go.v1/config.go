package opentracing

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/gofiber/fiber.v2"
)

const (
	enabled = fiber.PluginsRoot + ".opentracing.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable opentracing middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
