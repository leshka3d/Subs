package dataio

import (
	"io/ioutil"
	"log"
	"os"
	filepath "path/filepath"
	"runtime"
	"strings"
)

//DictWord  - frequency and traslation
type DictWord struct {
	Index   int
	Meaning string
}

func getDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	if strings.HasSuffix(os.Args[0], ".test.exe") {
		_, filename, _, _ := runtime.Caller(0)
		dir = strings.TrimSuffix(filename, "dataio/dataio.go")
	}
	if strings.HasSuffix(os.Args[0], ".test") {
		dir = strings.TrimSuffix(dir, "tests")
	}
	return strings.TrimSuffix(dir, "/")
}

// DataLoad load text file to line array
func DataLoad(f string) []string {
	dir := getDir()

	content, err := ioutil.ReadFile(dir + "/_data/" + f)
	if err == nil {
		txt := string(content)
		txt = strings.ReplaceAll(txt, "\r", "")
		arr := strings.Split(txt, "\n")
		for i, v := range arr {
			arr[i] = strings.Split(v, "-")[0]
		}
		return arr
	}
	return nil
}

//LoadDictionary - load dictionary data
func LoadDictionary() map[string]DictWord {
	// https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists
	// w_pg.txt   Project Gutenberg
	// w from TV show and movies
	//words.txt words & translation
	lines := DataLoad("w_pg.txt")
	index := 0
	m := make(map[string]DictWord, len(lines))
	for _, v := range lines {
		//parts := strings.Split(v, "-")
		i := index
		//m[strings.ToLower(parts[0])] = DictWord{i, parts[1]}
		m[v] = DictWord{i, ""}
		index++
	}
	return m
}
