package unit

import (
	"github.com/IshaanNene/GoShell/internal/core"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"testing"
)

// Helper function to execute a Cobra command with arguments and return the output
func executeCommand(cmd *cobra.Command, args ...string) (string, error) {
	cmd.SetArgs(args)
	output, err := cmd.ExecuteC()
	if err != nil {
		return "", err
	}
	return output.UsageString(), nil
}

func TestTouchCommand(t *testing.T) {
	fileName := "testfile.txt"

	// Test file creation
	if _, err := executeCommand(core.TouchCmd, fileName); err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}

	// Test if the file exists
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		t.Fatalf("Expected file %s to exist but it does not", fileName)
	}

	// Clean up after test
	os.Remove(fileName)
}

func TestLsCommand(t *testing.T) {
	// Setup: Create some files to list
	file1 := "file1.txt"
	file2 := "file2.txt"
	os.Create(file1)
	os.Create(file2)
	defer os.Remove(file1)
	defer os.Remove(file2)

	// Expected output should be the filenames listed
	expected := "file1.txt\nfile2.txt\n"

	// Call the ls command
	got, err := executeCommand(core.LsCmd)
	if err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}

	// Check if the output matches expected filenames
	if got != expected {
		t.Errorf("Expected %q but got %q", expected, got)
	}
}

func TestPwdCommand(t *testing.T) {
	expected, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting the current directory: %v", err)
	}

	// Call the pwd command
	got, err := executeCommand(core.PwdCmd)
	if err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}

	// Check if the output matches the expected directory
	if got != expected {
		t.Errorf("Expected %q but got %q", expected, got)
	}
}

func TestCdCommand(t *testing.T) {
	initialDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting the initial directory: %v", err)
	}

	// Test change to a different directory
	newDir := filepath.Join(initialDir, "..")

	// Simulate running the cd command
	if _, err := executeCommand(core.CdCmd, newDir); err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}

	// Verify the current directory has changed
	expected, err := filepath.Abs(newDir)
	if err != nil {
		t.Fatalf("Error getting the expected directory: %v", err)
	}

	got, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting the current directory: %v", err)
	}
	if got != expected {
		t.Errorf("Expected %q but got %q", expected, got)
	}

	// Change back to the original directory
	if _, err := executeCommand(core.CdCmd, initialDir); err != nil {
		t.Fatalf("Failed to return to initial directory: %v", err)
	}
}

func TestMkdirCommand(t *testing.T) {
	dirName := "testdir"

	// Test directory creation
	if _, err := executeCommand(core.MkdirCmd, dirName); err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}

	// Test if the directory exists
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		t.Fatalf("Expected directory %s to exist but it does not", dirName)
	}

	// Clean up after test
	os.Remove(dirName)
}

// Additional tests for other functionalities can be added here
