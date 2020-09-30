package tests

import (
	"testing"

	hide "../hide"
)

func TestOrder1(t *testing.T) {

	r := hide.RemoveByLevel("phone", 50)
	if r == "*" {
		return
	}
	t.Fail()
}
