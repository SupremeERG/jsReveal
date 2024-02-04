package runner

import (
	"flag"
	"fmt"
	"path/filepath"
	"runtime"
)

var helpMsg = `
Usage
-f  		-- Path to target JS file
-l  		-- Path to a file with JS URLs
-u  		-- URL to a singular JS file
-v  		-- Enable Verbosity
--endpoint  -- Use predefined regex file for API endpoints and directories
--api-key  	-- Use predefined regex file for API keys
-o			-- Send output to file (JSON)
-p --pretty -- Print output in a prettier format instead of RESULT::::SOURCE
`

type Options struct {
	JSFilePath      string
	JSLinksPath     string
	JSURL           string
	Source          int
	SearchEndpoints bool
	SearchAPIKey    bool
	RegexFilePath   string
	Verbose         bool
	FileOutput      string
}

func ParseOptions() Options {
	var options Options

	// Grab the current working directory, move up one level to grab the regex files.
	_, filename, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(filepath.Dir(filename)) // This goes up one directory

	flag.StringVar(&options.JSFilePath, "f", "", "Path to the .js file")
	flag.StringVar(&options.JSLinksPath, "l", "", "Path to the file with JS links")
	flag.StringVar(&options.JSURL, "u", "", "URL to a JS file")
	flag.BoolVar(&options.SearchEndpoints, "endpoint", false, "Use predefined regex file for API endpoints")
	flag.BoolVar(&options.SearchAPIKey, "api-key", false, "Use predefined regex file for API keys")
	flag.BoolVar(&options.Verbose, "v", false, "Enable verbose output")
	flag.StringVar(&options.FileOutput, "o", "", "Send output to file (JSON)")

	flag.Usage = func() {
		fmt.Print(helpMsg)
	}

	flag.Parse()

	/*
		parseCategories := []bool{options.SearchEndpoints, options.SearchAPIKey}
		for i := 0; i < len(parseCategories); i++ {
			/x*
				option 0 = options.SearchEndpoints
				option 1 = options.SearchAPIKey
			*x/
			if parseCategories[i] == true {
				if i == 0 {
					// options.RegexFilePath = filepath.Join(basepath, "api_endpoints_regex.txt")

				} else if i == 1 {
					options.RegexFilePath = filepath.Join(basepath, "api_key_regex.txt")

				}
			}
			if i == (len(parseCategories) - 1)
		}*/
	/*
		if options.SearchAPIKey {
			options.RegexFilePath = filepath.Join(basepath, "api_key_regex.txt")
		} else { // Default to API endpoint regex file
			options.RegexFilePath = filepath.Join(basepath, "regex.txt")
		}*/

	switch {
	default:
		fmt.Println("Scanning for endpoints (default option)")
		options.RegexFilePath = filepath.Join(basepath, "endpoints.txt")
	case options.SearchEndpoints:
		//		log.Fatal("endpoints regular expressions not implemented yet, sorry")
		options.RegexFilePath = filepath.Join(basepath, "endpoints.txt")
	case options.SearchAPIKey:
		options.RegexFilePath = filepath.Join(basepath, "api_key_regex.txt")

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
