package execute

import (
	"fmt"
	"os/exec"
)

type Execute struct {
}

func NewExecute() *Execute {
	return &Execute{}
}

// ExecuteCommand
func (e *Execute) ExecuteCommand(command string) (string, error) {

	fmt.Println("executing command: ", command)

	out, err := exec.Command("/bin/sh", command).Output()
	if err != nil {
		fmt.Printf("error command %s", err)
		return "", err
	}
	return string(out), nil
}
