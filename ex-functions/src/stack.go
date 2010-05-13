package main

import (
	"fmt"
	"strconv"
)

type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(k int) (ok bool) {
	if s.i+1 > 9 {
		return false
	}
	s.data[s.i] = k
	s.i++
	return true
}

func (s *stack) pop() (ret int, ok bool) {
	if s.i-1 < 0 {
		return 0, false
	}
	ret = s.data[s.i]
	s.i--
	return ret, true
}

func (s *stack) String() string {
	var str string
	for i := 0; i < s.i; i++ {
		str = str + "[" +
			strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}

func main() {
	s := new(stack) // returns pointer
	s.push(25)
	s.push(14)
	fmt.Printf("stack %v\n", s)
}