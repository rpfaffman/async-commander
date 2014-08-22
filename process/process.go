// & ampersand creates a point to the object immediately after it
// * indicates that the argument is of type pointer.  dereferences a pointer
// pointers can be nil

package process

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// ===PROCESS MANAGER CLASS===
type ProcessManager struct {
	processes map[string]Process
}

func NewProcessManager() *ProcessManager {
	return &ProcessManager{
		processes: make(map[string]Process),
	}
}

func (pm *ProcessManager) Spawn(command string) Process {
	identifier := strconv.Itoa(len(pm.processes) + 1)
	process := Process{
		identifier: identifier,
		command:    command,
		manager:    pm,
	}
	pm.processes[identifier] = process

	process.Execute()

	return process
}

func (pm *ProcessManager) Remove(identifier string) {
	delete(pm.processes, identifier)
}

func (pm *ProcessManager) List() map[string]Process {
	return pm.processes
}

func (pm *ProcessManager) SendInput(identifier string, command string) {
	pm.processes[identifier].Input(command)
}

// ===PROCESS CLASS===
type Process struct {
	identifier string
	command    string
	inputPipe  io.WriteCloser
	manager    *ProcessManager
}

func (p Process) Execute() {
	parts := strings.Fields(p.command)
	cmd := exec.Command(parts[0], parts[1:]...)
	p.inputPipe, _ = cmd.StdinPipe()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	go func() {
		cmd.Start()
		cmd.Wait()
		p.finish()
	}()
}

// need to implement
func (p Process) Input(input string) {
	fmt.Printf("Received input for %s: %s\n", p.identifier, input)
	parsedInput := []byte(input)
	p.inputPipe.Write(parsedInput)
}

func (p Process) finish() {
	p.manager.Remove(p.identifier)
	fmt.Printf("Process %s finished.\n", p.identifier)
}
