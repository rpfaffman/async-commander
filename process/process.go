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
	processes map[string]*Process
}

func NewProcessManager() *ProcessManager {
	return &ProcessManager{
		processes: make(map[string]*Process),
	}
}

func (pm *ProcessManager) Spawn(command string) *Process {
	identifier := strconv.Itoa(len(pm.processes) + 1)
	process := &Process{
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

func (pm *ProcessManager) List() map[string]*Process {
	return pm.processes
}

func (pm *ProcessManager) SendInput(identifier string, command string) {
	pm.processes[identifier].Input(fmt.Sprintf("%s\n", command))
}

func (pm *ProcessManager) RetrieveProcess(identifier string) *Process {
	return pm.processes[identifier]
}

// ===PROCESS CLASS===
type Process struct {
	identifier string
	command    string
	inputPipe  io.WriteCloser
	manager    *ProcessManager
}

func (p *Process) Execute() {
	parts := strings.Fields(p.command)
	cmd := exec.Command(parts[0], parts[1:]...)
	var err error
	p.inputPipe, err = cmd.StdinPipe()
	if err != nil {
		fmt.Println(err)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	go func() {
		cmd.Start()
		cmd.Wait()
		p.finish()
	}()
}

func (p *Process) Input(input string) {
	_, err := p.inputPipe.Write([]byte(input))
	if err != nil {
		fmt.Println(err)
	}
}

func (p *Process) finish() {
	p.manager.Remove(p.identifier)
	fmt.Printf("\nProcess %s finished.\n> ", p.identifier)
}

func (p *Process) Identifier() string {
	return p.identifier
}

func (p *Process) Command() string {
	return p.command
}
