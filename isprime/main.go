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
	for n > 1{
		if Isprime(n){
			return n
		}
		n--	
	}
	if n < 1 {
		return 0
	}
	return n
}

func FindNextPrime(n int) int {
	for n > 1 {
		if Isprime(n) {
			return n
		}
		n++
	}
	if n < 0 {
		return 0
	}
	return n
}

func main() {
	fmt.Println(FindNextPrime(-5))
	fmt.Println(FindNextPrime(4))
}
