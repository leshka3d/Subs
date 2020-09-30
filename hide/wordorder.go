package hide

import (
	"strconv"
	"strings"

	tools "../dataio"
)

var _endings []string
var _suffix []string
var _dictOrder map[string]tools.DictWord

func init() {
	_dictOrder = tools.LoadDictionary()
	_endings = tools.DataLoad("ending.txt")
	_suffix = tools.DataLoad("suffix.txt")
}

// RemoveByLevel 0-100
func RemoveByLevel(mes string, level int) string {

	if level == 0 {
		return mes
	}

	//расчитать индекс
	l := len(_dictOrder) / 100
	if level < 100 {
		l = l * level
	} else {
		l = len(_dictOrder)
	}

	dictOrder := _dictOrder
	txt := mes
	xtxt := txt

	for _, c := range Space {
		if c == " " {
			continue
		}
		xtxt = strings.ReplaceAll(xtxt, c, " ")
	}
	words := strings.Split(xtxt, " ")
	for _, v := range words {
		if len(v) > 0 {
			vv := strings.ToLower(v)
			w, ok := dictOrder[vv]
			if ok {
				if w.Index > l {
					//txt = strings.ReplaceAll(txt, v, v+"/"+strconv.Itoa(w.Index))
					// try to find synonym
				} else {
					txt = strings.ReplaceAll(txt, v, "*")
				}
				continue
			}
			c := false

			for _, suf := range _endings {
				txt, c = replaceModify(vv, v, suf, txt, l)
				if c {
					break
				}
			}
			for _, suf := range _suffix {
				txt, c = replaceModify(vv, v, suf, txt, l)
				if c {
					break
				}
			}
			if vv[len(vv)-1] == 's' {
				vv = strings.TrimSuffix(vv, "s")
				for _, suf := range _endings {
					if (suf != "s") && (suf != "es") {
						txt, c = replaceModify(vv, v, suf, txt, l)
						if c {
							break
						}
					}
				}
				for _, suf := range _suffix {
					txt, c = replaceModify(vv, v, suf, txt, l)
					if c {
						break
					}
				}
			}

		}

	}
	return txt
}

//replaceFromDict(txt, word, wSuf string, l int) (string, bool) {
func replaceModify(wordlow, word, suf, txt string, l int) (string, bool) {
	wLow := wordlow //strings.ToLower(word)
	wSuf := strings.TrimSuffix(wLow, suf)
	if wLow == wSuf {
		return txt, false
	}
	r := false
	txt, r = replaceFromDict(txt, word, wSuf, l)
	if r {
		return txt, r
	}
	/*
		w, ok := _dictOrder[wSuf]
		if ok {
			if w.Index > l {
				txt = strings.ReplaceAll(txt, word, word+"/^"+strconv.Itoa(w.Index))
			} else {
				txt = strings.ReplaceAll(txt, word, "#")
			}
			return txt, true
		}*/

	txt, r = replaceFromDict(txt, word, wSuf+"e", l)
	if r {
		return txt, r
	}
	/*w, ok = _dictOrder[wSuf+"e"]
	if ok {
		if w.Index > l {
			txt = strings.ReplaceAll(txt, word, word+"/^"+strconv.Itoa(w.Index))
		} else {
			txt = strings.ReplaceAll(txt, word, "#")
		}
		return txt, true
	}
	*/
	wSuf1 := strings.TrimSuffix(wSuf, "i")
	if wSuf1 != wSuf {
		txt, r = replaceFromDict(txt, word, wSuf1+"y", l)
		if r {
			return txt, r
		}
		/*w, ok = _dictOrder[wSuf]
		if ok {
			if w.Index > l {
				txt = strings.ReplaceAll(txt, word, word+"/^"+strconv.Itoa(w.Index))
			} else {
				txt = strings.ReplaceAll(txt, word, "#")
			}
			return txt, true
		}*/
	}
	sufLen := len(wSuf)
	if sufLen > 3 {
		if wSuf[sufLen-1] == wSuf[sufLen-2] {
			ru := []rune(wSuf)
			wSuf1 = string(ru[0 : sufLen-1])
			txt, r = replaceFromDict(txt, word, wSuf1, l)
			if r {
				return txt, r
			}
			/*
				w, ok = _dictOrder[wSuf]
				if ok {
					if w.Index > l {
						txt = strings.ReplaceAll(txt, word, word+"/^"+strconv.Itoa(w.Index))
					} else {
						txt = strings.ReplaceAll(txt, word, "#")
					}
					return txt, true
				}*/
		}

	}

	wSuf1 = wSuf + "er"
	txt, r = replaceFromDict(txt, word, wSuf1, l)
	if r {
		return txt, r
	}
	/*
		w, ok = _dictOrder[wSuf]
		if ok {
			if w.Index > l {
				txt = strings.ReplaceAll(txt, word, word+"/^"+strconv.Itoa(w.Index))
			} else {
				txt = strings.ReplaceAll(txt, word, "#")
			}
			return txt, true
		}*/

	return txt, false
}
func replaceFromDict(txt, word, wSuf string, l int) (string, bool) {
	w, ok := _dictOrder[wSuf]
	if ok {
		if w.Index > l {
			txt = strings.ReplaceAll(txt, word, word+"/^"+strconv.Itoa(w.Index))
		} else {
			txt = strings.ReplaceAll(txt, word, "#")
		}
		return txt, true
	}
	return txt, false
}

//-s
// -es
//-ed
//-ing
// -ly и другие
// таблица неправильных глаголов
