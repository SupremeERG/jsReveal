package runner

import (
	"flag"
	"fmt"
	"path/filepath"
	"runtime"
)

var helpMsg = `
Usage
-f  -- Path to target JS file
-l  -- Path to a file with JS URLs
-u  -- URL to a singular JS file
-v  -- Enable Verbosity
--api-endpoint  -- Use predefined regex file for API endpoints
--api-key       -- Use predefined regex file for API keys
`

type Options struct {
	JSFilePath     string
	JSLinksPath    string
	JSURL          string
	Source         int
	Verbose        bool
	UseAPIEndpoint bool
	UseAPIKey      bool
	RegexFilePath  string
}

func ParseOptions() Options {
	var options Options

	// Grab the current working directory, move up one level to grab the regex files.
	_, filename, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(filepath.Dir(filename)) // This goes up one directory

	flag.StringVar(&options.JSFilePath, "f", "", "Path to the .js file")
	flag.StringVar(&options.JSLinksPath, "l", "", "Path to the file with JS links")
	flag.StringVar(&options.JSURL, "u", "", "URL to a JS file")
	flag.BoolVar(&options.Verbose, "v", false, "Enable verbose output")
	flag.BoolVar(&options.UseAPIEndpoint, "api-endpoint", false, "Use predefined regex file for API endpoints")
	flag.BoolVar(&options.UseAPIKey, "api-key", false, "Use predefined regex file for API keys")

	flag.Usage = func() {
		fmt.Print(helpMsg)
	}

	flag.Parse()

	if options.UseAPIKey {
		options.RegexFilePath = filepath.Join(basepath, "api_key_regex.txt")
	} else { // Default to API endpoint regex file
		options.RegexFilePath = filepath.Join(basepath, "regex.txt")
	}

	switch {
	default:
		options.Source = 0
	case len(options.JSFilePath) > 0:
		options.Source = 1
	case len(options.JSLinksPath) > 0:
		options.Source = 2
	case len(options.JSURL) > 0:
		options.Source = 3
	}

	return options
}

func PrintHelp() {
	fmt.Println(helpMsg)
}
