package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	minRand = 1
	maxRand = 100
)

type Stack struct {
	impl []int
	len int
}

func (s *Stack) Push(val int) {
	s.impl[s.len] = val
	s.len++
}

func (s *Stack) Pop() (int, bool) {
	if s.len == 0 {
		return 0, false
	}

	s.len--

	return s.impl[s.len], true
}

func (s *Stack) Top() (int, bool) {
	if s.len == 0 {
		return 0, false
	}

	return s.impl[s.len-1], true
}

func (s *Stack) Empty() bool {
	return s.len == 0
}

func main()  {
	rand.Seed(time.Now().Unix())
	s := make([]int, 3, 3)
	for i := 0; i < len(s); i++ {
		s[i] = randValue()
	}
	st := &Stack{impl: s, len: len(s)}

	fmt.Printf("%v\n", st)
	pop, ok := st.Pop()
	fmt.Printf("%d, %t\n", pop, ok)
	top, ok := st.Top()
	fmt.Printf("%d, %t\n", top, ok)

	sorted := insertionSort(st)
	fmt.Printf("%v", sorted)
	selectionSort(sorted)
	fmt.Printf("%v\n", sorted)
}

func insertionSort(stack *Stack) *Stack {
	toSort := &Stack{impl: make([]int, stack.len, stack.len)}

	for !stack.Empty() {
		temp, _ := stack.Pop()

		top, _ := toSort.Top()
		for !toSort.Empty() && top > temp {
			value, _ := toSort.Pop()
			stack.Push(value)
		}
		toSort.Push(temp)
	}

	return toSort
}

func selectionSort(stack *Stack) {
	stackLen := stack.len
	tmpStack := &Stack{impl: make([]int, stack.len, stack.len)}

	for i := 0; i < stackLen; i++  {
		var largest int
		for j := 0; j < stackLen - i; j++ {
			curr, _ := stack.Pop()
			if curr > largest {
				largest = curr
			}
			tmpStack.Push(curr)
		}
		stack.Push(largest)
		largestSkipped := false
		for !tmpStack.Empty() {
			val, _ := tmpStack.Pop()
			if val == largest && !largestSkipped {
				largestSkipped = true
				continue
			}
			stack.Push(val)
		}
	}
}

func randValue() int {
	return rand.Intn(maxRand - minRand) + minRand
}