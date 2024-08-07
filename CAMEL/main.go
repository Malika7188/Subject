package main

import "fmt"

func CamelToSnakeCase(s string) string {
	// res := []string{}
	result := ""

	for i, ch := range s {
		if IsCup(ch) && IsCup(rune(s[i+1])) {
			return s
		}
		if IsCup(ch) && i != 0 {
			result += "_" + string(ch)
		} else {
			result += string(ch)
		}
	}
	return result
}
func IsCup(s rune) bool {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, c := range str {
		if s == c {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
