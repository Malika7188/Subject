package main

import "fmt"

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
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}
