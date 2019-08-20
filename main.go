package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/sean-eastbury/terraform-provider-hyperv/hyperv"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: hyperv.Provider,
	})
}
