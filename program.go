package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	act "./actions"
	messages "./messages"
)

func processLines(mm *messages.MessageMan, syn *sync.WaitGroup, print bool) {
	i := 0
	syn.Add(1)
	defer syn.Done()
	for true {
		fstr := mm.GetCurrentMessage()
		for _, v := range _actions {
			if v == nil {
				break
			}
			fstr = v(fstr)
		}

		mm.UpdateCurrentMessage(fstr)
		i++
		if print {
			if i%30 == 0 {
				fmt.Print(i)
				fmt.Print(" / ")
				fmt.Println(mm.GetCount())
			}
		}
		if !mm.NextMessage() {
			break
		}
	}
}

var _actions []func(string) string

func main() {
	var wg sync.WaitGroup

	cpu := runtime.NumCPU()
	//cpu = 1
	runtime.GOMAXPROCS(cpu)
	_actions = act.GetActions(os.Args[1], os.Args[2])

	fmt.Println(time.Now().Format(time.RFC850))
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])

	mm, e := messages.GetMessageMan(os.Args[3])
	if e != nil {
		log.Fatal(e)
		return
	}
	cnt := mm.GetCount()
	cnt = cnt / cpu

	for th := 1; th < cpu; th++ {
		mmCurrent := messages.NewMessageMan(mm, cnt*(th-1), cnt*th)
		go processLines(&mmCurrent, &wg, false)
	}
	mm2 := messages.NewMessageMan(mm, cnt*(cpu-1), mm.GetCount()-1)
	go processLines(&mm2, &wg, true)

	wg.Wait()
	mm.Save(os.Args[4])
	fmt.Println(time.Now().Format(time.RFC850))
	//fmt.Println(a)
}
