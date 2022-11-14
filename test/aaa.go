package test

import "fmt"

type aaa struct {
	x int
}

func ABC() {
	a := aaa{3}
	fmt.Println(a.x)

}
