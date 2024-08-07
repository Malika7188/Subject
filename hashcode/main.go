package main

import "fmt"

func HashCode(dec string) string {
	// hashed := make([]rune, len(dec))
	var hashed []rune
	size := len(dec)
	hash := 0

	for _, char := range dec {
		hash = (int(char) + size) % 127
		if hash < 32 || hash > 126 {
			hash += 33
		}
		hashed = append(hashed, rune(hash))
	}
	return string(hashed)
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
	fmt.Println(HashCode(" 14 Avril 1912"))
}
