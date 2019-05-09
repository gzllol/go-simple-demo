//Package stack implements with a fixed sized array
package stack

import "errors"

//Stack struct
type Stack struct {
	cap   int
	size  int
	array []int
}

//New a stack
func New(cap int) *Stack {
	s := new(Stack)
	s.cap = cap
	s.size = 0
	s.array = make([]int, cap, cap)
	return s
}

//Push an entry into stack
func (s *Stack) Push(a int) (int, error) {
	if s.cap == s.size {
		return 0, errors.New("full")
	}
	s.array[s.size] = a
	s.size++
	return 1, nil
}

//Pop an entry from stack
func (s *Stack) Pop() (int, error) {
	if s.cap == 0 {
		return 0, errors.New("empty")
	}
	r := s.array[s.size-1]
	s.size--
	return r, nil
}
