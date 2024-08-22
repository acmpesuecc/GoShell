// tests/integration/goshell_integration_test.go

package integration

import (
	"bytes"
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
	}{
		{
			name:           "Test 'ls' command",
			args:           []string{"ls"},
			expectedOutput: "Expected output from 'ls' command\n", // Update as necessary
		},
		{
			name:           "Test 'pwd' command",
			args:           []string{"pwd"},
			expectedOutput: "/path/to/current/dir\n", // Update as necessary
		},
		{
			name:           "Test 'touch' command",
			args:           []string{"touch", "testfile.txt"},
			expectedOutput: "file creation success\n",
		},
		{
			name:           "Test 'rm' command",
			args:           []string{"rm", "testfile.txt"},
			expectedOutput: "file deletion success\n", // Update as necessary
		},
		{
			name:           "Test 'mkdir' command",
			args:           []string{"mkdir", "testdir"},
			expectedOutput: "directory creation success\n", // Update as necessary
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Run the command
			cmd := exec.Command(binaryPath, tt.args...)
			var out bytes.Buffer
			cmd.Stdout = &out
			var stderr bytes.Buffer
			cmd.Stderr = &stderr
			err := cmd.Run()
			if err != nil {
				t.Fatalf("Error executing command: %v, stderr: %s", err, stderr.String())
			}
			if got := out.String(); got != tt.expectedOutput {
				t.Errorf("Expected %q but got %q", tt.expectedOutput, got)
			}
		})
	}
}
