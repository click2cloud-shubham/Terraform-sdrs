package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/terraform-providers/terraform-provider-opentelekomcloud/opentelekomcloud"
	
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// func main() {
// 	plugin.Serve(&plugin.ServeOpts{
// 		ProviderFunc: opentelekomcloud.Provider})
// }


func main() {
	plugin.Serve(&plugin.ServeOpts{
			ProviderFunc: func() terraform.ResourceProvider {
					return Provider()
			},
	})
}