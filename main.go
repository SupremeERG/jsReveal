package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/SupremeERG/jsReveal/internal/parser"
	"github.com/SupremeERG/jsReveal/pkg/fetchcode"
	"github.com/SupremeERG/jsReveal/pkg/misc"
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

func run(options runner.Options, outputChannel chan string, cleanPattern *regexp.Regexp, signalChannel chan int) {
	switch options.Source {
	default:
		//fmt.Println("./jsReveal -u <url to JS file>")
		return
	case 1:
		parser.ParseJS(options.JSFilePath, options.Verbose, options.MatchLine, options.RegexFilePath, outputChannel)
		signalChannel <- 0

	case 2: // -l flag for a file containing multiple JS file paths
		parser.ParseJSFromList(options.JSLinksPath, options.Verbose, options.MatchLine, options.RegexFilePath, outputChannel)
		signalChannel <- 0

	case 3:
		if options.Verbose {
			log.Println("Processing Code from " + cleanPattern.ReplaceAllString(options.JSURL, ""))
		}
		ch := make(chan string)
		go fetchcode.FetchJSFromURL(options.JSURL, ch)
		jsCode := <-ch
		parser.ParseJSFromCode(jsCode, options.JSURL, options.Verbose, options.MatchLine, options.RegexFilePath, outputChannel)
		signalChannel <- 0

	case 4:
		stdinScanner := bufio.NewScanner(os.Stdin)
		for {
			stdinScanner.Scan()
			line := stdinScanner.Text()
			if len(line) == 0 {
				break
			}
			ch := make(chan string)
			go fetchcode.FetchJSFromURL(line, ch)
			jsCode := <-ch
			parser.ParseJSFromCode(jsCode, line, options.Verbose, options.MatchLine, options.RegexFilePath, outputChannel)

		}
		err := stdinScanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		signalChannel <- 0


	}

	return
}

func main() {

	options := runner.ParseOptions()
	outputChannel := make(chan string)
	cleanPattern, _ := regexp.Compile(".*/")
	signalChannel := make(chan int)


	// Checks if arguments were passed
	if options.Source == 0 {
		fmt.Println("./jsReveal -u <url to JS file>")
		return
	}
	go run(options, outputChannel, cleanPattern, signalChannel)


	go func() {
		for signal := range signalChannel {
			if signal == 0 {// this means end the parsing and stop the script
				close(outputChannel)
			}
		}
	}()
	// Output Component
	if options.FileOutput != "" {

		jsonData := make(map[string]interface{})
		for output := range outputChannel {
			var parts []string
			var lineData map[string]interface{}

			if options.Verbose == true {

				parts = strings.Split(output, "::::")
				lineData = map[string]interface{}{
					"type":       parts[0],
					"confidence": parts[2],
					"source":     parts[3],
				}
			} else {
				parts = strings.Split(output, "::::")
				lineData = map[string]interface{}{
					"source": parts[1],
				}
			}

			existingData, err := misc.ReadExistingJSON(options.FileOutput)
			if err != nil {
				// Add the line data to the overall JSON data map
				jsonData[fmt.Sprintf("\"%s\"", parts[1])] = lineData

				go misc.WriteJSONToFile(options.FileOutput, jsonData)
				go fmt.Println(output)

			} else {
				// Add the line data to the overall JSON data map
				existingData[fmt.Sprintf("\"%s\"", parts[1])] = lineData
				go misc.WriteJSONToFile(options.FileOutput, existingData)
				go fmt.Println(output)
			}

		}
	} else {

		if options.PrettyPrint == true {
			for output := range outputChannel {
				newOut := strings.Replace(output, "::::", "\n", -1)
				fmt.Println(newOut + "\n")
			}
		} else {
			for output := range outputChannel {
				fmt.Println(output)
			}
		}
	}

}
