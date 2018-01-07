package mongodbatlas

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MONGODB_ATLAS_USERNAME", ""),
				Description: "MongoDB Atlas username",
			},
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MONGODB_ATLAS_API_KEY", ""),
				Description: "MongoDB Atlas API Key",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"mongodbatlas_cluster": resourceCluster(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		AtlasUsername: d.Get("username").(string),
		AtlasAPIKey:   d.Get("api_key").(string),
	}

	client := config.NewClient()

	return client, nil
}
