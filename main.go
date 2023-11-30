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
	outputChannel := make(chan string)
	cleanPattern, _ := regexp.Compile(".*/")
	go func() {
		switch options.Source {
		default:
			fmt.Println("./jsReveal -u <url to JS file>")
		case 1:
			parser.ParseJS(options.JSFilePath, options.Verbose, options.RegexFilePath, outputChannel)
		case 2: // -l flag for a file containing multiple JS file paths
			parser.ParseJSFromList(options.JSLinksPath, options.Verbose, options.RegexFilePath, outputChannel)
		case 3:
			if options.Verbose {
				log.Println("Processing Code from " + cleanPattern.ReplaceAllString(options.JSURL, ""))
			}
			ch := make(chan string)
			go fetchcode.FetchJSFromURL(options.JSURL, ch)
			jsCode := <-ch
			parser.ParseJSFromCode(jsCode, options.JSURL, options.Verbose, options.RegexFilePath, outputChannel)
		}
	}()

	// output functionality
	if options.FileOutput != "" {
		f, err := os.Create(options.FileOutput)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		w := bufio.NewWriter(f)
		for output := range outputChannel {
			fmt.Fprintln(w, output)
			w.Flush()
		}
	} else {
		for output := range outputChannel {
			fmt.Println(output)
		}
	}
}
