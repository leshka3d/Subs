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

func TestWordSimple2(t *testing.T) {
	r := hide.FixWord("I've gone")
	if r == "I've ****" {
		return
	}
	t.Fail()
}
func TestWordSimple3(t *testing.T) {
	r := hide.FixAll("I've gone")
	if r == "* *** ****" {
		return
	}
	t.Fail()
}
