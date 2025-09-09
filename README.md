# jsReveal

**Reveals interesting JS code to ethical hackers**

## About

`jsReveal` is a tool designed for security researchers and ethical hackers to analyze JavaScript files. It helps in discovering sensitive information such as API endpoints, API keys, and other interesting patterns within JavaScript code. You can provide a single JS file, a list of URLs to JS files, or even pipe URLs via stdin.

## Features

-   Parse local or remote JavaScript files.
-   Accepts a single file, a list of files, or input from stdin.
-   Uses predefined regex patterns to find API endpoints and API keys.
-   Verbose mode for detailed output.
-   Outputs findings in JSON format for easy integration with other tools.
-   Pretty print option for human-readable console output.

## Documentation

For a detailed guide on how to use `jsReveal`, please see the [Usage Guide](docs/usage.md).

To understand the internal workings and architecture of the project, refer to the [How It Works](docs/how-it-works.md) document.

## Installation

### Using `go install`

```bash
go install github.com/SupremeERG/jsReveal@latest
```

### From source

```bash
git clone https://github.com/SupremeERG/jsReveal.git
cd jsReveal
go install
```

## Usage

`jsReveal` can be used in several ways depending on the source of the JavaScript files.

### Basic Flags

| Flag         | Description                                                 |
| :----------- | :---------------------------------------------------------- |
| `-f`         | Path to a target JS file.                                   |
| `-l`         | Path to a file with JS URLs.                                |
| `-u`         | URL to a singular JS file.                                  |
| `-v`         | Enable verbosity for more detailed output.                  |
| `--endpoint` | Use a predefined regex file for API endpoints and directories. |
| `--api-key`  | Use a predefined regex file for API keys.                   |
| `-o`         | Send output to a file in JSON format.                       |
| `--pretty`   | Pretty print the output to the console.                     |

### Examples

**1. Analyzing a single local JS file:**

```bash
jsReveal -f /path/to/your/file.js
```

**2. Analyzing a single remote JS file:**

```bash
jsReveal -u https://example.com/static/app.js
```

**3. Analyzing a list of remote JS files from a file:**

```bash
jsReveal -l /path/to/your/js_links.txt
```

*`js_links.txt` content:*
```
https://example.com/static/app.js
https://example.com/static/vendor.js
```

**4. Using stdin to provide URLs:**

```bash
cat /path/to/your/js_links.txt | jsReveal
```

**5. Using verbosity and saving output to a file:**

```bash
jsReveal -u https://example.com/static/app.js -v -o output.json
```

**6. Using pretty print for console output:**

```bash
jsReveal -u https://example.com/static/app.js --pretty
```

## Contributing

Contributions are welcome! Please read our [Contributing Guidelines](docs/contributing.md) to get started. Also, please be sure to review our [Code of Conduct](docs/CODE_OF_CONDUCT.md).

## License

`jsReveal` is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
