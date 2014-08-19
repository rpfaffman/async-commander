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
			output_channel := exe_cmd(input)

			// currently this channel is not really used and the process
			// prints directly to stdout.  we need to find a way to pipe
			// this into the channel.
			go func() {
				c <- <-output_channel
			}()
		}
	}()

	return c
}

func exe_cmd(cmd string) <-chan string {
	c := make(chan string)

	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	go func() {
		if head == "cd" {
			os.Chdir(parts[0])
			c <- fmt.Sprint()
		} else {
			// write now, output does not go into the channel;
			// it prints directly to the stdout.  find a way to
			// send it into the channel.
			cmd := exec.Command(head, parts...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stdout
			cmd.Run()
			c <- fmt.Sprintln()
		}
	}()

	return c
}

func prompt() string {
	wd, _ := os.Getwd()
	return fmt.Sprintf("%s >>> ", wd)
}
