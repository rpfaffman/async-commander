package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/theelectricmiraclecat/async-commander/process"
)

var processManager = process.NewProcessManager()
var defaultProcess = ""

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
			} else if head == "switch" { // switch default input
				if len(parts) > 1 {
					defaultProcess = parts[1]
				} else {
					defaultProcess = ""
				}
			} else if head == "send" { // send input to pre-existing process
				input = strings.Join(parts[2:], " ")
				processManager.SendInput(parts[1], input)
			} else { // regular command
				runCommand(input)
			}
		}

	}
}

func runCommand(cmd string) {
	if defaultProcess != "" && processManager.RetrieveProcess(defaultProcess) != nil {
		processManager.SendInput(defaultProcess, cmd)
	} else {
		processManager.Spawn(cmd)
	}
}

func printPrompt() {
	processes := processManager.List()
	wd, _ := os.Getwd()

	if defaultProcess != "" {
		fmt.Printf("%s ) ", defaultProcess)
	}
	for _, process := range processes {
		fmt.Printf("[ %s - %s ] ", process.Identifier(), process.Command())
	}
	fmt.Println()
	fmt.Println(wd)
	fmt.Printf("> ")
}
