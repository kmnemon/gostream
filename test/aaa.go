package test

import "fmt"

type aaa struct {
	x int
}

func ABC(func(x int) any) {
	a := aaa{3}
	fmt.Println(a.x)

}
