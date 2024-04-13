package main

import (
    "os/exec"
)

// runCommand is a helper function to execute shell commands.
func runCommand(command string, args ...string) error {
    cmd := exec.Command(command, args...)
    return cmd.Run()
}
