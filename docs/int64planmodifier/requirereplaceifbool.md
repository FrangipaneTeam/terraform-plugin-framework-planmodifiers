# `RequireReplaceIfBool`

This plan modifier is used to require a resource to be replaced if a boolean attribute is set to a expected value.

## Who to use

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "disk_size": schema.Int64Attribute{
                Optional:            true,
                MarkdownDescription: "The size of the disk in MB.",
                PlanModifiers: []planmodifier.Int64{
                    fint64planmodifier.RequireReplaceIfBool(path.Root("enabled"), true)
                },
            },
            "enabled": schema.BoolAttribute{
                Optional:            true,
                MarkdownDescription: "Enable or disable ...",
            },
```
