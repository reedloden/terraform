package main

import (
	"github.com/hashicorp/terraform/builtin/providers/cloudstack"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cloudstack.Provider,
	})
}
