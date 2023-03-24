# Bool Plan Modifiers

Bool plan modifiers are used to modify the plan of a bool attribute.
I will be used into the `PlanModifiers` field of the `schema.BoolAttribute` struct.

## How to use it

```go
import (
    fboolplanmodifier "github.com/FrangipaneTeam/terraform-plugin-framework-planmodifiers/boolplanmodifier"
)
```

## List of Plan Modifiers

- [`SetDefault`](setdefault.md) - Sets a default value for the attribute.
- [`SetDefaultEnvVar`](setdefaultenvvar.md) - Sets a default value for the attribute from an environment variable.
- [`SetDefaultFunc`](setdefaultfunc.md) - Sets a default value for the attribute from a function.

### RequireReplace

- [`RequireReplaceIfBool`](requirereplaceifbool.md) - Forces the resource to be replaced when the specified boolean attribute is changed.
