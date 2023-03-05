package google

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceGoogleCloudFunctionsFunction() *schema.Resource {
	// Generate datasource schema from resource
	dsSchema := datasourceSchemaFromResourceSchema(ResourceCloudFunctionsFunction().Schema)

	// Set 'Required' schema elements
	addRequiredFieldsToSchema(dsSchema, "name")

	// Set 'Optional' schema elements
	addOptionalFieldsToSchema(dsSchema, "project", "region")

	return &schema.Resource{
		Read:   dataSourceGoogleCloudFunctionsFunctionRead,
		Schema: dsSchema,
	}
}

func dataSourceGoogleCloudFunctionsFunctionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := GetProject(d, config)
	if err != nil {
		return err
	}

	region, err := GetRegion(d, config)
	if err != nil {
		return err
	}

	cloudFuncId := &CloudFunctionId{
		Project: project,
		Region:  region,
		Name:    d.Get("name").(string),
	}

	d.SetId(cloudFuncId.CloudFunctionId())

	err = resourceCloudFunctionsRead(d, meta)
	if err != nil {
		return err
	}

	return nil
}
