package pkg1

import (
	"fmt"
	"os"
	"testing"
)

func TestPkg1(t *testing.T) {
	fmt.Fprintln(os.Stderr, "pkg1 blah")
}
