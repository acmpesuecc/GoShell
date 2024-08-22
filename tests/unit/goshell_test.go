// tests/unit/goshell_test.go

package unit

import (
	"github.com/IshaanNene/GoShell/internal/core"
	"os"
	"testing"
)

func TestTouchCommand(t *testing.T) {
	// Setup for testing the 'touch' command
	fileName := "testfile.txt"
	if err := core.Touch(fileName); err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		t.Fatalf("Expected file %s to exist but it does not", fileName)
	}
	os.Remove(fileName) // Clean up
}

func TestLsCommand(t *testing.T) {
	// Example setup for testing the 'ls' command
	// You would need to implement the actual logic in your core package
	expected := "some_expected_output"
	got := core.ListFiles() // Assuming you have a function that lists files
	if got != expected {
		t.Errorf("Expected %q but got %q", expected, got)
	}
}

func TestPwdCommand(t *testing.T) {
	// Example setup for testing the 'pwd' command
	expected := "/path/to/current/dir"       // Update as necessary
	got, err := core.PrintWorkingDirectory() // Assuming you have a function for this
	if err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}
	if got != expected {
		t.Errorf("Expected %q but got %q", expected, got)
	}
}

// Add more tests for other functionalities or methods
