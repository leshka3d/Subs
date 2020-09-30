package actions

import (
	"strconv"
	"strings"

	hide "../hide"
)

func empty(x string) string {
	return x
}

// GetActions - create configuration parameters object
func GetActions(p1, p2 string) []func(string) string {
	var l = len(p1)
	if l > 7 {
		l = 7
	}
	l++
	var aindex = 0
	var act = make([]func(string) string, l)

	wordTreshold, _ := strconv.Atoi(p2)

	if strings.Contains(p1, "A") {
		act[aindex] = hide.FixAx
		aindex++
	}
	if strings.Contains(p1, "P") {
		act[aindex] = hide.FixPrep
		aindex++
	}
	if strings.Contains(p1, "N") {
		act[aindex] = hide.FixPr
		aindex++
	}
	if strings.Contains(p1, "W") {
		act[aindex] = hide.FixWord
		aindex++
	}
	if strings.Contains(p1, "V") {
		// verb forms not ready yet
		act[aindex] = empty
		aindex++
	}
	if strings.Contains(p1, "D") {
		act[aindex] = hide.FixDigits
		aindex++
	}
	// => remove numbers
	// clean line if empty
	act[aindex] = func(s string) string {
		return hide.RemoveByLevel(s, wordTreshold)
	}
	return act
}
