package functions

import (
	"strings"
)

func HandleCarriage(input string) string {
	// Split the input by lines
	lines := strings.Split(input, "\n")
	var processedLines []string

	for _, line := range lines {
		// If the line contains a carriage return, handle it
		if strings.Contains(line, "\r") {
			parts := strings.Split(line, "\r")
			lastPart := parts[len(parts)-1]
			processedLines = append(processedLines, lastPart)
		} else {
			processedLines = append(processedLines, line)
		}
	}

	// Join the processed lines back together with newline characters
	return strings.Join(processedLines, "\n")
}
