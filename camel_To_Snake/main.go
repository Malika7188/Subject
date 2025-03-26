package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func CamelToSnakeCase(s string) string {
	if s == "" {
		return s
	}
	res := ""

	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && (s[i] >= 'A' && s[i] <= 'Z' && s[i+1] >= 'A' && s[i+1] <= 'Z') || (s[i] >= '0' && s[i] <= '9') || (i == len(s)-1 && s[i] >= 'A' && s[i] <= 'Z') {
			return s
		}
		if i > 0 && (s[i] >= 'A' && s[i] <= 'Z') {
			res += "_"
		}
		res += string(s[i])
	}
	return res
}

func main() {
	args := []string{
		"CamelCase",
		"camelCase",
		"HelloWorld",
		"132",
		" ",
		"",
		"A",
		"abcs",
		"AbC",
		"AbCEf",
		"abcAree",
		"ahe1Abde",
		"tesTing1",
		"SOME_VARIABLE",
		"ASuperLonGVariableName",
		"thisIsaTestOfCamelCase",
		"aA",
	}
	for _, arg := range args {
		challenge.Function("CamelToSnakeCase", CamelToSnakeCase, solutions.CamelToSnakeCase, arg)
	}
}
