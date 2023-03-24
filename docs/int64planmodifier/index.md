# Int64 Plan Modifiers

BoInt64ol plan modifiers are used to modify the plan of a int64 attribute.
I will be used into the `PlanModifiers` field of the `schema.Int64Attribute` struct.

## How to use it

```go
import (
    fint64planmodifier "github.com/FrangipaneTeam/terraform-plugin-framework-planmodifiers/int64planmodifier"
)
```

## List of Plan Modifiers

- [`SetDefault`](setdefault.md) - Sets a default value for the attribute.
- [`SetDefaultEnvVar`](setdefaultenvvar.md) - Sets a default value for the attribute from an environment variable.
- [`SetDefaultFunc`](setdefaultfunc.md) - Sets a default value for the attribute from a function.

### RequireReplace

- [`RequireReplaceIfBool`](requirereplaceifbool.md) - Forces the resource to be replaced when the specified boolean attribute is changed.
