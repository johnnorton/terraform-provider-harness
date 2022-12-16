package connector_test

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceConnectorKubernetesCloudCost(t *testing.T) {
	var (
		name         = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
		resourceName = "data.harness_platform_connector_kubernetes_cloud_cost.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConnector_kubernetesCloudCost(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "id", name),
					resource.TestCheckResourceAttr(resourceName, "identifier", name),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "connector_ref", name+"s"),
				),
			},
		},
	})
}

func testAccDataSourceConnector_kubernetesCloudCost(name string) string {
	return fmt.Sprintf(`
	resource "harness_platform_connector_kubernetes" "test" {
		identifier = "%[1]ss"
		name = "%[1]ss"
		description = "test"
		tags = ["foo:bar"]

		inherit_from_delegate {
			delegate_selectors = ["harness-delegate"]
		}
	}

		resource "harness_platform_connector_kubernetes_cloud_cost" "test" {
			identifier = "%[1]s"
			name = "%[1]s"
			description = "test"
			tags = ["foo:bar"]

			features_enabled = ["VISIBILITY", "OPTIMIZATION"]
			connector_ref = harness_platform_connector_kubernetes.test.id
		}

		data "harness_platform_connector_kubernetes_cloud_cost" "test" {
			identifier = harness_platform_connector_kubernetes_cloud_cost.test.identifier
		}
	`, name)
}
