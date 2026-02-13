# About

Pointless lib

```go
import "github.com/fpersson/mygolib/bar"  // Only imports bar with out any dependcy to mysql
```

```bash
# Only import bar in your code, then run:
go mod tidy
cat go.mod  # Check if MySQL dependency appears

# Or check what's actually used:
go list -m all
```