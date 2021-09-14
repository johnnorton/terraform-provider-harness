package provider

import (
	"context"

	"github.com/harness-io/harness-go-sdk/harness/api"
	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAWSLambdaService() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource for creating an AWS Lambda service",
		CreateContext: resourceAWSLambdaServiceCreateOrUpdate,
		ReadContext:   resourceAWSLambdaServiceRead,
		UpdateContext: resourceAWSLambdaServiceCreateOrUpdate,
		DeleteContext: resourceServiceDelete,
		Schema:        commonServiceSchema(),
		Importer: &schema.ResourceImporter{
			State: serviceStateImporter,
		},
	}
}

func resourceAWSLambdaServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	svcId := d.Get("id").(string)
	appId := d.Get("app_id").(string)

	var svc *cac.Service
	var err error

	if svc, err = c.ConfigAsCode().GetServiceById(appId, svcId); err != nil {
		return diag.FromErr(err)
	} else if svc == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	return readServiceAwsLambda(d, svc)
}

func readServiceAwsLambda(d *schema.ResourceData, svc *cac.Service) diag.Diagnostics {
	d.SetId(svc.Id)
	d.Set("name", svc.Name)
	d.Set("app_id", svc.ApplicationId)
	d.Set("description", svc.Description)

	if vars := flattenServiceVariables(svc.ConfigVariables); len(vars) > 0 {
		d.Set("variable", vars)
	}

	return nil
}

func resourceAWSLambdaServiceCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	var input *cac.Service
	var err error

	if d.IsNewResource() {
		input = cac.NewEntity(cac.ObjectTypes.Service).(*cac.Service)
	} else {
		if input, err = c.ConfigAsCode().GetServiceById(d.Get("app_id").(string), d.Id()); err != nil {
			return diag.FromErr(err)
		} else if input == nil {
			d.SetId("")
			d.MarkNewResource()
			return nil
		}
	}

	// Setup the object to be created
	input.Name = d.Get("name").(string)
	input.ArtifactType = cac.ArtifactTypes.AWSLambda
	input.DeploymentType = cac.DeploymentTypes.AWSLambda
	input.ApplicationId = d.Get("app_id").(string)
	input.Description = d.Get("description").(string)

	if vars := d.Get("variable"); vars != nil {
		input.ConfigVariables = expandServiceVariables(vars.(*schema.Set).List())
	}

	// Create Service
	newSvc, err := c.ConfigAsCode().UpsertService(input)
	if err != nil {
		return diag.FromErr(err)
	}

	return readServiceAwsLambda(d, newSvc)
}

func resourceAWSLambdaServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*api.Client)

	// Setup the object to create
	svcInput := &cac.Service{
		Name:           d.Get("name").(string),
		ArtifactType:   cac.ArtifactTypes.AWSLambda,
		DeploymentType: cac.DeploymentTypes.AWSLambda,
		ApplicationId:  d.Get("app_id").(string),
		Description:    d.Get("description").(string),
	}

	if vars := d.Get("variable"); vars != nil {
		svcInput.ConfigVariables = expandServiceVariables(vars.(*schema.Set).List())
	}

	// Create Service
	newSvc, err := c.ConfigAsCode().UpsertService(svcInput)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(newSvc.Id)

	return nil
}
