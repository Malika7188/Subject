package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	if len(os.Args) < 1 {
		return
	}
	arg := os.Args[1:]
	for _, str := range arg {
		ans := Lasty(str)
		Sp(ans)
	}
}

func Lasty(str string) string {
	res := ""
	for i, ch := range str {
		if i != len(str)-1 {
			if Iscap(ch) && str[i+1] != ' ' {
				ch = ch + 32
			}
			if Islow(ch) && str[i+1] == ' ' {
				ch = ch - 32
			}
		} else if i == len(str)-1 && Islow(ch) {
			ch -= 32
		}
		res += string(ch)
	}
	return res

}
func Sp(str string) {
	for _, ch := range str {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func Iscap(s rune) bool {
	return s >= 'A' && s <= 'Z'

}

func Islow(s rune) bool {
	return s >= 'a' && s <= 'z'
}
