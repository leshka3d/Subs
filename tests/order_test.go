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
func TestOrder3(t *testing.T) {

	r := hide.RemoveByLevel("caller", 50)
	if r == "*" {
		return
	}
	t.Fail()
}
func TestOrder2(t *testing.T) {

	r := hide.RemoveByLevel("walkman", 50)
	if r == "walkman" {
		return
	}
	t.Fail()
}
func TestOrder4(t *testing.T) {

	r := hide.RemoveByLevel("waolkmaniotic", 100)
	if r == "waolkmaniotic" {
		return
	}
	t.Fail()
}
func TestOrder5(t *testing.T) {

	r := hide.RemoveByLevel("waolkmaniotic", 99)
	if r == "waolkmaniotic" {
		return
	}
	t.Fail()
}
