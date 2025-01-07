package debug_test

import (
	"testing"

	"github.com/ak5w/gutil/debug"
)

func TestPrint(t *testing.T) {
	a := 1
	b := "abc"
	c := 123.456

	s := []int{1, 2, 3}
	debug.Print(a)
	debug.Print(b)
	debug.Print(c)
	debug.Print(s)
	debug.Print(a, b, c, s)
}
