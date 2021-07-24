package main

import (
	"fmt"
	"os"

	envoy "github.com/open-policy-agent/opa-envoy-plugin/plugin"
	"github.com/open-policy-agent/opa/cmd"
	"github.com/open-policy-agent/opa/runtime"
)

func main() {

	// register opa-envoy plugin
	runtime.RegisterPlugin("envoy.ext_authz.grpc", envoy.Factory{}) // for backwards compatibility
	runtime.RegisterPlugin(envoy.PluginName, envoy.Factory{})

	// register custom decision logger
	runtime.RegisterPlugin(PluginName, Factory{})

	if err := cmd.RootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
