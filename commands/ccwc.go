package commands

import (
	"fmt"
	"os"
	"strings"
)

// Emulating wc command
// wc command is used to count the number of lines, words and characters in a file
// wc -c test.txt

func Call(input string) bool {
	commands := strings.Fields(strings.TrimSpace(input))
	if commands[0] == "wc" {
		if commands[1] == "-c" {
			getFileSize(commands[2])
			return true
		}
		if commands[1] == "-l" {
			countLines(commands[2])
			return true
		}
	}
	return false
}

func countLines(filepath string) {
	var err error
	f, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading file")
	}
	lines := strings.Split(string(f), "\n")
	fmt.Printf("%s Lines : %d", filepath, len(lines))
}

func getFileSize(filepath string) {
	var err error
	var f *os.File
	f, err = os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file")
	}
	info, err := f.Stat()
	if err != nil {
		fmt.Println("Error getting file info")
	}
	fmt.Printf("%s : %d", filepath, info.Size())
}
