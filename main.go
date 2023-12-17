package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

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

func run(options runner.Options, outputChannel chan string, cleanPattern *regexp.Regexp) {
	switch options.Source {
	default:
		//fmt.Println("./jsReveal -u <url to JS file>")
		return
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

	return
}

func main() {
	options := runner.ParseOptions()
	outputChannel := make(chan string)
	cleanPattern, _ := regexp.Compile(".*/")

	// Checks if arguments were passed
	if options.Source == 0 {
		fmt.Println("./jsReveal -u <url to JS file>")
		return
	} else {
		go run(options, outputChannel, cleanPattern)

		// Output Component
		if options.FileOutput != "" {
			f, err := os.Create(options.FileOutput)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			w := bufio.NewWriter(f)

			jsonData := make(map[string]interface{})
			for output := range outputChannel {
				parts := strings.Fields(output)

				/*
					// Create a map to hold the JSON structure
					jsonData := map[string]interface{}{
						"type":       parts[0],
						"confidence": parts[2],
						"source":     parts[3],
					}

					// Add the line data to the overall JSON data map

					// Convert the map to JSON
					jsonBytes, err := json.MarshalIndent(map[string]interface{}{fmt.Sprintf("\"%s\"", parts[1]): jsonData}, "", "    ")
					if err != nil {
						fmt.Println("Error marshaling JSON:", err)
						return
					}*/
				lineData := map[string]interface{}{
					"type":       parts[0],
					"confidence": parts[2],
					"source":     parts[3],
				}

				// Add the line data to the overall JSON data map
				jsonData[fmt.Sprintf("\"%s\"", parts[1])] = lineData

				jsonBytes, err := json.MarshalIndent(jsonData, "", "    ")
				if err != nil {
					fmt.Println("Error marshaling JSON:", err)
					return
				}

				// Write the JSON to the file
				fmt.Fprintln(w, string(jsonBytes))
				//fmt.Fprintln(w, output)
				w.Flush()
			}
		} else {
			for output := range outputChannel {
				fmt.Println(output)
			}
		}
	}
}
