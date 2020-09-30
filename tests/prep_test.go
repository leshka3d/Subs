package tests

import (
	"testing"

	hide "../hide"
)

func TestPrepSimple1(t *testing.T) {
	r := hide.FixPrep("in")
	if r == "**" {
		return
	}
	t.Fail()
}

func TestPrepSimple2(t *testing.T) {
	r := hide.FixPrep("Get me in!")
	if r == "Get me **!" {
		return
	}
	t.Fail()
}
