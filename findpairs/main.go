package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/01-edu/z01"
)

func Pairs(n []int, target int) [][]int {
	res := [][]int{}

	for i := 0; i < len(n); i++ {
		for j := i + 1; j < len(n); j++ {
			if n[i]+n[j] == target {
				res = append(res, []int{i, j})
			}
		}
	}
	return res

}
func parseArray(input string) []int {
	input = strings.Trim(input, "[]")
	parts := strings.Split(input, ",")
	arr := []int{}
	for _, slice := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(slice))
		arr = append(arr, num)
	}
	return arr
}

func print(val [][]int) {
	length := len(val)

	z01.PrintRune('[')

	// fmt.Println(length)
	for i := 0; i < length; i++ {
		z01.PrintRune('[')
		for j := 0; j < len(val[i]); j++ {
			if j != 0 {
				z01.PrintRune(' ')
			}
			z01.PrintRune(rune(val[i][j]) + '0')

		}
		z01.PrintRune(']')
		if i < length-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune(']')
	z01.PrintRune('\n')
}

func main() {

	if len(os.Args) != 3 {
		return
	}
	args := parseArray(os.Args[1])
	num, _ := strconv.Atoi(os.Args[2])

	print(Pairs(args, num))
	// nums, target := []int{1, 2, 3, 4, 5}, 6
	// nu, tar := []int{-1, 2, -3, 4, -5},1
	// print(Pairs(nums, target))
	// print(Pairs(nu, tar))
}
