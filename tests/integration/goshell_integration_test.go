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
			expectedOutput: "Expected output from 'ls' command",
		},
		{
			name:           "Test 'pwd' command",
			args:           []string{"pwd"},
			expectedOutput: "Expected output from 'pwd' command",
		},
		{
			name:           "Test 'touch' command",
			args:           []string{"touch", "testfile.txt"},
			expectedOutput: "file creation success",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(binaryPath, tt.args...)
			var out bytes.Buffer
			cmd.Stdout = &out
			err := cmd.Run()
			if err != nil {
				t.Fatalf("Error executing command: %v", err)
			}
			if got := out.String(); got != tt.expectedOutput {
				t.Errorf("Expected %q but got %q", tt.expectedOutput, got)
			}
		})
	}
}
