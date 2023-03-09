// Package int64planmodifier provides a plan modifier for int64 values.
package int64planmodifier

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
func SetDefault(i int64) planmodifier.Int64 {
	return setDefaultFunc(
		func(_ context.Context, _ planmodifier.Int64Request, resp *DefaultFuncResponse) {
			resp.Value = i
		},
		"Set default value",
		"Set default value",
	)
}
