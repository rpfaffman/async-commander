package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	c := read_input()

	fmt.Print(prompt())

	for {
		select {
		case output := <-c:
			fmt.Println(output)
		}
		fmt.Print(prompt())
	}
}

func read_input() <-chan string {
	c := make(chan string)
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	go func() {
		for scanner.Scan() {
			input = scanner.Text()
			c <- exe_cmd(input)
		}
	}()
	return c
}

func exe_cmd(cmd string) string {
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	if head == "cd" {
		os.Chdir(parts[0])
		return fmt.Sprintln()
	} else {
		out, err := exec.Command(head, parts...).Output()
		if err != nil {
			return fmt.Sprintf("%s", err)
		}
		return fmt.Sprintf("%s", out)
	}
}

func prompt() string {
	wd, _ := os.Getwd()
	return fmt.Sprintf("%s >>> ", wd)
}
