package app

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// CheckSystemApps checks system apps are in path
func CheckSystemApps() {
	errors := []string{}

	if !inPath("sudo") {
		errors = append(errors, "- missing sudo: install `sudo`")
	}

	if !inPath("wg") {
		errors = append(errors, "- missing wg: install `wireguard`")
	}

	if !inPath("wg-quick") {
		errors = append(errors, "- missing wg-quick: install `wireguard-tools`")
	}

	if !inPath("resolvconf") {
		resolvectl, err := which("resolvectl")
		if err == nil {
			errors = append(errors, fmt.Sprintf("- create a symlink: `sudo ln -s %s /usr/local/bin/resolvconf`", resolvectl))
		} else {
			errors = append(errors, "- missing resolvconf: install `resolvconf`")
		}
	}

	if len(errors) > 0 {
		fmt.Printf("You are missing some system tools:\n%s\n", strings.Join(errors, "\n"))
		os.Exit(2)
	}
}

// Which locates a binary in the current $PATH.
func which(binName string) (string, error) {
	return exec.LookPath(binName)
}

// InPath returns whether a binary is if your path
func inPath(binName string) bool {
	if _, err := exec.LookPath(binName); err != nil {
		return false
	}

	return true
}
