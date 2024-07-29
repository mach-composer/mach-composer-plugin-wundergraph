package plugin

import (
	"github.com/mach-composer/mach-composer-plugin-sdk/plugin"

	"github.com/mach-composer/mach-composer-plugin-wundergraph/internal"
)

// Serve serves the plugin
func Serve() {
	p := internal.NewWundergraphPlugin()
	plugin.ServePlugin(p)
}
