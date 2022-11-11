package gostream

import "fmt"

type A interface {
	AA()
}

type S1 struct {
	x int
}

func (s1 S1) AA() {
	s1.x = 4
}

type S2 struct {
	x int
}

func (s2 *S2) AA() {
	s2.x = 3
}

func abc() {
	s1 := S1{0}
	s1.AA()
	fmt.Println(s1.x)

	s2 := S2{0}
	s2.AA()
	fmt.Println(s2.x)
}
