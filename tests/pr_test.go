package tests

import (
	"testing"

	hide "../hide"
)

func TestPrSimple1(t *testing.T) {

	r := hide.FixPr("you")
	if r == "***" {
		return
	}
	t.Fail()
}
func TestPrSimple2(t *testing.T) {
	r := hide.FixPr("public you")
	if r == "public ***" {
		return
	}
	t.Fail()
}

func TestPrSimple3(t *testing.T) {
	r := hide.FixPr("you and me can test")
	if r == "*** and ** can test" {
		return
	}
	t.Fail()
}
func TestPrSimple4(t *testing.T) {
	r := hide.FixPr("i've")
	if r == "*'ve" {
		return
	}
	t.Fail()
}
