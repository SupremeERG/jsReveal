package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	// "strings"
	"github.com/SupremeERG/jsReveal/pkg/regex"

	"github.com/dlclark/regexp2"
)

type RegexProperties struct {
	MatchLine       bool   `json:"match_line"`
	CaseInsensitive bool   `json:"case_insensitive"`
	Type            string `json:"type"`
}

func compilePattern(pattern string, regexProperties RegexProperties) (*regexp2.Regexp, error) {
	var flags regexp2.RegexOptions
	validPattern := pattern
	if regexProperties.MatchLine {
		validPattern = fmt.Sprintf("%s.*(?:\n|$)", validPattern)
	}
	if regexProperties.CaseInsensitive {
		flags |= regexp2.IgnoreCase
	}

	return regexp2.Compile(validPattern, flags)
}

func fetchPatterns() ([]string, error) {
	file, err := os.Open("regex.txt") // Replace with your file path
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var patterns []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		patterns = append(patterns, scanner.Text())
	}

	return patterns, scanner.Err()
}

func parseJS(jsFilePath string) {
	var matchTest *regexp2.Match
	var matches = []string{}

	jsCode, err := os.ReadFile(jsFilePath)
	if err != nil {
		fmt.Println("Error reading JS file:", err)
		return
	}

	// Grab the regex patterns from the file
	patterns, err := fetchPatterns()
	if err != nil {
		log.Fatal("Error reading regex patterns: ", err)
	}

	regexProperties := RegexProperties{MatchLine: false, CaseInsensitive: true} // Example properties

	for _, pattern := range patterns {
		regexpPattern, err := compilePattern(pattern, regexProperties)
		if err != nil {
			log.Fatal("Error compiling regular expression '", pattern, "': ", err)
		}
		matchTest, _ = regexpPattern.FindStringMatch(string(jsCode))

		for matchTest != nil {
			matches = append(matches, matchTest.String())
			match := matchTest.String()
			if len(match) > 1000 {
				match = match[:250] + "\n" // Prevents large blocks of code
			}
			fmt.Printf("Pattern: %s\nMatch: %s\n\n", pattern, match)
			matchTest, _ = regexpPattern.FindNextMatch(matchTest)
		}
	}
}

// parseJSFromCode parses JavaScript code from a string.
func parseJSFromCode(jsCode string) {
	var matchTest *regexp2.Match
	var matches = []string{}

	// Grab the regex patterns from the file
	patterns, err := fetchPatterns()
	if err != nil {
		log.Fatal("Error reading regex patterns: ", err)
	}

	regexProperties := RegexProperties{MatchLine: false, CaseInsensitive: true} // Example properties

	for _, pattern := range patterns {
		regexpPattern, err := compilePattern(pattern, regexProperties)
		if err != nil {
			log.Fatal("Error compiling regular expression '", pattern, "': ", err)
		}
		matchTest, _ = regexpPattern.FindStringMatch(jsCode)

		for matchTest != nil {
			matches = append(matches, matchTest.String())
			match := matchTest.String()
			if len(match) > 1000 {
				match = match[:250] + "\n" // Prevents large blocks of code
			}
			fmt.Printf("Pattern: %s\nMatch: %s\n\n", pattern, match)
			matchTest, _ = regexpPattern.FindNextMatch(matchTest)
		}
	}
}

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

// Function to fetch JS code from a URL
func fetchJSFromURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	jsFilePath := flag.String("f", "", "Path to the .js file")
	jsLinksPath := flag.String("l", "", "Path to the file with JS links")
	jsURL := flag.String("u", "", "URL to a JS file")
	verbose := flag.Bool("v", false, "Enable verbose output")
	flag.Parse()

	if *verbose {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	} else {
		log.SetFlags(0)
	}

	log.SetPrefix("jsreveal: ")

	// Check if the JS links file path is provided
	if *jsLinksPath != "" {
		if *verbose {
			log.Println("Processing JS links from file:", *jsLinksPath)
		}

		links, err := readJSLinks(*jsLinksPath)
		if err != nil {
			log.Fatal("Error reading JS links: ", err)
		}

		for _, link := range links {
			jsCode, err := fetchJSFromURL(link)
			if err != nil {
				log.Printf("Error fetching JS from URL %s: %v\n", link, err)
				continue
			}
			parseJSFromCode(jsCode) // Assuming parseJSFromCode accepts a string of JS code
		}
		return
	} else if *jsURL != "" {
		jsCode, err := fetchJSFromURL(*jsURL)
		if err != nil {
			log.Printf("Error fetching JS from URL %s: %v\n", *jsURL, err)
		}
		parseJSFromCode(jsCode)
		return
	}

	if *jsFilePath == "" {
		fmt.Println("Usage: go run main.go -f <input_file.js> [-v for verbose]\n-f - file with js code\n-l - file with links to js code\n-v - for verbose")
		return
	}

	/*
		if *verbose {
			log.Println("Verbose mode enabled")
		}*/ // little unnecessary if they added the option

	parseJS(*jsFilePath)
}
