package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    // Run Terraform init
    if err := runCommand("terraform", "init"); err != nil {
        fmt.Printf("Failed to run 'terraform init': %s\n", err)
        return
    }

    // Run Terraform plan
    if err := runCommand("terraform", "plan"); err != nil {
        fmt.Printf("Failed to run 'terraform plan': %s\n", err)
        return
    }

    // Run Terraform apply
    if err := runCommand("terraform", "apply", "-auto-approve"); err != nil {
        fmt.Printf("Failed to run 'terraform apply': %s\n", err)
        return
    }
}

func runCommand(command string, args ...string) error {
    cmd := exec.Command(command, args...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}
