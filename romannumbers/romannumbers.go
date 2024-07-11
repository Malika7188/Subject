package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: rn <number>")
		return
	}

	input := os.Args[1]

	value, err := strconv.Atoi(input)
	if err != nil || value <= 0 || value >= 4000 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	roman := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	numbers := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

	index := len(numbers) - 1
	var final string

	for value > 0 {
		for numbers[index] <= value {
			if value-numbers[index] == 0 {
				if roman[index] == "CM" || roman[index] == "CD" || roman[index] == "XC" || roman[index] == "XL" || roman[index] == "IX" || roman[index] == "IV" {
					fmt.Print("(", string(roman[index][1]), " - ", string(roman[index][0]), ")",)
					final += roman[index]
					value -= numbers[index]
					break
				}
				final += roman[index]
				fmt.Print(roman[index])
				value -= numbers[index]
				break
			} else {
				if roman[index] == "CM" || roman[index] == "CD" || roman[index] == "XC" || roman[index] == "XL" || roman[index] == "IX" || roman[index] == "IV" {
					fmt.Print("(", string(roman[index][1]), " - ", string(roman[index][0]), ")", "+")
					final += roman[index]
					value -= numbers[index]
				} else {
					final += roman[index]
					fmt.Print(roman[index], "+")
					value -= numbers[index]
				}
			}
		}
		index--
	}

	fmt.Println()
	fmt.Println(final)
}
