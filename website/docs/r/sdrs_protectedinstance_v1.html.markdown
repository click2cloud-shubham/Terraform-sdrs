---
layout: "opentelekomcloud"
page_title: "OpenTelekomCloud: opentelekomcloud_sdrs_protectedinstance_v1"
sidebar_current: "docs-opentelekomcloud-resource-sdrs-protectedinstance-v1"
description: |-
  Manages a V1 SDRS protected instance resource within OpenTelekomCloud.
---

# opentelekomcloud_sdrs_protectedinstance_v1

Manages a SDRS protected instance resource within OpenTelekomCloud.

## Example Usage

```hcl
resource "opentelekomcloud_sdrs_protectedinstance_v1" "group_1" {
	server_group_id = "group_1"
	server_id = "test description"
	name = "group_1"
	description = "test description"
    primary_subnet_id = "Specifies the network ID of the subnet "
    primary_ip_address = "Specifies the IP address of the primary"
    flavorRef= "Specifies the flavor ID"
}

```

## Argument Reference

The following arguments are supported:

* `GroupID` - (Required) Specifies the ID of the protection group where a protected instance is added..

* `server_id` - (Required) Specifies the ID of the production site server.

* `name` - (Required) The name of a protected instance.

* `description` - (Optional) The description of a protection instance. 

* `primary_subnet_id` - (Optional) Specifies the network ID of the subnet.

* `primary_ip_address` - (Optional) Specifies the IP address of the primary NIC on the DR site server.
                                    
* `flavorRef` - (Optional) Same name as the production site server.

## Attributes Reference

The following attributes are exported:

* `id` -  ID of the protected instance.

