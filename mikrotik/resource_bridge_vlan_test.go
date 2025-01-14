package mikrotik

import (
	"fmt"
	"testing"

	"github.com/ddelnano/terraform-provider-mikrotik/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestBridgeVlan_basic(t *testing.T) {

	resourceName := "mikrotik_bridge_vlan.testacc"

	createdBridgeVlan := client.BridgeVlan{}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckBridgeVlanDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
						resource "mikrotik_bridge" "default" {
							name = "test_bridge"
						}

						resource "mikrotik_bridge_vlan" "testacc" {
							bridge   = mikrotik_bridge.default.name
							vlan_ids = [10, 15, 18]
							untagged = ["*0"]
						}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccBridgeVlanExists(resourceName, &createdBridgeVlan),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "bridge", "test_bridge"),
					resource.TestCheckResourceAttr(resourceName, "untagged.#", "1"),
				),
			},
			{
				Config: `
				resource "mikrotik_bridge" "default" {
					name = "test_bridge"
					}

					resource "mikrotik_bridge_vlan" "testacc" {
						bridge   = mikrotik_bridge.default.name
						vlan_ids = [10, 15, 18]
						untagged = []
						}
						`,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccBridgeVlanExists(resourceName, &createdBridgeVlan),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "bridge", "test_bridge"),
					resource.TestCheckResourceAttr(resourceName, "untagged.#", "0"),
				),
			},
		},
	})
}

func testAccBridgeVlanExists(resourceID string, record *client.BridgeVlan) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		r, ok := s.RootModule().Resources[resourceID]
		if !ok {
			return fmt.Errorf("resource %q not found in state", resourceID)
		}
		if r.Primary.ID == "" {
			return fmt.Errorf("resource %q has empty primary ID in state", resourceID)
		}
		c := client.NewClient(client.GetConfigFromEnv())
		remoteRecord, err := c.FindBridgeVlan(r.Primary.ID)
		if err != nil {
			return err
		}
		*record = *remoteRecord

		return nil
	}
}

func testAccCheckBridgeVlanDestroy(s *terraform.State) error {
	c := client.NewClient(client.GetConfigFromEnv())
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "mikrotik_bridge_vlan" {
			continue
		}

		remoteRecord, err := c.FindBridgeVlan(rs.Primary.ID)
		if err != nil && !client.IsNotFoundError(err) {
			return fmt.Errorf("expected not found error, got %+#v", err)
		}

		if remoteRecord != nil {
			return fmt.Errorf("bridge vlan %q still exists in remote system", remoteRecord.Id)
		}
	}

	return nil
}
