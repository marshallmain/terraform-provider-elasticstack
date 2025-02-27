package server_host

import (
	"context"

	fleetapi "github.com/elastic/terraform-provider-elasticstack/generated/fleet"
	"github.com/elastic/terraform-provider-elasticstack/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type serverHostModel struct {
	Id      types.String `tfsdk:"id"`
	HostID  types.String `tfsdk:"host_id"`
	Name    types.String `tfsdk:"name"`
	Hosts   types.List   `tfsdk:"hosts"`
	Default types.Bool   `tfsdk:"default"`
}

func (model *serverHostModel) populateFromAPI(ctx context.Context, data *fleetapi.ServerHost) (diags diag.Diagnostics) {
	if data == nil {
		return nil
	}

	model.Id = types.StringValue(data.Id)
	model.HostID = types.StringValue(data.Id)
	model.Name = types.StringValue(data.Name)
	model.Hosts = utils.SliceToListType_String(ctx, data.HostUrls, path.Root("hosts"), &diags)
	model.Default = types.BoolPointerValue(data.IsDefault)

	return
}

func (model serverHostModel) toAPICreateModel(ctx context.Context) (body fleetapi.CreateFleetServerHostJSONRequestBody, diags diag.Diagnostics) {
	body = fleetapi.CreateFleetServerHostJSONRequestBody{
		HostUrls:  utils.ListTypeToSlice_String(ctx, model.Hosts, path.Root("hosts"), &diags),
		Id:        model.HostID.ValueStringPointer(),
		IsDefault: model.Default.ValueBoolPointer(),
		Name:      model.Name.ValueString(),
	}
	return
}

func (model serverHostModel) toAPIUpdateModel(ctx context.Context) (body fleetapi.UpdateFleetServerHostJSONRequestBody, diags diag.Diagnostics) {
	body = fleetapi.UpdateFleetServerHostJSONRequestBody{
		HostUrls:  utils.SliceRef(utils.ListTypeToSlice_String(ctx, model.Hosts, path.Root("hosts"), &diags)),
		IsDefault: model.Default.ValueBoolPointer(),
		Name:      model.Name.ValueStringPointer(),
	}
	return
}
