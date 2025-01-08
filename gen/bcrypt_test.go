package gen_test

import (
	"testing"

	"github.com/ak5w/gutil/gen"
)

func TestBcrypt(t *testing.T) {
	passwd := "adddfffrrerer*^ffK)#M>"
	h := gen.BcryptHash(passwd)

	if !gen.BcryptCheck(passwd, h) {
		t.Errorf("passwd %s check fails", passwd)
	}
}
