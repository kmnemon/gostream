package gostream

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	a := []int{6, 5, 3, 4, 5}

	x := StreamOf(a).
		Map(func(x int) int {
			return x + 1
		}).
		ToList()

	expect := []int{7, 6, 4, 5, 6}

	if !equalSliceHelper(expect, x) {
		t.Error("Map operator has some problem")
	}

}

func TestDistinct(t *testing.T) {
	a := []int{6, 5, 3, 4, 5}

	x := StreamOf(a).
		Distinct().
		ToList()

	expect := []int{6, 5, 3, 4}

	if !equalSliceHelper(expect, x) {
		t.Error("Distinct operator has some problem")
	}
}

func TestReduceWithInitValue(t *testing.T) {
	a := []int{6, 5, 3, 4, 5}

	x := StreamOf(a).
		ReduceWithInitValue(1, func(x int, y int) int {
			return x + y
		}).
		ToList()

	fmt.Println(x)

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

	// ages := map[string]int{
	// 	"rr": 1000,
	// 	"vv": 2000,
	// 	"cc": 500,
	// }

	// StreamOfMap(ages).Map(func(e EntrySet[string, int]) EntrySet[string, int] {
	// 	e.V = e.V + 1
	// 	return e
	// }).ForEach(func(e EntrySet[string, int]) {
	// 	fmt.Println(e)
	// })

}

func equalSliceHelper[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
