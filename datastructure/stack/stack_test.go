//Package stack implements with a fixed sized array
package stack

import "testing"

func TestStack_Push(t *testing.T) {
	s := New(10)
	s.Push(1)
	s.Push(3)
	s.Push(20)
	a, _ := s.Pop()
	print(a)
}
