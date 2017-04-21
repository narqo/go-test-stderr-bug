package pkg2

import (
	"fmt"
	"os"
	"testing"
)

func TestPkg2(t *testing.T) {
	fmt.Fprint(os.Stderr, "pkg2 blah")
}
