// package main

// import (
// 	"os"
// 	"strconv"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	if len(os.Args) != 2 {
// 		return
// 	}
// 	num, err := strconv.Atoi(os.Args[1])
// 	if err != nil || num <= 0 || num > 4000 {
// 		return
// 	}

// 	result, ans := Roman(num)

// 	for _, val1 := range ans {
// 		z01.PrintRune(val1)
// 	}
// 	z01.PrintRune('\n')

// 	for _, val := range result {
// 		z01.PrintRune(val)
// 	}
// 	z01.PrintRune('\n')
// }

// func Roman(num int) (string, string) {
// 	result := ""
// 	ans := ""

// 	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
// 	roman := []string{"M", "CM", "D", "CD", "C", "LC", "L", "XL", "X", "IX", "V", "IV", "I"}
// 	calculations := []string{"M", "(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(V-I)", "I"}

// 	for i, val := range values {
// 		for num >= val {
// 			num -= val
// 			result += roman[i]
// 			if len(ans) > 0 {
// 				ans += "+"
// 			}
// 			ans += calculations[i]
// 		}
// 	}
// 	return result, ans
// }

package main

import (
	// "fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

var romanValues = []struct {
	key   int
	value string
	calc  string
}{
	{1000, "M", "M"},
	{900, "CM", "(M-C)"},
	{500, "D", "D"},
	{400, "CD", "(D-C)"},
	{100, "C", "C"},
	{90, "XC", "(C-X)"},
	{50, "L", "L"},
	{40, "XL", "(L-X)"},
	{10, "X", "X"},
	{9, "IX", "(X-I)"},
	{5, "V", "V"},
	{4, "IV", "(V-I)"},
	{1, "I", "I"},
}

func Rom(number int)(string, string) {
	result := ""
	calculation := ""
	for _, val := range romanValues {
		for number >= val.key {
			result += val.value
			calculation += val.calc
			number -= val.key
			if number > 0 {
				calculation += "+"
			}
		}
	}
	return result, calculation
}

func main() {
	number, err := strconv.Atoi(os.Args[1])
	if err != nil || number < 0 || number >= 4000 {
		return
	}
	result, calculation := Rom(number)

	for _, char := range calculation {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')

	for _, val := range result {
		z01.PrintRune(val)
	}
	z01.PrintRune('\n')
}
