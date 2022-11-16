package gostream

import (
	"fmt"
	"testing"
)

type AA struct {
	a int
	b int
}

func TestTest(t *testing.T) {
	a := []int{6, 5, 3, 4, 5}

	StreamOf(a).
		Filter(func(x int) bool {
			return x <= 5
		}).
		Limit(3).
		Map(func(x int) int {
			return x * 3
		}).
		Sorted().
		ForEach(func(x int) {
			fmt.Println(x)
		})

	// b := []AA{
	// 	{1, 2},
	// 	{3, 4},
	// }

	// less := func(a AA, b AA) bool {
	// 	return b.a < a.a
	// }

	// StreamOf(b).
	// 	// Map(func(x int) int {
	// 	// 	return x * 3
	// 	// }).
	// 	SortedWith(less).
	// 	ForEach(func(x AA) {
	// 		fmt.Println(x)
	// 	})
}
