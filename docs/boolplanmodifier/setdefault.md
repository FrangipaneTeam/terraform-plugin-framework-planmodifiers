# `SetDefault`

This plan modifier is used to set a default value for a boolean attribute.

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
                    fboolplanmodifier.SetDefault(true),
                },
            },
```
