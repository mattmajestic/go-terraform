package main_test

import (
	"os/exec"
	"testing"
)

// runMainProgram executes the main.go program.
func runMainProgram() error {
	cmd := exec.Command("go", "run", "main.go")
	return cmd.Run()
}

func TestMainProgramExecutionAndTerraformOperations(t *testing.T) {
	// Run the main program
	if err := runMainProgram(); err != nil {
		t.Fatalf("Failed to run the main program: %s\n", err)
	}

	// Run Terraform apply
	if err := runCommand("terraform", "apply", "-auto-approve"); err != nil {
		t.Fatalf("Failed to apply Terraform changes: %s\n", err)
	}

	// Check if apply was successful (check for any output, state file, etc.)
	// If apply was successful, proceed with destroy
	// You need to implement this check based on your project's specific requirements

	// Run Terraform destroy
	if err := runCommand("terraform", "destroy", "-auto-approve"); err != nil {
		t.Fatalf("Failed to destroy Terraform resources: %s\n", err)
	}
}

// runCommand is a helper function to execute shell commands.
func runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	return cmd.Run()
}
