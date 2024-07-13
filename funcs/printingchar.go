package ascii

import (
	"bufio"
	"fmt"
	"os"
)

func Printingchar(r rune, file *os.File) []string {
	n := int(r - 32)
	startLine := 3 + (n)*9
	// fmt.Println(startLine)
	endLine := startLine + 8
	// Seek to the beginning of the file
	_, err := file.Seek(0, 0)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(10)
	}
	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	// Scan lines from the file
	lineNumber := 1
	var lines []string
	for scanner.Scan() {
		lineNumber++
		if lineNumber < startLine {
			continue // Skip until the start line }
		}
		if lineNumber > endLine {
			break // Stop scanning after the end line
		}
		line := scanner.Text() // Get the current line
		// fmt.Println(line)      // Process the line (print in this case)
		lines = append(lines, line)
	}
	// Check for errors encountered during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
	return lines
}
