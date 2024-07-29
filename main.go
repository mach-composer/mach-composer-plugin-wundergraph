package main

import (
	"github.com/mach-composer/mach-composer-plugin-sdk/plugin"
	"github.com/mach-composer/mach-composer-plugin-wundergraph/internal"
)

func main() {
	p := internal.NewWundergraphPlugin()
	plugin.ServePlugin(p)
}
