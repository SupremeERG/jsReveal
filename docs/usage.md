# How to Use jsReveal

`jsReveal` is a versatile tool for analyzing JavaScript files from various sources. This guide provides a detailed walkthrough of its features and command-line options.

## Command-Line Flags

Here is a complete list of the available flags:

| Flag         | Description                                                 |
| :----------- | :---------------------------------------------------------- |
| `-f`         | Path to a target JS file on your local machine.             |
| `-l`         | Path to a text file containing a list of JS file URLs.      |
| `-u`         | URL to a single remote JS file.                             |
| `-v`         | Enable verbosity for more detailed output.                  |
| `--endpoint` | Use a predefined regex file for API endpoints and directories. |
| `--api-key`  | Use a predefined regex file for API keys.                   |
| `-o`         | Send output to a file in JSON format.                       |
| `--pretty`   | Pretty print the output to the console for better readability.|

## Usage Scenarios

### 1. Analyzing a Local JavaScript File

If you have a JavaScript file saved on your computer, you can analyze it using the `-f` flag.

```bash
jsReveal -f /path/to/your/file.js
```

### 2. Analyzing a Remote JavaScript File

To analyze a JavaScript file hosted on a web server, use the `-u` flag followed by the URL.

```bash
jsReveal -u https://example.com/static/app.js
```

### 3. Analyzing Multiple Remote Files

If you have a list of URLs to JavaScript files, you can save them in a text file (one URL per line) and use the `-l` flag.

**Example `js_links.txt`:**
```
https://example.com/static/app.js
https://example.com/static/vendor.js
https://another-site.com/main.js
```

**Command:**
```bash
jsReveal -l /path/to/your/js_links.txt
```

### 4. Piping URLs from stdin

`jsReveal` also supports reading URLs from standard input. This is useful for chaining commands together.

```bash
cat /path/to/your/js_links.txt | jsReveal
```

You can also use it with other tools like `gau` or `hakrawler`:

```bash
gau example.com | grep '\.js$' | jsReveal
```

### 5. Customizing Output

#### JSON Output

For programmatic use or integration with other tools, you can save the output in JSON format using the `-o` flag.

```bash
jsReveal -u https://example.com/static/app.js -o findings.json
```

#### Verbose Output

To get more detailed information about the findings, use the `-v` flag. This will include the type of finding and a confidence score.

```bash
jsReveal -u https://example.com/static/app.js -v
```

#### Pretty Print

For a more human-readable output in the console, use the `--pretty` flag.

```bash
jsReveal -u https://example.com/static/app.js --pretty
```

### 6. Using Custom Regex Patterns

While `jsReveal` comes with predefined regex patterns for API keys and endpoints, you can supply your own files with custom patterns.

**For endpoints:**
```bash
jsReveal -u https://example.com/static/app.js --endpoint /path/to/my_endpoint_regex.txt
```

**For API keys:**
```bash
jsReveal -u https://example.com/static/app.js --api-key /path/to/my_apikey_regex.txt
```
