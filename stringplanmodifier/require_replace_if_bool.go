package stringplanmodifier

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

/*
RequireReplaceIfBool

returns a plan modifier that requires replacement
if the attribute value is equal to the excepted value.
*/
func RequireReplaceIfBool(path path.Path, exceptedValue bool) planmodifier.String {
	description := fmt.Sprintf("Attribute require replacement if `%s` is `%v`", path.String(), exceptedValue)
	return stringplanmodifier.RequiresReplaceIf(stringplanmodifier.RequiresReplaceIfFunc(func(ctx context.Context, req planmodifier.StringRequest, resp *stringplanmodifier.RequiresReplaceIfFuncResponse) {
		boolValue := &types.Bool{}

		resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path, boolValue)...)
		if resp.Diagnostics.HasError() {
			return
		}

		if boolValue.ValueBool() == exceptedValue {
			resp.RequiresReplace = true
		}
	}), description, description)
}
