package parser

import (
	"fmt"
	"log"
	"os"

	"github.com/SupremeERG/jsReveal/pkg/fetchcode"
	"github.com/SupremeERG/jsReveal/pkg/misc"
	"github.com/SupremeERG/jsReveal/pkg/regexmod"
	"github.com/dlclark/regexp2"
)

// the algorithm that scans the JS code
func parse(patterns []string, jsCode string, verbosity bool, source string, regexFile string) {
	var matchTest *regexp2.Match
	var matches = []string{}

	for _, pattern := range patterns {
		regexProperties := regexmod.DetermineProperties(pattern, regexFile)
		regexpPattern, err := regexmod.CompilePattern(pattern, regexProperties)
		if err != nil {
			log.Fatal("Error compiling regular expression '", pattern, "': ", err)
		}
		matchTest, _ = regexpPattern.FindStringMatch(jsCode)

		for matchTest != nil && misc.Contains(matches, matchTest.String()) == false {

			matches = append(matches, matchTest.String())
			match := matchTest.String()
			if len(match) > 1000 {
				match = match[:250] + "\n" // Prevents large blocks of code
			}
			if verbosity == true {
				fmt.Printf("Category: %s\nMatch: %s\nConfidence: %s\nSource: %s\n\n\n", regexProperties.Type, match, regexProperties.Confidence, source)
			} else {
				fmt.Printf("%s\t(%s)\n\n", match, source)
			}
			matchTest, _ = regexpPattern.FindNextMatch(matchTest)
		}
	}
}

// parseJS parses JavaScript code from a file
func ParseJS(jsFilePath string, verbosity bool) {

	jsCode, err := os.ReadFile(jsFilePath)
	if err != nil {
		fmt.Println("Error reading JS file:", err)
		return
	}

	// Grab the regex patterns from the file
	patterns, regexFile, err := fetchcode.FetchPatterns("regex.txt")
	if err != nil {
		log.Fatal("Error reading regex patterns: ", err)
	}

	parse(patterns, string(jsCode), verbosity, jsFilePath, regexFile)

	/*
		for _, pattern := range patterns {
			regexProperties := regexmod.DetermineProperties(pattern, "")
			regexpPattern, err := regexmod.CompilePattern(pattern, regexProperties)
			if err != nil {
				log.Fatal("Error compiling regular expression '", pattern, "': ", err)
			}
			matchTest, _ = regexpPattern.FindStringMatch(string(jsCode))

			for matchTest != nil && misc.Contains(matches, matchTest.String()) == false {
				matches = append(matches, matchTest.String())
				match := matchTest.String()
				if len(match) > 1000 {
					match = match[:250] + "\n" // Prevents large blocks of code
				}

				if verbosity == true {
					fmt.Printf("Category: %s\nMatch: %s\nConfidence: %s\n", regexProperties.Type, match, regexProperties.Confidence)
				} else {
					fmt.Print(match, "\n\n") // removing the pattern from print, it is only needed to test
				}
				matchTest, _ = regexpPattern.FindNextMatch(matchTest)
			}
		}*/
}

// ParseJSFromCode parses JavaScript code from a string and applies regex patterns to find matches.
func ParseJSFromCode(jsCode string, source string, verbosity bool) {
	// Fetch regex patterns from the file
	patterns, regexFile, err := fetchcode.FetchPatterns("regex.txt")
	if err != nil {
		log.Fatalf("Error reading regex patterns from regex.txt: %v", err)
	}

	// Apply regex patterns to the JavaScript code
	applyRegexPatterns(patterns, jsCode, verbosity, source, regexFile)
}

// applyRegexPatterns applies a set of regex patterns to a given string of JavaScript code.
func applyRegexPatterns(patterns []string, jsCode string, verbosity bool, source string, regexFile string) {
	for _, pattern := range patterns {
		regexProperties := regexmod.DetermineProperties(pattern, regexFile)
		// Assuming DetermineProperties only returns the properties without error

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
			fmt.Printf("Category: %s\nMatch: %s\nConfidence: %s\nSource: %s\n\n", regexProperties.Type, match, regexProperties.Confidence, source)
		} else {
			fmt.Printf("%s\t(%s)\n", match, source)
		}

		matchTest, err = regexpPattern.FindNextMatch(matchTest)
		if err != nil {
			log.Printf("Error finding next match: %v", err)
			break
		}
	}
}
