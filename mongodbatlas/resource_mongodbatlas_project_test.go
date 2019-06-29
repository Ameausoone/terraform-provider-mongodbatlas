package mongodbatlas

import (
	"fmt"
	"testing"

	ma "github.com/akshaykarle/go-mongodbatlas/mongodbatlas"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccMongodbatlasProject_basic(t *testing.T) {
	projectName := "testAcc"
	resourceName := "mongodbatlas_project.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMongodbatlasProject(projectName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "org_id"),
					resource.TestCheckResourceAttrSet(resourceName, "created"),
					resource.TestCheckResourceAttrSet(resourceName, "cluster_count"),
					resource.TestCheckResourceAttr(resourceName, "name", projectName),
				),
			},
		},
	})
}

func TestAccProject_importBasic(t *testing.T) {
	resourceName := "mongodbatlas_project.test"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMongodbatlasProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMongodbatlasProject("test"),
			},
			{
				ResourceName:      resourceName,
				ImportStateId:     testAccMongodbAtlasProjectId,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccMongodbatlasProject(projectName string) string {
	return fmt.Sprintf(`resource "mongodbatlas_project" "test" {
  org_id = "%s"
  name = "%s"
}`, testAccMongodbAtlasOrgId, projectName)
}

func testAccCheckMongodbatlasProjectDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*ma.Client)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "mongodbatlas_project" {
			continue
		}

		projects, _, err := client.Projects.List()

		if err == nil {
			if len(projects) != 0 {
				return fmt.Errorf("projects %q still exists", rs.Primary.ID)
			}
		}

		// Verify the error
		if err != nil {
			return fmt.Errorf("Error listing MongoDB Projects: %s", err)
		}
	}

	return nil
}
