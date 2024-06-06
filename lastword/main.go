package main

import (
	"fmt"
	"os"
)

func main() {
	// args := os.Args[1]

	// if len(os.Args) == 1 || len(os.Args) > 2 {
	// 	return
	// }
	// word := ""


	// for i := len(args)-1; i > 0; i--{
	// 	if word == " " && string(args[i]) == " " {
	// 		continue
	// 	}
	// 	if word != " " && string(args[i]) == " " {
	// 		break
	// 	}
	// 	word = string(args[i]) + word
	// }
	// fmt.Println(word)

	// for args[len(args)-1] == ' ' {
	// 	args = args[:len(args)-1]
	// }

	// for i := len(args) - 1; i >= 0; i-- {
	// 	if args[i] == ' ' {
	// 		word = args[i+1:]
	// 		break
	// 	}
	// }

	// fmt.Println(word)

	args := os.Args[1]

	if len(os.Args) == 1 {
		return
	}
	word := ""

	for i := len(args)-1; i > 0; i-- {
		if word == " " && string(args[i]) == " " {
			continue
		}
		if word != " " && string(args[i]) == " " {
			break
		}
		word = string(args[i]) + word
	}
	fmt.Println(word)
}
