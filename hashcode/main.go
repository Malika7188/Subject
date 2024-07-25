package main

import "fmt"

func HashCode(dec string) string {
	hashed := make([]rune, len(dec))
	size := len(dec)
	
	for i, char := range dec {
		hashvals := (int(char) + size)%127
		
		if hashvals < 32 || hashvals > 126 {
			hashvals += 33
		}
		hashed[i] = rune(hashvals)	
	}
	return string(hashed)
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
