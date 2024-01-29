package main

import (
	"fmt"
	"strconv"

	"github.com/MFQWKMR4/MyFPGo/pkg/hkt"
	"github.com/MFQWKMR4/MyFPGo/pkg/list"
	"github.com/MFQWKMR4/MyFPGo/pkg/maybe"
)

// example
func main() {

	a := maybe.Some(100)
	b := maybe.Some("10000000000")
	c := list.New([]int{1, 2, 3})

	var intToStr = func(i int) string {
		return strconv.Itoa(i)
	}

	var count = func(i string) int {
		return len(i)
	}

	var double = func(i int) int {
		return i * 2
	}

	var minus50 = func(i int) int {
		tmp := i - 50
		if tmp < 0 {
			return 0
		}
		return tmp
	}

	var divideBy = func(i int) hkt.K1[maybe.MaybeId, int] {
		if i == 0 {
			return maybe.None[int]()
		}
		return maybe.Some(100 / i)
	}

	var genN = func(n int) hkt.K1[list.ListId, int] {
		results := make([]int, n)
		for i := 0; i < n; i++ {
			results[i] = i + 1
		}
		return list.New(results)
	}

	res1 := hkt.Chain3(
		maybe.F(intToStr), // 100 -> "100"
		maybe.F(count),    // "100" -> 3
		maybe.F(double),   // 3 -> 6
	)(a)

	fmt.Println(res1)

	res2 := hkt.Chain3(
		maybe.F(intToStr), // 100 -> "100"
		maybe.F(count),    // "100" -> 3
		maybe.M(divideBy), // 3 -> 33
	)(a)

	fmt.Println(res2)

	res3 := hkt.Chain3(
		maybe.F(count),    // "100" -> 3
		maybe.F(minus50),  // 3 -> 0
		maybe.M(divideBy), // 0 -> None[int]
	)(b)

	fmt.Println(res3)

	res4 := hkt.Chain2(
		list.F(intToStr), // 1, 2, 3 -> "1", "2", "3"
		list.F(count),    // "1", "2", "3" -> 1, 1, 1
	)(c)

	fmt.Println(res4)

	res5 := hkt.Chain2(
		list.M(genN),   // 1 ,2, 3 -> 1, 1, 2, 1, 2, 3
		list.F(double), // 1, 1, 2, 1, 2, 3 -> 2, 2, 4, 2, 4, 6
	)(c)

	fmt.Println(res5)
}
