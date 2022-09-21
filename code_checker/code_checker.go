package main

import (
	"errors"
	"fmt"
	"strings"
)

type Stack struct {
	stack []string
}

func NewStack() Stack {
	return Stack{
		stack: []string{},
	}
}

func (s *Stack) Push(newData string) {
	s.stack = append(s.stack, newData)
}

func (s *Stack) Pop() (string, error) {
	max := len(s.stack) - 1
	if max >= 0 {
		result := s.stack[max]
		s.stack = s.stack[0:max]
		return result, nil
	}
	return "", errors.New("Empty stack")
}

func CheckData(intput string, popKey string, pushKey string) bool {
	data := strings.Split(intput, "")
	s := NewStack()
	for i := 0; i < len(data); i++ {
		if data[i] == popKey {
			s.Push(data[i])
		} else if data[i] == pushKey {
			_, err := s.Pop()
			if err != nil {
				return false
			}
		}
	}
	_, err := s.Pop()
	return err != nil
}

func main() {
	fmt.Println(CheckData("{{{}}}", "{", "}"))
	fmt.Println(CheckData("{{{{}}", "{", "}"))
	fmt.Println(CheckData("{}{}", "{", "}"))
	fmt.Println(CheckData("{}}{{}", "{", "}"))
}
