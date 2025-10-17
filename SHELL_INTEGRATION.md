# Shell Integration for GoShell

## The `cd` Command Problem

The `cd` command in GoShell requires special integration with your shell because a child process (GoShell) cannot change the working directory of its parent process (your shell).

## Solution: Shell Wrapper Function

We provide a wrapper function that integrates GoShell with your shell, allowing the `cd` command to work properly.

## Installation

### Step 1: Add the wrapper to your shell configuration

**For Bash users** (add to `~/.bashrc`):
```bash
# GoShell integration
source /path/to/GoShell/goshell-wrapper.sh
```

**For Zsh users** (add to `~/.zshrc`):
```zsh
# GoShell integration
source /path/to/GoShell/goshell-wrapper.sh
```

### Step 2: Reload your shell configuration
```bash
# For Bash
source ~/.bashrc

# For Zsh
source ~/.zshrc
```

### Step 3: Test it!
```bash
goshell cd /tmp
# You should see: Directory changed to: /tmp
# And pwd should show you're actually in /tmp

pwd
# Output: /tmp
```

## How It Works

1. When you run `goshell cd <directory>`, the wrapper function intercepts the command
2. GoShell validates the directory and outputs a special marker: `GOSHELL_CD:/path/to/dir`
3. The wrapper detects this marker and executes `cd` in your current shell
4. Your shell actually changes directories! ✨

## Alternative Usage (Without Wrapper)

If you don't want to use the wrapper, you can use:
```bash
cd $(goshell cd docs)
```

But this is less user-friendly and doesn't provide error messages as nicely.

## Troubleshooting

### "goshell: command not found"
The wrapper looks for the `goshell` binary in:
1. The current directory (`./goshell`)
2. Your system PATH

Make sure you've built GoShell and either:
- Run it from the GoShell directory, or
- Install it to a directory in your PATH

### Directory not changing
1. Make sure you've sourced the wrapper: `source ~/.bashrc`
2. Check that the wrapper function is loaded: `type goshell`
3. Verify the binary path in `goshell-wrapper.sh` is correct
