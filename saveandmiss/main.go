package main

import "fmt"

func SaveAndMiss(arg string, num int) string {
	res := ""
<<<<<<< HEAD
	if num <= 0 {
		return arg
	}

	for i := 0; i < len(arg); i+=num*2 {
		end := i + num
		if end > len(arg) {
			end = len(arg)
		}
		res += string(arg[i:end])
	}
	return res
}
=======
	result := ""
	// if arg == ""{
	// 	return ""
	// }
	// slice := []rune(arg)
	// res := []rune{}
	// for i := 0; i < len(slice); i++{
		
	// }

	// return string(res)

	 
	if num <= 0 {
		return arg
	}
	// for _, ch := range arg {
	// 	if ch == ' ' {
	// 		continue
	// 	}
	// 	res += string(ch)
	// }
	// for  i := 0; i < len(res); i+=num {
	// 	if i + num > len(res) {
	// 		result += res[i:]
	// 	} else {
	// 		result += string(res[i:i+num])
	// 	}
	// 	i++
	// }
	// return result
	
	for i := 0; i < len(arg); i++ {
		split := i+num
		if split > len(arg) {
			split = len(arg)
		}
		res += arg[i:split]
		result += res
		res = ""
		i = i+(num*2)-1
	}

	return result
}


>>>>>>> d9e9c31757a4f1c215a1464dc081e74b1d5b88c7
func main() {
	fmt.Println(SaveAndMiss("123456789", 3))
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(SaveAndMiss("", 3))
	fmt.Println(SaveAndMiss("hello you all ! ", 0))
	fmt.Println(SaveAndMiss("what is your name?", 0))
	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5))
}
