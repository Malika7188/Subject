package main

import (
	"fmt"
	"os"
	"strconv"
)
func RM(number int)(string, string) {
	result := ""
	ans := ""

	values := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
	roman := []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}
	calculations := []string{"M","(M-C)","D","(D-C)","C","(C-X)","L","(L-X)","X","(X-I)","V","(V-I)","I"}

	for i, val := range values {
		for number >= val {
			number -= val
			result += roman[i]

			if len(ans) > 0 {
				ans += "+"
			}
			ans += calculations[i]
		}
	}
	return result, ans
}
func main() {
	if len(os.Args) != 2 {
		return
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 0 || num >= 4000 {
		fmt.Println("Error: cannot convert to digit")
		return
	}

	result, ans := RM(num)
	fmt.Println(ans)
	fmt.Println(result)
}
