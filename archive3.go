package main

import (
	"fmt"
)

type structA struct {
	val int
}

type structB struct {
	val int
}

func (a structA) setVal(i int) {
	a.val = i
}

func (b *structB) setVal(i int) {
	b.val = i
}

func printStructsABCaseWithPointers() {
	s1 := structA{}
	s2 := &structB{}

	s1.setVal(2) // s1 is an existing value
	s2.setVal(3) // s2 has a pointer receiver

	fmt.Println(s1.val) // 0
	fmt.Println(s2.val) // 3
}
