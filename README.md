# AY SHELL (The Born again Bourne again shell)

AY SHELL is a simple, yet powerful, shell program written in Go that allows users to execute both built-in and external commands. It aims to provide a lightweight and efficient command-line interface.

## Features

*   **Built-in Commands:** AY SHELL comes with a set of essential built-in commands:
    *   `cd <directory>`: Change the current working directory.
    *   `ls [directory]`: List directory contents.
    *   `mv <source> <destination>`: Move or rename files and directories.
    *   `cat <file>`: Display file contents.
    *   `mkdir <directory>`: Create a new directory.
    *   `rm <file/directory>`: Remove files or directories.
    *   `touch <file>`: Create an empty file or update timestamps.
    *   `help`: Display help information about built-in commands.
    *   `exit`: Terminate the AY SHELL session.
*   **External Command Execution:** Execute any external command available in your system's PATH. AY SHELL handles argument parsing for these commands.
*   **Customizable Icon:** Includes an `aysh.ico` file and `ayshell.rc` resource file for embedding a custom icon into the executable (requires a resource compiler like `windres`).

## Building from Source

To build AY SHELL from source, you'll need Go installed on your system. If you want to embed the icon on Windows, you'll also need a resource compiler like `windres` (often part of MinGW/GCC).

1.  **Clone the repository (if applicable) or navigate to the project directory.**
2.  **Compile the resource file (for Windows icon embedding - optional):**
    ```bash
    windres ayshell.rc -o ayshell.syso
    ```
3.  **Build the Go program:**
    ```bash
    go build -o aysh.exe
    ```
    (On Linux/macOS, you might use `go build -o aysh`)

## Usage

Once built, you can run AY SHELL from your terminal:

```bash
./aysh.exe  # On Windows
./aysh      # On Linux/macOS
```

This will start the AY SHELL prompt, where you can begin typing commands.

## Icon Embedding Note

The project includes `aysh.ico` and `ayshell.rc`. To embed this icon into the final `aysh.exe` on Windows, you need to compile the `.rc` file into a `.syso` object file using a resource compiler like `windres` (commonly found in MinGW distributions). Then, the Go compiler will automatically link this resource file during the build process if it's present in the main package directory.

Example `ayshell.rc` content:
```
1 ICON "aysh.ico"
```