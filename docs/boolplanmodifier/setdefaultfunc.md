# `SetDefaultFunc`

This plan modifier is used to set a default value for a boolean using a custom function.

## How to use it

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "enabled": schema.BoolAttribute{
                Optional:            true,
                MarkdownDescription: "Enable or disable ...",
                PlanModifiers: []planmodifier.Bool{
                    fboolplanmodifier.SetDefaultFunc(fboolplanmodifier.DefaultFunc(func(ctx context.Context, req planmodifier.BoolRequest, resp *fboolplanmodifier.DefaultFuncResponse) {
                        if os.Getenv("CAV_VAR_1") == "foo" && os.Getenv("CAV_VAR_2") == "bar" {
                            resp.Value = true
                            return
                        }
                    })),
                },
            },
```
