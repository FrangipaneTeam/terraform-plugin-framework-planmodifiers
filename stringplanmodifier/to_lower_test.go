// Package stringplanmodifier provides a plan modifier for string values.
package stringplanmodifier_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/FrangipaneTeam/terraform-plugin-framework-planmodifiers/stringplanmodifier"
)

func TestToLowerPlanModifyString(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         types.String
		exceptedVal types.String
		expectError bool
	}

	tests := map[string]testCase{
		"unknown String": {
			val:         types.StringUnknown(),
			exceptedVal: types.StringNull(),
		},
		"null String": {
			val:         types.StringNull(),
			exceptedVal: types.StringNull(),
		},
		"valid String": {
			val:         types.StringValue("TEST"),
			exceptedVal: types.StringValue("test"),
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			request := planmodifier.StringRequest{
				Path:           path.Root("test"),
				PathExpression: path.MatchRoot("test"),
				ConfigValue:    test.val,
			}

			resp := &planmodifier.StringResponse{}
			stringplanmodifier.ToLower().PlanModifyString(context.Background(), request, resp)

			if diff := cmp.Diff(test.exceptedVal, resp.PlanValue); diff != "" && !test.expectError {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
