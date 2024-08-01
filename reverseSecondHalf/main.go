package main

import "fmt"

func ReverseSecondHalf(str string) string {
	if len(str) == 0 {
		return "invalid input\n"
	}
	if len(str) == 1 {
		return str
	}
	sechalf := len(str) / 2
	res := str[sechalf:]

	result := ""
	for i := len(res) - 1; i >= 0; i-- {
		result += string(res[i])
	}
	return result + "\n"
}
func main() {
	fmt.Print(ReverseSecondHalf("This is the 1st half This is the 2nd half"))
	fmt.Print(ReverseSecondHalf(""))
	fmt.Print(ReverseSecondHalf("Hello World"))
	// fmt.Print(ReverseSecondHalf("me and her n"))
}
