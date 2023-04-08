# Linter for Unreal Engine

This linter finds bad smells in an Unreal Project.

## Run

1. Install dependencies

   This tool needs Doxygen and Golang. Use your favourite package manager for this.

   Example for macOS:

    ```bash
    brew install doxygen
    brew install go
    ```

2. Run

    ```bash
    go run .
    # or
    go run . -path="the/path/to/your/unreal/project"
    ```
    