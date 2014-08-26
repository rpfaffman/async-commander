package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/theelectricmiraclecat/async-commander/process"
)

var ProcessManager = process.NewProcessManager()

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	var parts []string
	var head string

	for scanner.Scan() {
		input = scanner.Text()
		parts = strings.Fields(input)

		if len(parts) == 0 {
			printPrompt()
		} else {
			head = parts[0]
			if head == "cd" { // change directories
				os.Chdir(parts[1])
				printPrompt()
			} else if head == "send" { // send input to pre-existing process
				input = fmt.Sprintf("%s\n", strings.Join(parts[2:], " "))
				ProcessManager.SendInput(parts[1], input)
			} else { // regular command
				runCommand(input)
			}
		}

	}
}

func runCommand(cmd string) {
	ProcessManager.Spawn(cmd)
}

func printPrompt() {
	wd, _ := os.Getwd()

	fmt.Println(ProcessManager.List())
	fmt.Println(wd)
	fmt.Printf("> ")
}
