package messages

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func getMess(data string) (Message, error) {
	if len(data) == 0 {
		return Message{"", "", ""}, errors.New("incorrect data")
	}
	parts := strings.Split(data, "\r\n")

	l := len(parts)

	if l < 3 {
		return Message{"", "", ""}, errors.New("line count is incorrect")
	}
	m := ""
	for i := 2; i < l; i++ {
		m = m + " " + parts[i]
	}

	a := Message{parts[0], parts[1], m}
	return a, nil
}

// GetCurrentMessage !
func (mm *MessageMan) GetCurrentMessage() string {
	return mm.mes[mm.current].Message
}

// GetCount - get messages count
func (mm *MessageMan) GetCount() int {
	return len(mm.mes)
}

//NextMessage - next current message
func (mm *MessageMan) NextMessage() bool {
	mm.current++
	if mm.end < mm.current {
		return false
	}
	return true
}

// UpdateCurrentMessage !
func (mm *MessageMan) UpdateCurrentMessage(s string) {
	mm.mes[mm.current].Message = s
}

//Save to file
func (mm *MessageMan) Save(fileName string) {
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	m := mm.mes
	for _, v := range m {
		fmt.Fprintln(f, v.ID)
		fmt.Fprintln(f, v.Time)
		fmt.Fprintln(f, v.Message)
		fmt.Fprintln(f, "")
	}

}

// NewMessageMan same data new current
func NewMessageMan(mm MessageMan, start, end int) MessageMan {
	mm.current = start
	mm.end = end
	return mm
}

//GetMessageMan from file
func GetMessageMan(fileName string) (MessageMan, error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return MessageMan{mes: nil}, err
	}
	txt := string(content)
	mes := strings.Split(txt, "\r\n\r\n")
	l := len(mes)
	messList := make([]Message, l)
	var index int = 0
	for i := 0; i < l; i++ {
		messList[index], err = getMess(mes[i])
		if err == nil {
			index++
		}
	}
	var mm = MessageMan{mes: messList, current: 0, end: len(messList) - 1}

	return mm, nil
}
