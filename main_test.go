package main_test

import (
    "fmt"
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
        fmt.Println("\x1b[31m🔴 Failed to run the main program\x1b[0m")
        t.Fatalf("Error: %s\n", err)
    } else {
        fmt.Println("\x1b[32m🟢 Main program ran successfully\x1b[0m")
    }

    // Run Terraform apply
    if err := runCommand("terraform", "apply", "-auto-approve"); err != nil {
        fmt.Println("\x1b[31m🔴 Failed to apply Terraform changes\x1b[0m")
        t.Fatalf("Error: %s\n", err)
    } else {
        fmt.Println("\x1b[32m🟢 Terraform apply ran successfully\x1b[0m")
    }

    // Check if apply was successful (check for any output, state file, etc.)
    // If apply was successful, proceed with destroy
    // You need to implement this check based on your project's specific requirements

    // Run Terraform destroy
    if err := runCommand("terraform", "destroy", "-auto-approve"); err != nil {
        fmt.Println("\x1b[31m🔴 Failed to destroy Terraform resources\x1b[0m")
        t.Fatalf("Error: %s\n", err)
    } else {
        fmt.Println("\x1b[32m🟢 Terraform destroy ran successfully\x1b[0m")
    }
}

// runCommand is a helper function to execute shell commands.
func runCommand(command string, args ...string) error {
    cmd := exec.Command(command, args...)
    return cmd.Run()
}