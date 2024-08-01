package main

import "fmt"

func CamelToSnakeCase(s string) string {
	res := []rune{}

	if consUpp(s) {
		return s
	} else if validS(s) {
		return s
	}

	// res := []rune{}
	for  i, char := range s {

		if IsUpp(char) && i != 0{
			res = append(res, '_')
			res = append(res, char)
		}else{
			res = append(res, char)
		}
	}
	// fmt.Println("s")
	return string(res)

}
func consUpp(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if IsUpp(rune(s[i])) && IsUpp(rune(s[i+1])) {
			return true
		}
	}
	return false
}
func IsUpp(s rune) bool {
	return s >= 'A' && s <= 'Z'
	
}
func IsLower(s rune) bool {
	return s >= 'a' && s <= 'z'
	
}
func validS(s string) bool {
	for _, c := range s {
		if !(c >= 'a' && c <= 'z' ||  c >= 'A' && c <= 'Z'){
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
