# `SetDefaultFunc`

This plan modifier is used to set a default value for a string using a custom function.

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
                    fstringplanmodifier.SetDefaultFunc(fstringplanmodifier.DefaultFunc(func(ctx context.Context, req planmodifier.StringRequest, resp *fstringplanmodifier.DefaultFuncResponse) {
                        if strings.Contains(req.PlanValue, "foo") {
                            resp.Value = "bar"
                            return
                        }
                    })),
                },
            },
```
