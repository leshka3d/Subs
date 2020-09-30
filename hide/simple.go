package hide

import (
	"errors"
	"regexp"
	"strings"

	tools "../dataio"
)

var _star []string
var _ax []string
var _pr []string
var _p []string
var _words []string

// Space  space like chars
var Space []string

func init() {
	_star = []string{"", " * ", " ** ", " *** ", " **** ", " ***** ", " ****** "}
	_ax = tools.DataLoad("axlist.txt")
	_pr = tools.DataLoad("prlist.txt")
	_words = tools.DataLoad("list.txt")
	_p = tools.DataLoad("prep.txt")
	Space = []string{" ", ",", ".", "]", "[", "-", "+", "(", ")", "\"", "<", ">", "*", "?", "!", "\""}
}

//Star - get asterix string
func star(i int) string {
	if i > 6 {
		i = 6
	}
	if i < 0 {
		panic(errors.New("value out of range"))
	}
	return _star[i]
}
func cleanStarString(s string) string {

	for true {
		l := len(s)
		s = strings.ReplaceAll(s, "  *", " *")
		s = strings.ReplaceAll(s, "*  ", "* ")
		s = strings.ReplaceAll(s, " !", "!")
		s = strings.ReplaceAll(s, " ?", "?")
		if l == len(s) {
			break
		}
	}
	s = strings.TrimSpace(s)
	return s
}
func caseInsensitiveReplace(subject string, search string, replace string) string {
	searchRegex := regexp.MustCompile("(?i)" + search)
	return searchRegex.ReplaceAllString(subject, replace)
}

// FixAll - remove all nessesery data as configured
func FixAll(l string) string {
	l = fixCommon(l, _p)
	l = fixCommon(l, _words)
	l = fixCommon(l, _pr)
	l = fixCommon(l, _ax)
	return l
}

//FixPrep - remove prepositions
func FixPrep(l string) string {
	return fixCommon(l, _p)
}

//FixWord - remove most common words
func FixWord(l string) string {
	return fixCommon(l, _words)
}

// FixPr remove proniuns
func FixPr(l string) string {
	return fixCommon(l, _pr)
}

// FixAx  - remove axulary werbs
func FixAx(l string) string {
	return fixCommon(l, _ax)
}
func space2reg(s string) string {
	if s == " " {
		return s
	}
	return "\\" + s
}
func fixCommon(l string, wList []string) string {
	_ax := wList
	//fmt.Println("Recovered in f", len(wList))
	//_spaceReg := []string{" ", "\\,", "\\.", "\\]", "\\[", "\\-", "\\+", "\\(", "\\)", "\\\"", "\\<", "\\>", "\\*", "\\?"}
	space := Space
	l = strings.ReplaceAll(l, "\r", " ")
	l = strings.ReplaceAll(l, "\n", " ")
	l = strings.ReplaceAll(l, "'", " '")
	l = " " + l + " "

	for i := 0; i < len(_ax); i++ {
		starCnt := len(_ax[i])
		if _ax[i][0] == '\'' {
			for k := 0; k < len(space); k++ {
				l = caseInsensitiveReplace(l, _ax[i]+space2reg(space[k]), star(starCnt)+space[k])
			}
		} else {

			for j := 0; j < len(space); j++ {
				for k := 0; k < len(space); k++ {

					l = caseInsensitiveReplace(l, space2reg(space[j])+_ax[i]+space2reg(space[k]), space[j]+star(starCnt)+space[k])
				}
			}
		}
	}
	l = cleanStarString(l)
	l = strings.ReplaceAll(l, " '", "'")
	return l
}
