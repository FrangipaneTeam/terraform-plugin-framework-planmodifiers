# `ToLower`

This plan modifier is used to force the string to be lowercase.

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
                    fstringplanmodifier.ToLower(),
                },
            },
```

```tf title="main.tf"
resource "resource_x" "example" {
  name = "FOO"
}
```

```tf title="terraform.tfstate"
{
  "version": 4,
  "terraform_version": "1.0.0",
  "resources": [
    {
      "mode": "managed",
      "type": "resource_x",
      "name": "example",
      "provider": "provider[\"registry.terraform.io/hashicorp/x\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "id": "example",
            "name": "foo",
          },
        },
      ],
    },
  ],
}
```

