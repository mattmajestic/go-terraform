package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the database connection string from the .env file
	connStr := os.Getenv("DATABASE_URL")
	// Open a connection to the database
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func logScriptRun(scriptName, status string) {
	// Determine the operating system for a simplified hostname
	var hostname string
	if runtime.GOOS == "windows" {
		hostname = "Windows"
	} else if runtime.GOOS == "linux" {
		hostname = "Linux"
	} else {
		hostname = "Other"
	}

	// Insert the log into the database
	_, err := db.Exec("INSERT INTO terraform (hostname, script_name, status) VALUES ($1, $2, success)", hostname, scriptName, status)
	if err != nil {
		log.Printf("Error logging script run: %v", err)
	}
}

func runTerraformCommands() {
    // Define the Terraform commands to run
    commands := [][]string{
        {"init"},
        {"apply", "-auto-approve"},
        {"destroy", "-auto-approve"},
    }

    for _, command := range commands {
        cmd := exec.Command("terraform", command...)
        err := cmd.Run()
        if err != nil {
            log.Printf("Terraform command '%s' failed: %v", command[0], err)
            return
        }
    }
}

func main() {
	// Initialize the Gin router
	r := gin.Default()

	// Define a simple GET route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// Execute Terraform commands and log the outcome
	runTerraformCommands()

	// Start the Gin server on port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Unable to start server:", err)
		logScriptRun("terraform.go", "Failure")
	}
}