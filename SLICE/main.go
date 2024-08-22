package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func Slice(a []string, nbrs ...int) []string {
	res := []string{}
	leng := len(a)

	if len(nbrs) == 1 {
		num := nbrs[0]
		if num < 0 {
			num += leng
		}
		res = append(res, a[num:]...)
	}
	if len(nbrs) >= 2 {
		num1 := nbrs[0]
		num2 := nbrs[1]

		if num1 < 0 || num2 < 0 {
			num1 += leng
			num2 += leng
		}
		if num1 > num2 {
			return []string(nil)

		}

		res = append(res, a[num1:num2]...)
	}
	return res
}

func main() {
	elems := [][]interface{}{
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			1,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			-3,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			2, 4,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			-2, -1,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			2, 0,
		},
	}

	s := random.StrSlice(chars.Words)

	elems = append(elems, []interface{}{s, -len(s) - 10, -len(s) - 5})

	for i := 0; i < 3; i++ {
		s = random.StrSlice(chars.Words)
		elems = append(elems, []interface{}{s, random.IntBetween(-len(s)-10, len(s)+10), random.IntBetween(-len(s)-8, len(s)+10)})
	}

	for _, a := range elems {
		challenge.Function("Slice", Slice, solutions.Slice, a...)
	}
}
