package main

import (
	"log"
	"reflect"
)

var testCases = []struct {
	in   string
	want string
}{
	{
		in:   " ",
		want: "\n",
	},
	{
		in:   "FOR PONY",
		want: "PONY\n",
	},
	{
		in:   "this ... is sparta, then again, maybe not",
		want: "not\n",
	},
	{
		in:   " lorem,ipsum ",
		want: "lorem,ipsum\n",
	},
}

func main() {
	for _, tc := range testCases {
		got := LastWord(tc.in)
		if !reflect.DeepEqual(got, tc.want) {
			log.Fatalf("%s(%q) == %q instead of %q\n",
				"LastWord",
				tc.in,
				got,
				tc.want,
			)
		}
	}
}

func LastWord(s string) string {
	if s == "" {
		return "\n"
	}
	word := false
	res := ""
	for i := len(s) - 1; i > 0; i-- {
		if s[i] != ' ' {
			res = string(s[i]) + res
			word = true
		} else if word {
			break
		}
	}

	return res + "\n"
}

// func main() {
// 	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
// 	fmt.Print(LastWord(" lorem,ipsum "))
// 	fmt.Print(LastWord(" "))
// }
