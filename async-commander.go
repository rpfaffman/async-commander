// OBJECTIVES

// label processes
// - send commands to those processes
// - mute/unmute those processes

// have outputs of process go through channel

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Process struct {
	identifier string
	command    string
	channel    chan string
}

var PROCESSES map[string]Process = make(map[string]Process)

func main() {
	c := fan_in()

	for {
		select {
		case output := <-c:
			fmt.Println(output)
		}
		fmt.Print(prompt())
	}
}

func create_process(command string) Process {
	identifier := strconv.Itoa(len(PROCESSES) + 1)
	channel := exe_cmd(command, identifier)
	process := Process{identifier, command, channel}
	PROCESSES[identifier] = process
	return process
}

// This function should create processes and keep
// track of them.  Maybe build a map of identifier to
// process struct.
//
// It also maintains its own output channel through
// which all output is fanned in.
func fan_in() <-chan string {
	c := make(chan string)

	// processes := make(map[string]Process)

	scanner := bufio.NewScanner(os.Stdin)
	var input string

	go func() {
		for scanner.Scan() {
			input = scanner.Text()

			// currently this channel is not really used and the process
			// prints directly to stdout.  we need to find a way to pipe
			// this into the channel.

			process := create_process(input)

			go func() {
				c <- <-process.channel
			}()
		}
	}()

	return c
}

func exe_cmd(cmd, identifier string) chan string {
	c := make(chan string)

	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	go func() {
		if head == "cd" {
			os.Chdir(parts[0])

			c <- fmt.Sprintln()
		} else if head == "send" {
			process_identifier := parts[0]
			cmd := strings.Join(parts[1:], " ")
			send_input_to(process_identifier, cmd)
		} else {
			// right now, output does not go into the channel;
			// it prints directly to the stdout.  find a way to
			// send it into the channel.
			cmd := exec.Command(head, parts...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stdout
			cmd.Run()
			c <- fmt.Sprintln()

			go func() {
				var input string
				for {
					input = <-c
					run_cmd(input)
				}
			}()
		}
		// listen for additional commands to this process
		// delete the process after it is complete
		delete(PROCESSES, identifier)
	}()

	return c
}

// does not keep track of channels or attempt ot route special commands
// like 'cd' or 'send'
func run_cmd(command string) {
	parts := strings.Fields(command)
	head := parts[0]
	parts = parts[1:len(parts)]
	cmd := exec.Command(head, parts...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Run()
}

func send_input_to(identifier, input string) {
	if process, ok := PROCESSES[identifier]; ok {
		channel := process.channel
		channel <- fmt.Sprint(input)
	} else {
	}
}

func prompt() string {
	var prompt []string
	for k, v := range PROCESSES {
		prompt = append(prompt, fmt.Sprintf("[%s - %s]", k, v.command))
	}

	wd, _ := os.Getwd()

	prompt = append(prompt, fmt.Sprintf("\n%s >>> ", wd))

	return strings.Join(prompt, " ")
}
