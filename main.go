package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	// "strings"

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

func parseJS() {
	var matchTest *regexp2.Match
	var matches = []string{}

	jsCode, err := os.ReadFile(os.Args[1])
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

func main() {
	log.SetFlags(0)
	log.SetPrefix("jsreveal: ")
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file.js>")
		return
	}

	parseJS()
}
