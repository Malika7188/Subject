package main

//import "golang.org/x/text/number"
import "fmt"

func Isprime(number int) bool {
	if number < 1 {
		return false
	}
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func FindPrevPrime(n int) int {
	if !Isprime(n) {
		n--
	}
	if n < 2 {
		return -1
	}
	return n
}

func FindNextPrime(n int) int {
	if !Isprime(n) {
		n++
	}
	return n
}

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}
