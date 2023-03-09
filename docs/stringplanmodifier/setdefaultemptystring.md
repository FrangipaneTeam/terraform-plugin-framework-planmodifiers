# `SetDefaultEmptyString`

This plan modifier is used to set a default empty value for a string.

## Who to use

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "name": schema.StringAttribute{
                Optional:            true,
                MarkdownDescription: "A name for ...",
                PlanModifiers: []planmodifier.String{
                    fstringplanmodifier.SetDefaultEmptyString(),
                },
            },
```
