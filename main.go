// This is a no-op go program
// It exists as a HACK to avoid `go install ./...` errorsi
// for the upstream pre-commit hook
package noop

import (
	"fmt"
)

func main() {
	fmt.Print("no-op")
}
