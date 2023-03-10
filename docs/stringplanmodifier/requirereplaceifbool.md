# `RequireReplaceIfBool`

This plan modifier is used to require a resource to be replaced if a boolean attribute is set to a expected value.

## How to use it

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "name": schema.StringAttribute{
                Optional:            true,
                MarkdownDescription: "A name for ...",
                PlanModifiers: []planmodifier.String{
                    fstringplanmodifier.RequireReplaceIfBool(path.Root("enabled"), true)
                },
            },
            "enabled": schema.BoolAttribute{
                Optional:            true,
                MarkdownDescription: "Enable or disable ...",
            },
```
