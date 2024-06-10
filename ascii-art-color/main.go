package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"ascii-artcolor/functions"
)

var colorMap = map[string]string{
	"black":   "30",
	"red":     "31",
	"green":   "32",
	"yellow":  "33",
	"blue":    "34",
	"magenta": "35",
	"cyan":    "36",
	"white":   "37",
}

func main() {
	ar := os.Args[1]
	if strings.HasPrefix(ar, "-") && !strings.HasPrefix(ar, "--") {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("EX: go run . --output=<fileName.txt> something standard")
		fmt.Printf("EXPECING: --")
		os.Exit(0)
	}
	if strings.HasPrefix(ar, "--") && !strings.Contains(ar, "=") {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("EX: go run . --output=<fileName.txt> something standard")
		fmt.Println("EXPECTING: =")
		os.Exit(0)
	}
	// Define the output flag
	//outputFlag := flag.String("output", "", "Specify the output file name.")
	colorFlag := flag.String("color", "", "Specify the color")
	subsrtFlag := flag.String("substr", "", "specify the substring color")
	// Parse the flags
	flag.Parse()

	// Custom usage message to inform the user how to use the program
	usageMessage := "Usage: go run . --output=<fileName.txt> [STRING] [BANNER]"

	// Get the non-flag arguments
	args := flag.Args()

	// Validate the number of arguments. Expect 1 or 2 arguments: the input string and optionally the font file
	if len(args) == 0 || len(args) > 2 {
		log.Fatal(usageMessage)
		return
	}

	// Assign the first argument to inputString
	inputString := args[1]
	substr := args[0]

	var fontFile string
	// If there is a second argument, assign it to fontFile; otherwise, use "standard" as the default font
	if len(args) == 2 {
		fontFile = args[1]
	} else {
		fontFile = "standard"
	}

	// Load the font using the FontLoader function from the "functions" package
	fontArray, err := functions.FontLoader(fontFile)
	if err != nil {
		log.Fatal(err)
	}

	// Replace "\\n" in the input string with actual newline characters
	inputString = strings.ReplaceAll(inputString, "\\n", "\n")
	// Split the input string into words using newline as the delimiter
	words := strings.Split(inputString, "\n")
	// Create a slice to hold the concatenated lines of ASCII art

	// Open the output file if specified
	// var outputFile *os.File
	// if *outputFlag != "" {
	// 	outputFile, err = os.Create(*outputFlag)
	// 	if err != nil {
	// 		log.Fatal("Error creating output file:", err)
	// 	}
	// 	defer outputFile.Close() // Close the file when main() exits
	// } else {
	// 	// If no output file is specified, use os.Stdout (standard output)
	// 	outputFile = os.Stdout
	// }
	// Process the color flag
	var colorCode string
	var ok bool
	// ar colorName string

	if *colorFlag != "" {
		colorCode, ok = colorMap[*colorFlag]
		if !ok {
			log.Fatal("Color not found")
		}
	}
	if *subsrtFlag == "" {
		fmt.Println("No substring specified")
		os.Exit(0)
	}

	count := 0 // Initialize a count for empty lines
	for _, word := range words {
		// Handle empty lines by printing a newline in the output
		if word == "" {
			count++
			if count < len(words) {
				fmt.Println() // Print newline to the output file
				continue
			}
			continue
		}
		// Create a slice to hold the concatenated lines of ASCII art for this word
		concatenatedLines := make([]string, 8)

		// Process each character in the word
		for _, char := range word {
			asciiChar, err := functions.PrintChar(char, fontArray)
			if err != nil {
				log.Fatal("Error printing ascii", err)
			}
			// Concatenate the lines of the ASCII art for the character
			for i, line := range asciiChar {
				concatenatedLines[i] += line
			}
		}
		if *colorFlag != "" {
			concatenatedLines = functions.Colorstring(concatenatedLines, colorCode)
			concatenatedLines = functions.SubstringColoring(inputString, substr, colorCode)

		}
		// Print the concatenated lines to the output file
		for _, lines := range concatenatedLines {
			fmt.Println(lines) // Print ASCII art line to the output file
		}
	}

	// // If an output file is specified, infoSrm the user
	// if *outputFlag != "" {
	// 	fmt.Println("Output written to", *outputFlag)
	// }
}
