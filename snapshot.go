package main

import (
	"bufio"
	"fmt"
	"github.com/CodisLabs/codis/pkg/utils/errors"
	"os"
	"strings"
)

// 1. stack by go
// 2. text list & undo

type Stack struct {
	elements []string
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Pop() (v string, err error) {
	if len(s.elements) > 0 {
		v = s.elements[len(s.elements)-1]
		s.elements = s.elements[:len(s.elements)-1]
		return
	}
	err = errors.New("stack is null")
	return
}

func (s *Stack) Push(v string) {
	s.elements = append(s.elements, v)
}

type InputText struct {
	text string
	snapShot *SnapShot
}

func NewInputText(snapShot *SnapShot) *InputText {
	return &InputText{snapShot: snapShot}
}

func (i *InputText) List() string {
	return i.text
}

func (i *InputText) Append(v string) {
	i.snapShot.record(i.text)
	i.text += v + "\n"
}

func (i *InputText) Undo() {
	if text, err := i.snapShot.Undo(); err == nil {
		i.text = text
	}
}

type SnapShot struct {
	s *Stack	// record all data
}

func NewSnapShot(s *Stack) *SnapShot {
	return &SnapShot{s: s}
}

func (snap *SnapShot) record(v string) {
	snap.s.Push(v)
}

func (snap *SnapShot) Undo() (v string, err error) {
	return snap.s.Pop()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	snapShot := NewSnapShot(NewStack())
	i := NewInputText(snapShot)
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if text == ":list" {
			fmt.Println(i.List())
		} else if text == ":undo" {
			i.Undo()
			fmt.Println(i.List())
		} else {
			i.Append(text)
		}
	}
}