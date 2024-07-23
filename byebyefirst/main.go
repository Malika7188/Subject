package main

import "fmt"

func ByeByeFirst(strings []string) []string {
	result := []string{}

	if len(strings) == 0 {
		return result
	}
	for i := 0; i < len(strings); i++ {
		if strings[i] == strings[0] {
			result = strings[1:]
		}
	}
	return result
}
func main() {
	fmt.Println(ByeByeFirst([]string{}))
	fmt.Println(ByeByeFirst([]string{"one arg"}))
	fmt.Println(ByeByeFirst([]string{"first", "second"}))
	fmt.Println(ByeByeFirst([]string{"", "abcd", "efg"}))
}
