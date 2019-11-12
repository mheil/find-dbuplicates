package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()

	var reader *bufio.Reader
	if len(args) > 0 {
		reader = bufio.NewReader(strings.NewReader(args[0] + "\n"))
	} else {
		reader = bufio.NewReader(os.Stdin)
	}

	lineOccurences := findDuplicatedLines(reader)

	for lineContent, occurances := range lineOccurences {
		if len(occurances) < 2 {
			continue
		}

		fmt.Printf("line '%v' has been found in multiple lines: %v\n", lineContent, occurances)
	}
}

func findDuplicatedLines(reader *bufio.Reader) map[string][]int {
	lineOccurences := make(map[string][]int)

	lineNumber := 1
	for {
		var lineContent, err = reader.ReadString('\n')
		if err != nil {
			break
		}

		lineContent = strings.TrimSpace(lineContent)

		var lineNumbers, isPresent = lineOccurences[lineContent]
		if !isPresent {
			lineNumbers = make([]int, 0)
		}

		lineNumbers = append(lineNumbers, lineNumber)
		lineOccurences[lineContent] = lineNumbers

		lineNumber++
	}

	return lineOccurences
}
