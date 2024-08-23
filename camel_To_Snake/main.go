package main

import "fmt"

func CamelToSnakeCase(s string) string {
	if s == "" {
		return s
	}
	if !Valid(s) {
		return s
	}
	res := ""
	for i := 0; i < len(s); i++ {
		if i != 0 && Up(rune(s[i])) && i+1 < len(s) && Low(rune(s[i+1])) {
			res += "_" + string(s[i])
			// res += string(s[i])
		} else if Low(rune(s[i])) || i == 0 && Up(rune(s[i])) {
			res += string(s[i])
		} else {
			return s
		}
	}
	return res

}
func Valid(str string) bool {
	for _, c := range str {
		if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
			return false
		}
	}
	return true
}
func Up(s rune) bool {
	return s >= 'A' && s <= 'Z'
}
func Low(s rune) bool {
	return s >= 'a' && s <= 'z'
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
