package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Emulating wc command
// wc command is used to count the number of lines, words and characters in a file
// wc -c test.txt

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	if !call(line) {
		fmt.Println("Invalid command")
	}
}

func call(input string) bool {
	commands := strings.Fields(strings.TrimSpace(input))
	if commands[0] == "wc" {
		if commands[1] == "-c" {
			getFileSize(commands[2])
			return true
		}
	}
	return false
}

func getFileSize(filepath string) {
	var err error
	var f *os.File
	f, err = os.Open(filepath)
	info, err := f.Stat()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s : %d", filepath, info.Size())
}
