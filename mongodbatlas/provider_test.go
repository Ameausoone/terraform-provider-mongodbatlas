package mongodbatlas

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider
var testAccMongodbAtlasOrgId string
var testAccMongodbAtlasProjectId string

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"mongodbatlas": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("MONGODB_ATLAS_USERNAME"); v == "" {
		t.Fatal("MONGODB_ATLAS_USERNAME must be set for acceptance tests")
		if v := os.Getenv("MONGODB_ATLAS_API_KEY"); v == "" {
			t.Fatal("MONGODB_ATLAS_API_KEY must be set for acceptance tests")
		}
	}
	if testAccMongodbAtlasOrgId = os.Getenv("MONGODB_ATLAS_TESTACC_ORG_ID"); testAccMongodbAtlasOrgId == "" {
		// default org_id
		testAccMongodbAtlasOrgId = "5b71ff2f96e82120d0aaec14"
	}
	if testAccMongodbAtlasProjectId = os.Getenv("MONGODB_ATLAS_TESTACC_PROJECT_ID"); testAccMongodbAtlasProjectId == "" {
		// default org_id
		testAccMongodbAtlasProjectId = "5b71ff2f96e82120d0aaec14"
	}
}
