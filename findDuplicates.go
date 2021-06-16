package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	filename := flag.String("file", "", "file to read from")
	flag.Parse()

	var result map[string][]int
	if *filename != "" {
		result = collectLineOccurrencesInFile(*filename)
	} else {
		result = collectLineOccurrencesInStdin()
	}

	printResult(result)
}

// Collects occurrences of lines in file with filename.
// Returns map with line content as key and array of line numbers in which the line occurred as value.
func collectLineOccurrencesInFile(filename string) map[string][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	return collectLineOccurrences(bufio.NewReader(file))
}

// Collects occurrences of lines in stdin.
// Returns map with line content as key and array of line numbers in which the line occurred as value.
func collectLineOccurrencesInStdin() map[string][]int {
	return collectLineOccurrences(bufio.NewReader(os.Stdin))
}

// Prints the lineOccurrences which occurred more than one time.
func printResult(lineOccurrences map[string][]int) {
	for lineContent, occurrences := range lineOccurrences {
		if len(occurrences) < 2 {
			continue
		}

		fmt.Printf("line '%v' has been found in multiple lines: %v\n", lineContent, occurrences)
	}
}

// Collects occurrences of lines provided by reader.
// Returns map with line content as key and array of line numbers in which the line occurred as value.
func collectLineOccurrences(reader *bufio.Reader) map[string][]int {
	lineOccurrences := make(map[string][]int)

	lineNumber := 1
	for {
		var lineContent, err = reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		lineContent = strings.TrimSpace(lineContent)

		var lineNumbers, isPresent = lineOccurrences[lineContent]
		if !isPresent {
			lineNumbers = make([]int, 0)
		}

		lineNumbers = append(lineNumbers, lineNumber)
		lineOccurrences[lineContent] = lineNumbers

		lineNumber++
	}

	return lineOccurrences
}
