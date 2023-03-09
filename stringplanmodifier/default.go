// Package stringplanmodifier provides a plan modifier for string values.
package stringplanmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// SetDefault
//
// SetDefault returns a plan modifier that sets the plan value to the
// provided value if the following conditions are met:
//
//   - The resource is planned for update.
//   - The plan and state values are not equal.
//   - The plan or state values are not null or known
func SetDefault(str string) planmodifier.String {
	return setDefaultFunc(
		func(_ context.Context, _ planmodifier.StringRequest, resp *DefaultFuncResponse) {
			resp.Value = str
		},
		"Set default value",
		"Set default value",
	)
}
