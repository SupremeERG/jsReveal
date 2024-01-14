package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/SupremeERG/jsReveal/pkg/fetchcode"
	"github.com/SupremeERG/jsReveal/pkg/misc"
	"github.com/SupremeERG/jsReveal/pkg/regexmod"
	"github.com/SupremeERG/opshins"
	"github.com/dlclark/regexp2"
)

// parse applies regex patterns to a given string of JavaScript code.
func parse(patterns []string, jsCode string, verbosity bool, source string, regexFile string, outputChan chan<- string) {
	var matchTest *regexp2.Match
	var matches = []string{}

	for _, pattern := range patterns {
		regexProperties := regexmod.DetermineProperties(pattern, regexFile)
		regexpPattern, err := regexmod.CompilePattern(pattern, regexProperties)
		if err != nil {
			log.Fatalf(`{"Error compiling regular expression", "Pattern": "%s", "Error": "%v"}`, pattern, err)
		}
		matchTest, _ = regexpPattern.FindStringMatch(jsCode)

		for matchTest != nil && !misc.Contains(matches, matchTest.String()) {
			matches = append(matches, matchTest.String())
			match := matchTest.String()
			if verbosity == true {
				outputChan <- fmt.Sprintf("%s\n%s\n%s\n%s\n\n\n", regexProperties.Type, match, regexProperties.Confidence, source)

			} else {
				outputChan <- fmt.Sprintf("%s\n%s\n\n", match, source)
			}
			matchTest, _ = regexpPattern.FindNextMatch(matchTest)

		}
	}

}

// ParseJS parses JavaScript code from a file using specified regex patterns.
func ParseJS(jsFilePath string, verbosity bool, regexFilePath string, outputChan chan string) {
	jsCode, err := os.ReadFile(jsFilePath)
	if err != nil {
		log.Fatalf(`{"Error reading JS file": "%v"}`, err)
	}

	patterns, regexFile, err := fetchcode.FetchPatterns(regexFilePath) // Adjust to capture all returned values
	if err != nil {
		log.Fatalf(`{"Error reading regex patterns": "%s", "Error": "%v"}`, regexFilePath, err)
	}

	parse(patterns, string(jsCode), verbosity, jsFilePath, regexFile, outputChan)
}

// FetchJSFromURL fetches JavaScript code from a URL
func FetchJSFromURL(jsURL string) (string, error) {
	resp, err := http.Get(jsURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// section for -l
const MaxConcurrentJobs = 10 // Adjust the concurrency level as needed

func ParseJSFromList(listFilePath string, verbosity bool, regexFilePath string, outputChan chan string) {
	if strings.HasSuffix(listFilePath, ".js") == true {
		if opshins.PromptYN(fmt.Sprintf("%s looks like a javascript file, not a wordlist of JS URLs. Are you sure you want to use this file?", listFilePath), "yes", " > ") == "no" {
			close(outputChan)
			return
		}
	}

	listContent, err := os.ReadFile(listFilePath)
	if err != nil {
		log.Fatalf(`{"Error reading list file": "%v"}`, err)
	}

	jsURLs := strings.Split(string(listContent), "\n")

	// Create a channel for job distribution
	jobs := make(chan string, len(jsURLs))

	// Start a fixed number of worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < MaxConcurrentJobs; i++ {
		wg.Add(1)
		go urlWorker(&wg, jobs, verbosity, regexFilePath, outputChan)
	}

	// Distribute work

	for _, url := range jsURLs {
		jobs <- url
	}
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()
}

func urlWorker(wg *sync.WaitGroup, jobs <-chan string, verbosity bool, regexFilePath string, outputChan chan string) {
	defer wg.Done()
	for jsURL := range jobs {
		if jsURL == "" {
			continue
		}

		jsURL = strings.TrimSpace(jsURL)
		if _, err := url.ParseRequestURI(jsURL); err != nil {
			log.Printf(`{"Invalid URL": "%s", "Error": "%v"}`, jsURL, err)
			continue
		}

		jsCode, err := FetchJSFromURL(jsURL)
		if err != nil {
			log.Printf(`{"Failed to fetch JS from URL": "%s", "Error": "%v"}`, jsURL, err)
			continue
		}

		ParseJSFromCode(jsCode, jsURL, verbosity, regexFilePath, outputChan)
	}
}

/// end section for -l

// ParseJSFromCode parses JavaScript code from a string using specified regex patterns.
func ParseJSFromCode(jsCode string, source string, verbosity bool, regexFilePath string, outputChan chan string) {
	patterns, regexFile, err := fetchcode.FetchPatterns(regexFilePath) // Adjust to capture all returned values
	if err != nil {
		log.Fatalf("Error reading regex patterns from %s: %v", regexFilePath, err)
	}

	parse(patterns, jsCode, verbosity, source, regexFile, outputChan)
	//applyRegexPatterns(patterns, jsCode, verbosity, source, regexFile)
}
