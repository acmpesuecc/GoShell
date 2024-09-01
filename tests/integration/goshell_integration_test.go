package integration

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
)

func TestCommandExecution(t *testing.T) {
	// Path to the built binary
	binaryPath := "./goshell"

	tests := []struct {
		name           string
		args           []string
		expectedOutput string
		postCheck      func(t *testing.T) // Additional checks after command execution
		cleanup        func()             // Cleanup actions after each test
	}{
		{
			name:           "Test 'ls' command",
			args:           []string{"ls"},
			expectedOutput: "goshell\n", // Adjust according to the actual output
		},
		{
			name: "Test 'pwd' command",
			args: []string{"pwd"},
			expectedOutput: func() string {
				dir, _ := os.Getwd()
				return dir + "\n"
			}(),
		},
		{
			name: "Test 'touch' command",
			args: []string{"touch", "testfile.txt"},
			postCheck: func(t *testing.T) {
				// Verify the file was created
				if _, err := os.Stat("testfile.txt"); os.IsNotExist(err) {
					t.Fatalf("Expected file 'testfile.txt' to exist but it does not")
				}
			},
			cleanup: func() {
				// Clean up the created file
				os.Remove("testfile.txt")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(binaryPath, tt.args...)
			var out bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &out

			// Execute the command
			err := cmd.Run()
			if err != nil {
				t.Fatalf("Error executing command: %v", err)
			}

			// Check the output if expectedOutput is provided
			if tt.expectedOutput != "" {
				if got := out.String(); got != tt.expectedOutput {
					t.Errorf("Expected %q but got %q", tt.expectedOutput, got)
				}
			}

			// Execute any post-check actions
			if tt.postCheck != nil {
				tt.postCheck(t)
			}

			// Execute cleanup actions
			if tt.cleanup != nil {
				tt.cleanup()
			}
		})
	}
}
