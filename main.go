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
	options := runner.Run()

	cleanPattern, _ := regexp.Compile(".*/") // Used to make links more readable
	switch options.Source {
	default: // no option selected
		fmt.Println("./jsReveal -u <url to JS file>")
	case 1: // single file
		parser.ParseJS(options.JSFilePath, options.Verbose)

	case 2: // multiple URLs
		if options.Verbose {
			log.Println("Processing JS links from file:", cleanPattern.ReplaceAllString(options.JSLinksPath, ""))
		}

		links, err := readJSLinks(options.JSLinksPath)
		if err != nil {
			log.Fatal("Error reading JS links: ", err)
		}

		ch := make(chan string)
		for _, link := range links {
			go fetchcode.FetchJSFromURL(link, ch)

			go parser.ParseJSFromCode(<-ch, link, options.Verbose) // Assuming parseJSFromCode accepts a string of JS code

		}

	case 3: // single url
		if options.Verbose {
			log.Println("Processing Code from " + cleanPattern.ReplaceAllString(options.JSURL, ""))
		}
		ch := make(chan string)
		fetchcode.FetchJSFromURL(options.JSURL, ch)

		parser.ParseJSFromCode(<-ch, options.JSURL, options.Verbose)
	}
}
