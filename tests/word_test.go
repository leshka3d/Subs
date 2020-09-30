package tests

import (
	"testing"

	hide "../hide"
)

func TestWordSimple1(t *testing.T) {
	r := hide.FixWord("wEll")
	if r == "****" {
		return
	}
	t.Fail()
}
