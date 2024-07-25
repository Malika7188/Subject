package main

import "fmt"

func ZipString(s string) string {
	seen := make(map[rune]int)
	result := []rune{}

	for _, char := range s {
		if seen[char] == 0 {
			result = append(result, char)
		}
		seen[char]++
	}
	res := ""
	for _, char := range result {
		count := seen[char]
		res += Itoa(count) + string(char)
	}
	return res
}

func Itoa(number int) string {
	result := ""
	if number == 0 {
		return "0"
	}
	for number > 0 {
		digit := number % 10
		result = string(digit+'0') + result
		number /= 10
	}
	return result
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}
