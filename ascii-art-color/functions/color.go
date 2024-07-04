package functions

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func ColoredSub(substr string, fontarray [95][8]string, colorcode string) []string {
	concatenatedLines := make([]string, 8)

	for _, char := range substr {
		Ascii, err := PrintChar(char, fontarray)
		if err != nil {
			log.Fatal(err)
		}
		for i, line := range Ascii {
			if colorcode != "" {
				concatenatedLines[i] += fmt.Sprintf("\033[%sm%s\033[0m", colorcode, line)
			} else {
				concatenatedLines[i] += line
			}
		}
	}
	return concatenatedLines
}

func RgbToAnsi(color string) string {
	color = strings.ToLower(color)
	rgb := strings.TrimPrefix(color, "rgb(")
	rgb = strings.TrimSuffix(rgb, ")")
	slice := strings.Split(rgb, ",")

	r, _ := strconv.Atoi(strings.TrimSpace(slice[0]))

	g, _ := strconv.Atoi(strings.TrimSpace(slice[1]))

	b, _ := strconv.Atoi(strings.TrimSpace(slice[2]))

	ansi := fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)

	return ansi
}
