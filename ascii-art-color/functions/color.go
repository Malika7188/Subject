package functions

import "fmt"

func Colorstring(lines []string, color string) []string {
	for i := range lines {
		lines[i] = fmt.Sprintf("\033[%sm%s\033[0m", color, lines[i])
	}
	return lines
}
