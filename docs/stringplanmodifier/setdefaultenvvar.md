# `SetDefaultEnvVar`

This plan modifier is used to set a default value for a string from an environment variable.

## How to use it

```sh
export CAV_VAR_DEFAULT_NAME="default-name"
```

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "name": schema.StringAttribute{
                Optional:            true,
                MarkdownDescription: "A name for ...",
                PlanModifiers: []planmodifier.String{
                    fstringplanmodifier.SetDefaultEnvVar("CAV_VAR_DEFAULT_NAME"),
                },
            },
```
