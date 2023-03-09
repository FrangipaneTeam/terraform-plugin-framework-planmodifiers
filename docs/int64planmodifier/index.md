# StringPlanModifiers

BoInt64ol plan modifiers are used to modify the plan of a int64 attribute.
I will be used into the `PlanModifiers` field of the `schema.Int64Attribute` struct.

## Who to use

```go
import (
    fint64planmodifier "github.com/FrangipaneTeam/terraform-plugin-framework-planmodifiers/int64planmodifier"
)
```

## List of Plan Modifiers

- [`SetDefault`](setdefault.md) - Sets a default value for the attribute.
- [`SetDefaultEnvVar`](setdefaultenvvar.md) - Sets a default value for the attribute from an environment variable.
- [`SetDefaultFunc`](setdefaultfunc.md) - Sets a default value for the attribute from a function.
