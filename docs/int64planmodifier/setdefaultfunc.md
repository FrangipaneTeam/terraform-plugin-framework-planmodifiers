# `SetDefaultFunc`

This plan modifier is used to set a default value for a int64 using a custom function.

## How to use it

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "enabled": schema.BoolAttribute{
                Optional:            true,
                MarkdownDescription: "The size of the disk in MB.",
                PlanModifiers: []planmodifier.Bool{
                    fint64planmodifier.SetDefaultFunc(fint64planmodifier.DefaultFunc(func(ctx context.Context, req planmodifier.Int64Request, resp *fint64planmodifier.DefaultFuncResponse) {
                        resp.Value = req.PlanValue * 1024
                    })),
                },
            },
```
