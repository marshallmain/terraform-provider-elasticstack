package detections

import (
	"context"

	"github.com/elastic/terraform-provider-elasticstack/generated/detections"
	"github.com/elastic/terraform-provider-elasticstack/internal/clients"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = (*Resource)(nil)

func NewEqlRuleResource() resource.Resource {
	return &Resource{}
}

type Resource struct {
	client *clients.ApiClient
}

type commonParams struct {
	Author []types.String `tfsdk:"author"`
}

type eqlRuleResourceModel struct {
	Id           types.String `tfsdk:"id"`
	SpaceID      types.String `tfsdk:"space_id"`
	Name         types.String `tfsdk:"name"`
	Description  types.String `tfsdk:"description"`
	RiskScore    types.Int32  `tfsdk:"risk_score"`
	Severity     types.String `tfsdk:"severity"`
	Query        types.String `tfsdk:"query"`
	CommonParams commonParams `tfsdk:"common_params"`
}

func (r *Resource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	client, diags := clients.ConvertProviderData(request.ProviderData)
	response.Diagnostics.Append(diags...)
	r.client = client
}

func (r *Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_kibana_eql_rule"
}

func (r *Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"space_id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"description": schema.StringAttribute{
				Required: true,
			},
			"risk_score": schema.Int32Attribute{
				Required: true,
			},
			"severity": schema.StringAttribute{
				Required: true,
			},
			"query": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	detectionsClient, err := r.client.GetDetectionsClient()
	if err != nil {
		resp.Diagnostics.AddError("unable to get detections client", err.Error())
		return
	}

	var data eqlRuleResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create API call logic

	apiData := data.ToCreateRequest(ctx)
	authCtx := r.client.SetDetectionsAuthContext(ctx)

	respModel, res, err := detectionsClient.CreateRule(authCtx, data.SpaceID.ValueString()).RuleCreateProps(detections.RuleCreateProps{EqlRuleCreateProps: &apiData}).KbnXsrf("true").ElasticApiVersion("2023-10-31").Execute()

	if err != nil {
		resp.Diagnostics.AddError("Failed to create eql rule", err.Error())
		return
	}
	defer res.Body.Close()

	data.FromApiResponse(respModel)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	detectionsClient, err := r.client.GetDetectionsClient()
	if err != nil {
		resp.Diagnostics.AddError("unable to get detections client", err.Error())
		return
	}

	var data eqlRuleResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	authCtx := r.client.SetDetectionsAuthContext(ctx)

	// Read API call logic
	respModel, res, err := detectionsClient.ReadRule(authCtx, data.SpaceID.ValueString()).Id(data.Id.ValueString()).KbnXsrf("true").ElasticApiVersion("2023-10-31").Execute()

	if err != nil && res == nil {
		resp.Diagnostics.AddError("Failed to read eql rule", err.Error())
		return
	}
	defer res.Body.Close()

	data.FromApiResponse(respModel)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	detectionsClient, err := r.client.GetDetectionsClient()
	if err != nil {
		resp.Diagnostics.AddError("unable to get detections client", err.Error())
		return
	}

	var data eqlRuleResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiData := data.ToUpdateRequest(ctx)
	authCtx := r.client.SetDetectionsAuthContext(ctx)

	// Update API call logic
	respModel, res, err := detectionsClient.UpdateRule(authCtx, data.SpaceID.ValueString()).RuleUpdateProps(detections.RuleUpdateProps{EqlRuleUpdateProps: &apiData}).KbnXsrf("true").ElasticApiVersion("2023-10-31").Execute()

	if err != nil && res == nil {
		resp.Diagnostics.AddError("Failed to read eql rule", err.Error())
		return
	}
	defer res.Body.Close()

	data.FromApiResponse(respModel)
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	detectionsClient, err := r.client.GetDetectionsClient()
	if err != nil {
		resp.Diagnostics.AddError("unable to get detections client", err.Error())
		return
	}

	var data eqlRuleResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	authCtx := r.client.SetDetectionsAuthContext(ctx)

	// Delete API call logic
	_, res, err := detectionsClient.DeleteRule(authCtx, data.SpaceID.ValueString()).Id(data.Id.ValueString()).KbnXsrf("true").ElasticApiVersion("2023-10-31").Execute()

	if err != nil && res == nil {
		resp.Diagnostics.AddError("Failed to delete eql rule", err.Error())
		return
	}
	defer res.Body.Close()
}

func (m eqlRuleResourceModel) ToCreateRequest(ctx context.Context) detections.EqlRuleCreateProps {
	apiModel := detections.EqlRuleCreateProps{
		Name:        m.Name.ValueString(),
		Description: m.Description.ValueString(),
		RiskScore:   m.RiskScore.ValueInt32(),
		Severity:    detections.Severity(m.Severity.ValueString()),
		Query:       m.Query.ValueString(),
		Type:        "eql",
		Language:    "eql",
	}

	return apiModel
}

func (m eqlRuleResourceModel) ToUpdateRequest(ctx context.Context) detections.EqlRuleUpdateProps {
	apiModel := detections.EqlRuleUpdateProps{
		Name:        m.Name.ValueString(),
		Description: m.Description.ValueString(),
		RiskScore:   m.RiskScore.ValueInt32(),
		Severity:    detections.Severity(m.Severity.ValueString()),
		Query:       m.Query.ValueString(),
		Type:        "eql",
		Language:    "eql",
	}

	return apiModel
}

func (m *eqlRuleResourceModel) FromApiResponse(apiResp *detections.RuleResponse) {
	m.Id = types.StringValue(apiResp.EqlRule.Id)
	m.Name = types.StringValue(apiResp.EqlRule.Name)
	m.Description = types.StringValue(apiResp.EqlRule.Description)
	m.RiskScore = types.Int32Value(apiResp.EqlRule.RiskScore)
	m.Severity = types.StringValue(string(apiResp.EqlRule.Severity))
	m.Query = types.StringValue(apiResp.EqlRule.Query)
}
