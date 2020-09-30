package hide

import tools "../dataio"

var _num []string

func init() {
	_num = tools.DataLoad("numbers.txt")
}

// FixDigits - hide digits from text
func FixDigits(l string) string {
	l = fixDigits(l)
	l = fixCommon(l, _num)
	return l
}

func fixDigits(l string) string {
	r := []rune(l)

	for i, v := range r {
		if (v >= 48) && (v <= 57) {
			r[i] = 42
		}
	}
	return string(r)
}
