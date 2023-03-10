// Package stringplanmodifier provides a plan modifier for string values.
package stringplanmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// StringChangeFunc is a function that can be used to change a string value.
type StringChangeFunc func(context.Context, planmodifier.StringRequest, *StringChangeFuncResponse)

// StringChangeFuncResponse is the response type for a DefaultFunc.
type StringChangeFuncResponse struct {
	// Diagnostics report errors or warnings related to this logic. An empty
	// or unset slice indicates success, with no warnings or errors generated.
	Diagnostics diag.Diagnostics

	// Value is the value to use by default if the attribute is not configured.
	Value string
}

// setStringChangeFunc
//
// Change the value of the strinf
func setChangeStringFunc(f StringChangeFunc, description, markdownDescription string) planmodifier.String {
	return stringChangeFuncPlanModifier{
		f:                   f,
		description:         description,
		markdownDescription: markdownDescription,
	}
}

// defaultFuncPlanModifier is an plan modifier that sets RequiresReplace
// on the attribute if a given function is true.
type stringChangeFuncPlanModifier struct {
	f                   StringChangeFunc
	description         string
	markdownDescription string
}

// Description returns a human-readable description of the plan modifier.
func (m stringChangeFuncPlanModifier) Description(_ context.Context) string {
	return m.description
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m stringChangeFuncPlanModifier) MarkdownDescription(_ context.Context) string {
	return m.markdownDescription
}

// PlanModifyString implements the plan modification logic.
func (m stringChangeFuncPlanModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}

	funcResp := &StringChangeFuncResponse{}

	m.f(ctx, req, funcResp)

	resp.Diagnostics.Append(funcResp.Diagnostics...)
	resp.PlanValue = types.StringValue(funcResp.Value)
}
