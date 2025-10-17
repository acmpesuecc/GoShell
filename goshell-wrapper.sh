#!/bin/bash
# GoShell wrapper function for proper cd command integration
# 
# Installation:
# Add this to your ~/.bashrc or ~/.zshrc:
# source /path/to/goshell-wrapper.sh

goshell() {
    # Path to the actual goshell binary
    local goshell_bin="./goshell"
    
    # If goshell is installed globally, use that instead
    if command -v goshell &> /dev/null; then
        goshell_bin="goshell"
    fi
    
    # Execute goshell and capture output
    local output
    output=$("$goshell_bin" "$@" 2>&1)
    local exit_code=$?
    
    # Check if this is a cd command response
    if [[ "$output" == GOSHELL_CD:* ]]; then
        # Extract the directory path
        local new_dir="${output#GOSHELL_CD:}"
        
        # Change directory in the current shell
        if cd "$new_dir" 2>/dev/null; then
            echo "Directory changed to: $new_dir"
            return 0
        else
            echo "Error: Failed to change directory to $new_dir" >&2
            return 1
        fi
    else
        # Not a cd command, just print the output
        echo "$output"
        return $exit_code
    fi
}

# Export the function so it's available in the shell
export -f goshell 2>/dev/null || true
