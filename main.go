package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/vrutkovs/terraform-provider-ignition/v2/ignition"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ignition.Provider})
}
