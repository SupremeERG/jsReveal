# How jsReveal Works

This document provides a high-level overview of `jsReveal`'s internal architecture and data flow.

## Core Components

The tool is structured into several key packages:

-   **`main`**: The entry point of the application. It orchestrates the entire process, from parsing options to managing output.
-   **`runner`**: Responsible for parsing command-line arguments and setting up the configuration for the run.
-   **`internal/parser`**: Contains the core logic for fetching, parsing, and analyzing JavaScript code.
-   **`pkg/fetchcode`**: A utility package for fetching content, such as JavaScript from URLs or regex patterns from files.
-   **`pkg/regexmod`**: A helper package for compiling and managing the regex patterns used for matching.

## Data Flow

The process can be broken down into the following steps:

1.  **Parsing Options**: When you run `jsReveal`, the `main` function first calls `runner.ParseOptions()`. This function, located in `runner/options.go`, uses Go's `flag` package to parse all the command-line arguments you provide. It determines the source of the JavaScript (file, URL, list, or stdin) and the type of analysis to perform (endpoints, API keys, etc.).

2.  **Initiating the Run**: Based on the parsed options, the `main` function calls the `run()` function. This function acts as a controller, directing the flow based on the input source.

3.  **Fetching JavaScript Code**:
    *   **Local File (`-f`)**: `parser.ParseJS` is called, which reads the file directly from the disk.
    *   **Remote URL (`-u`)**: `fetchcode.FetchJSFromURL` is called to download the JavaScript code as a string.
    *   **List of URLs (`-l`)**: `parser.ParseJSFromList` is called. It reads the list of URLs and spins up a pool of concurrent workers (`urlWorker`). Each worker fetches the JavaScript from a URL. This concurrent approach significantly speeds up the analysis of multiple files.
    *   **Stdin**: The `run` function reads from `os.Stdin` line by line, treating each line as a URL to be fetched.

4.  **Fetching Regex Patterns**: The parser functions call `fetchcode.FetchPatterns()` to read the regex patterns from the appropriate file (`endpoints.txt`, `api_key_regex.txt`, etc.).

5.  **Pattern Matching**: The core analysis happens in the `parser.parse()` function. It receives the JavaScript code (as a string) and the list of regex patterns. It iterates through each pattern and uses the `regexp2` library—a more powerful regex engine than Go's native one—to find all matching substrings within the code.

6.  **Handling Output**: As matches are found, they are sent back to the `main` function via a Go channel (`outputChannel`). The `main` function then handles the output based on your flags:
    *   If the `-o` flag is used, the results are written to a JSON file.
    *   If the `--pretty` flag is used, the results are printed to the console in a formatted, human-readable way.
    *   Otherwise, the results are printed to the console in the default `match::::source` format.

7.  **Concurrency**: The use of Go channels and goroutines is central to `jsReveal`'s design. This allows for non-blocking I/O operations (like fetching from multiple URLs) and efficient processing, making the tool fast and responsive.
