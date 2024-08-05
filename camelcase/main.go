package main

import "fmt"

func CamelToSnakeCase(s string) string {
	// res := []string{}
	result := ""

	for i, ch := range s {
		if i > 0 && (ch >= 'A' && ch <= 'Z') {
			result += s[:i] + "_" + s[i:]
			break
		} else if i != 0 && (ch >= '0' && ch <= '9') {
			return s
		} else if i != 0 && (s[i+1] == 'A' || s[i+1] == 'Z') {
			return s
		}
	}
	return result
}
func IsUp(s string) bool {
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
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
