package runner

import (
	"flag"
	"fmt"
)

var helpMsg = `
Usage
-f  -- Path to target JS file
-l  -- Path to a file with JS URLs
-u  -- URL to a singular JS file
-v  -- Enable Verbosity

`

type Options struct {
	JSFilePath  string
	JSLinksPath string
	JSURL       string
	Source      int // the source that the script will use (could be js file, links file, singular link)
	Verbose     bool
}

func ParseOptions() Options {
	var options Options

	flag.StringVar(&options.JSFilePath, "f", "", "Path to the .js file")
	flag.StringVar(&options.JSLinksPath, "l", "", "Path to the file with JS links")
	flag.StringVar(&options.JSURL, "u", "", "URL to a JS file")
	flag.BoolVar(&options.Verbose, "v", false, "Enable verbose output")

	flag.Usage = func() {
		fmt.Print(helpMsg)
	}

	flag.Parse()

	switch {
	default:
		options.Source = 0
	case len(options.JSFilePath) > 0:
		options.Source = 1 // for a JS file
	case len(options.JSLinksPath) > 0:
		options.Source = 2 // for multple urls
	case len(options.JSURL) > 0:
		options.Source = 3 // for a single url
	}

	return options

}

/*
if *verbose {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	} else {
		log.SetFlags(0)
	}*/

func PrintHelp() {
	fmt.Println(helpMsg)
}
