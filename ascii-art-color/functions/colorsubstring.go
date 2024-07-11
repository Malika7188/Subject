package functions

import (
	"strings"
	"fmt"
)
func SubstringColoring(str []string, substr, color string) []string {

	colorCode := fmt.Sprintf("\033[%sm", color)
	resetCode := "\033[0m"
	
	//find the starting index of the substring
	startIndex := strings.Index(, substr)
	if startIndex == -1 { //checks if the substring is available and if not it returns the original string
		return str
	}
	//split the string into 3 parts: rhe innitial string, substring and the after string, which is te result after coloring
	initially := str[:startIndex]
	after := str[startIndex+len(substr):]

	return initially + colorCode + substr + resetCode + after

	s:=[]string{"hey","jkt"}
}