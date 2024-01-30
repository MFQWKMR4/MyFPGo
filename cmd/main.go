package main

import (
	"errors"
	"fmt"
	"strconv"

	. "github.com/MFQWKMR4/MyFPGo/pkg/hkt"
	"github.com/MFQWKMR4/MyFPGo/pkg/list"
	"github.com/MFQWKMR4/MyFPGo/pkg/maybe"
	"github.com/MFQWKMR4/MyFPGo/pkg/result"
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

	var divideBy = func(i int) K1[maybe.Maybe, int] {
		if i == 0 {
			return maybe.None[int]()
		}
		return maybe.Some(100 / i)
	}

	var genN = func(n int) K1[list.List, int] {
		results := make([]int, n)
		for i := 0; i < n; i++ {
			results[i] = i + 1
		}
		return list.New(results)
	}

	var add3 = func(i int) K2[result.Result, int, error] {
		if i < 7 {
			return result.Ok[int, error](i + 3)
		}
		return result.Err[int, error](errors.New("too big"))
	}

	res1 := Chain3(
		maybe.F(intToStr), // 100 -> "100"
		maybe.F(count),    // "100" -> 3
		maybe.F(double),   // 3 -> 6
	)(a)

	fmt.Println(res1)

	res2 := Chain3(
		maybe.F(intToStr), // 100 -> "100"
		maybe.F(count),    // "100" -> 3
		maybe.M(divideBy), // 3 -> 33
	)(a)

	fmt.Println(res2)

	res3 := Chain3(
		maybe.F(count),    // "100" -> 3
		maybe.F(minus50),  // 3 -> 0
		maybe.M(divideBy), // 0 -> None[int]
	)(b)

	fmt.Println(res3)

	res4 := Chain2(
		list.F(intToStr), // 1, 2, 3 -> "1", "2", "3"
		list.F(count),    // "1", "2", "3" -> 1, 1, 1
	)(c)

	fmt.Println(res4)

	res5 := Chain3(
		list.M(genN),   // 1 ,2, 3 -> 1, 1, 2, 1, 2, 3
		list.F(double), // 1, 1, 2, 1, 2, 3 -> 2, 2, 4, 2, 4, 6
		list.F(double), // 2, 2, 4, 2, 4, 6 -> 4, 4, 8, 4, 8, 12
	)(c)

	fmt.Println(res5)

	res6 := Chain22(
		result.Mw(add3), // Ok(5) -> Ok(8)
		result.Mw(add3), // Ok(8) -> Err("too big")
	)(result.Ok[int, error](5))

	fmt.Println(res6)

	//
	//

	sample1()
}

func sample1() {

	a := 1
	b := 2

	c := list.New([]*int{&a, nil, &b})
	d := list.Map[*int, K1[maybe.Maybe, int]](maybe.From[int])(c)

	var double = func(i int) int {
		return i * 2
	}

	var isDefined = func(i K1[maybe.Maybe, int]) bool {
		return i != maybe.None[int]()
	}

	data := list.Chain(d)(
		list.Filter(isDefined),
		list.Map(maybe.Map(double)),
	)

	list.ForEach(func(a K1[maybe.Maybe, int]) {
		fmt.Println(a)
	})(data)
}
