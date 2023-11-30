package parser

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
			log.Fatalf(`{"Error compiling regular expression": "%s", "Pattern": "%s", "Error": "%v"}`, pattern, err)
		}
		matchTest, _ = regexpPattern.FindStringMatch(jsCode)

		for matchTest != nil && !misc.Contains(matches, matchTest.String()) {
			matches = append(matches, matchTest.String())
			match := matchTest.String()
			outputChan <- fmt.Sprintf(`{"Category": "%s", "Match": "%s", "Confidence": "%s", "Source": "%s"}`, regexProperties.Type, match, regexProperties.Confidence, source)
			matchTest, _ = regexpPattern.FindNextMatch(matchTest)
			close(outputChan)
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
		go worker(&wg, jobs, verbosity, regexFilePath, outputChan)
	}

	// Distribute work
	for _, url := range jsURLs {
		jobs <- url
	}
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()
}

func worker(wg *sync.WaitGroup, jobs <-chan string, verbosity bool, regexFilePath string, outputChan chan string) {
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

/*

// applyRegexPatterns applies a set of regex patterns to a given string of JavaScript code.
func applyRegexPatterns(patterns []string, jsCode string, verbosity bool, source string, regexFile string) {
	for _, pattern := range patterns {
		regexProperties := regexmod.DetermineProperties(pattern, regexFile)
		regexpPattern, err := regexmod.CompilePattern(pattern, regexProperties)
		if err != nil {
			log.Fatalf("Error compiling regular expression '%s': %v", pattern, err)
		}

		findAndPrintMatches(regexpPattern, jsCode, regexProperties, verbosity, source)
	}
}

// findAndPrintMatches finds and prints matches of a compiled regex pattern in a given JavaScript code.
func findAndPrintMatches(regexpPattern *regexp2.Regexp, jsCode string, regexProperties regexmod.RegexProperties, verbosity bool, source string) {
	var matches []string
	matchTest, err := regexpPattern.FindStringMatch(jsCode)
	if err != nil {
		log.Printf("Error finding matches for the pattern: %v", err)
		return
	}

	for matchTest != nil {
		match := matchTest.String()

		// Check for duplicates
		if misc.Contains(matches, match) {
			continue
		}

		matches = append(matches, match)

		// Truncate long matches
		if len(match) > 1000 {
			match = match[:250] + "..."
		}

		// Print match based on verbosity
		if verbosity {
			outputChan <- fmt.Sprintf("Category: %s\nMatch: %s\nConfidence: %s\nSource: %s\n\n", regexProperties.Type, match, regexProperties.Confidence, source)
		} else {
			outputChan <- fmt.Sprintf("%s\t(%s)\n", match, source)
		}

		matchTest, err = regexpPattern.FindNextMatch(matchTest)
		if err != nil {
			log.Printf("Error finding next match: %v", err)
			break
		}
	}
}
*/
