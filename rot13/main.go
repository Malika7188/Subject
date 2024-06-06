package main

import "fmt"

//"net/netip"
//"os"

func Rot13(s string) string {
	// args := os.Args
	result := []rune{}
	var Newchar rune

	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			Newchar = (char-'a'+13)%26 + 'a'
			result = append(result, Newchar)
		} else if char >= 'A' && char <= 'Z' {
			Newchar = (char-'A'+13)%26 + 'A'
			result = append(result, Newchar)
		} else {
			result = append(result, char)
		}
	}
	fmt.Println(string(result))
	return string(result)
}

func main() {
	//Rot13("abc")
	Rot13("hello there")
}
