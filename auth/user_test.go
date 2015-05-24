package auth

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	inputs := []struct{ in []string }{
		{[]string{"jad-b", "jad-b", "jad-b@jad-b.com"}},
	}
	for _, inputs := range inputs {
		val := inputs.in
		user := NewUser(val[0], val[1], val[2])
		exp := &User{val[0], val[1], val[2]}
		if !user.equal(exp) {
			t.Errorf("%+v != %+v", user, exp)
		}
	}
}
