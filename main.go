package main

import (
	"bufio"
	"fmt"
	"go-1/commands"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	if !commands.Call(line) {
		fmt.Println("Invalid command")
	}
}
