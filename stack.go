package main

import (
	"fmt"
)

// creating custom int type of stack
type stack []int

// defining is_empty method on the stack
func (s *stack) is_empty() bool {

	return len(*s) == 0
}

// defining push method on the stack
func (s *stack) push(elm int) {
	*s = append((*s), elm)
}

// pop method on the stack
func (s *stack) pop() {

	if s.is_empty() {

		fmt.Println(" stack is empty")
		return
	} else {

		ind := len(*s) - 1
		elm_pop := (*s)[ind]
		fmt.Printf("%v : is poped from the stack \n", elm_pop)
		*s = (*s)[:ind]

	}

}

// defining empty() method onthe stack
func (s *stack) empty() {

	if s.is_empty() {
		fmt.Println("The sack is empty already")
		return
	}

	*s = nil
}

func main() {

	var s stack

	// pushing elements to the stack
	fmt.Println("creating and intializing stack...")
	for i := 0; i < 10; i++ {
		s.push(i)
	}

	// printing stack

	fmt.Printf("stack-len:%v\n stack-cap:%v\nstack:%v\n ", len(s), cap(s), s)

	// pop the las element
	fmt.Println("Poping the last element of the stack...")
	s.pop()

	// making  the stack empty

	fmt.Println("removing all elements from the stack")
	s.empty()
	fmt.Printf("Stack :%v \n", s)

}
