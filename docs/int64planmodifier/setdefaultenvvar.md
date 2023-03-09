# `SetDefaultEnvVar`

This plan modifier is used to set a default value for a int64 from an environment variable.

## Who to use

```sh
export CAV_VAR_DEFAULT_DISK_SIZE="100"
```

```go
// Schema defines the schema for the resource.
func (r *xResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = schema.Schema{
        (...)
            "disk_size": schema.Int64Attribute{
                Optional:            true,
                MarkdownDescription: "The size of the disk in MB.",
                PlanModifiers: []planmodifier.Int64{
                    fint64planmodifier.SetDefaultEnvVar("CAV_VAR_DEFAULT_DISK_SIZE"),
                },
            },
```
