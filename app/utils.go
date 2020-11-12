package app

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)

// SudoExec will run a particular comand and return the output or an error.
// All commands are run via sudo.
func SudoExec(command string) (string, error) {
	cmd := exec.Command("sh", "-c", "sudo "+command)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf(strings.TrimSpace(string(output)))
	}

	return string(output), nil
}

// Select will present a list of choices
func Select(question string, choices []string) string {
	prompt := promptui.Select{
		Label:    question,
		Items:    choices,
		HideHelp: true,
	}

	_, result, _ := prompt.Run()

	return result

}

// InArray returns if a string is in a string slice
func InArray(val string, array []string) bool {
	for i := range array {
		if ok := array[i] == val; ok {
			return true
		}
	}
	return false
}
