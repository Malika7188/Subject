package main

import "fmt"

func ConcatParams(args []string) string {
	res := ""
	for i := 0; i < len(args); i++ {
		if i > 0 {
			res += "\n"
		}
		res += args[i]
		
	}
	
	return res
}
func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}
