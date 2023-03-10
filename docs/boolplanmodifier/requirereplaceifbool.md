# `RequireReplaceIfBool`

This plan modifier is used to require a resource to be replaced if a boolean attribute is set to a expected value.

## How to use it

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "enabled": schema.StringAttribute{
                Optional:            true,
                MarkdownDescription: "Enable or disable ...",
                PlanModifiers: []planmodifier.String{
                    fboolplanmodifier.RequireReplaceIfBool(path.Root("force"), true)
                },
            "force": schema.Int64Attribute{
                Optional:            true,
                MarkdownDescription: "Force the resource ...",
            },
```
