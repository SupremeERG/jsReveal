package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/SupremeERG/jsReveal/internal/parser"
	"github.com/SupremeERG/jsReveal/pkg/fetchcode"
	"github.com/SupremeERG/jsReveal/runner"
)

// function to read JS links from a file
func readJSLinks(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var links []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		links = append(links, scanner.Text())
	}

	return links, scanner.Err()
}

func main() {
	options := runner.ParseOptions()

	cleanPattern, _ := regexp.Compile(".*/")
	switch options.Source {
	default:
		fmt.Println("./jsReveal -u <url to JS file>")
	case 1:
		parser.ParseJS(options.JSFilePath, options.Verbose, options.RegexFilePath)
	case 2: // -l flag for a file containing multiple JS file paths
		parser.ParseJSFromList(options.JSLinksPath, options.Verbose, options.RegexFilePath)
	case 3:
		if options.Verbose {
			log.Println("Processing Code from " + cleanPattern.ReplaceAllString(options.JSURL, ""))
		}
		ch := make(chan string)
		go fetchcode.FetchJSFromURL(options.JSURL, ch)
		jsCode := <-ch
		parser.ParseJSFromCode(jsCode, options.JSURL, options.Verbose, options.RegexFilePath)
	}
}
