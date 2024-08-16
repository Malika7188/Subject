package main

import "fmt"

func FromTo(from int, to int) string {
	res := ""

	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}
	if from == to {
		return Itoa(from) + "\n"
	}
	if from < to {
		for i := from; i <= to; i++ {
			if i < 10 {
				res += "0" + Itoa(i)
			} else {
				res += Itoa(i)
			}
			if i != to {
				res += ", "
			}
		}
	} else {
		for i := from; i >= to; i-- {
			if i < 10 {
				res += "0" + Itoa(i)
			} else {
				res += Itoa(i)
			}
			if i != to {
				res += ", "
			}
		}
	}
	return res + "\n"

}
func Itoa(num int) string {
	neg := false
	if num < 0 {
		neg = true
		num *= -1
	}
	if num == 0 {
		return "0"
	}
	res := ""
	for num > 0 {
		digit := num % 10
		res = string(digit+'0') + res
		num /= 10
	}
	if neg {
		return "-" + res
	}
	return res
}

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}
