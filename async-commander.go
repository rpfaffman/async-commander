package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// func main() {
// 	var input string
// 	for {
// 		fmt.Scanf("%s", &input)
// 		input = strings.Join([]input, " ")
// 		fmt.Println("EXE:", input)
// 		//input = strings.Join(input, " ")
// 		//exe_cmd(input)
// 	}
// }

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	fmt.Print(prompt())

	for scanner.Scan() {
		input = scanner.Text()
		fmt.Println(input)
		exe_cmd(input)
		fmt.Print(prompt())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func exe_cmd(cmd string) {
	// splitting head => g++ parts => rest of the command
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	if head == "cd" {
		fmt.Println("You tried to change dirs", parts[0])
		os.Chdir(parts[0])
		return
	} else {
		out, err := exec.Command(head, parts...).Output()
		if err != nil {
			fmt.Printf("%s", err)
		}
		fmt.Printf("%s", out)
	}

	// wg.Done() // Need to signal to waitgroup that this goroutine is done
}

func prompt() string {
	wd, _ := os.Getwd()
	return fmt.Sprintf("%s >>> ", wd)
}
