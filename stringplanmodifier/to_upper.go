// Package stringplanmodifier provides a plan modifier for string values.
package stringplanmodifier

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func ToUpper() planmodifier.String {
	return setChangeStringFunc(
		func(_ context.Context, req planmodifier.StringRequest, resp *StringChangeFuncResponse) {
			resp.Value = types.StringValue(strings.ToUpper(req.ConfigValue.ValueString()))
		},
		"Force to upper case",
		"Force to upper case",
	)
}
