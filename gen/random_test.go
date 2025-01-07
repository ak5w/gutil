package gen_test

import (
	"testing"

	"github.com/ak5w/gutil/gen"
)

func isLowerChar(c uint8) bool {
	return isDigital(c) || (c >= 'a' && c <= 'z')
}

func isDigital(c uint8) bool {
	return c >= '0' && c <= '9'
}

func TestRandomLower(t *testing.T) {
	l := gen.RandomLower(5)
	for i := 0; i < len(l); i++ {
		if !isLowerChar(l[i]) {
			t.Errorf("The %c is not lower in %s", l[i], l)
		}
	}
}

func TestRandomNumber(t *testing.T) {
	l := gen.RandomNumber(20)
	for i := 0; i < len(l); i++ {
		if !isDigital(l[i]) {
			t.Errorf("The %c is not digital in %s", l[i], l)
		}
	}
}
