# `SetDefaultEnvVar`

This plan modifier is used to set a default value for a boolean from an environment variable.

## Who to use

```sh
export CAV_VAR_DEFAULT_NAME="true"
```

```go
// Schema defines the schema for the resource.
func (r *vappResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "enabled": schema.BoolAttribute{
                Optional:            true,
                MarkdownDescription: "",
                PlanModifiers: []planmodifier.Bool{
                    fboolplanmodifier.SetDefaultEnvVar("CAV_VAR_DEFAULT_NAME"),
                },
            },
```
