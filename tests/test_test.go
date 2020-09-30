package tests

import (
	"os"
	"testing"

	hide "../hide"
)

func TestXxx(t *testing.T) {
	r := hide.FixPr("Xxx")
	r = hide.FixPrep(r)
	r = hide.FixAx(r)
	r = hide.FixWord(r)
	if r != "Xxx" {
		t.Fail()
	}
}
func TestDigits1(t *testing.T) {

	r := hide.FixDigits("123 AA two")
	if r != "*** AA ***" {
		t.Fail()
	}
}
func TestMain(m *testing.M) {
	//setup()
	code := m.Run()
	//shutdown()
	os.Exit(code)

}
