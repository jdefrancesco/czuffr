package zfuzz

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Chatgpt pseudo code to remind me

func runFuzzing(targetProgram string, targetArgs []string) {
	for _, arg := range targetArgs {
		if arg == "FUZZ" {
			// Generate fuzz payload
			fuzzPayload := "AAAA" // Placeholder for fuzzing input
			arg = fuzzPayload
		}

		// Prepare command execution
		cmd := exec.Command(targetProgram, targetArgs...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Run the command
		fmt.Println("Executing:", targetProgram, strings.Join(targetArgs, " "))
		if err := cmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Command failed with %s\n", err)
		}
	}
}
