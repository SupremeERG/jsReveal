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

// parseJS parses JavaScript code from a file
func ParseJS(jsFilePath string, verbosity bool) {

	var matchTest *regexp2.Match
	var matches = []string{}

	jsCode, err := os.ReadFile(jsFilePath)
	if err != nil {
		fmt.Println("Error reading JS file:", err)
		return
	}

	// Grab the regex patterns from the file
	patterns, err := fetchcode.FetchPatterns()
	if err != nil {
		log.Fatal("Error reading regex patterns: ", err)
	}

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
	}
}

// parseJSFromCode parses JavaScript code from a string.
func ParseJSFromCode(jsCode string, source string, verbosity bool) {
	var matchTest *regexp2.Match
	var matches = []string{}

	// Grab the regex patterns from the file
	patterns, err := fetchcode.FetchPatterns()
	if err != nil {
		log.Fatal("Error reading regex patterns: ", err)
	}

	/* // to fetch from JSON file
	categories, _ := fetchcode.FetchPatternsFromJSON()
	for pattern, regexProperties := range categories {*/
	for _, pattern := range patterns {
		regexProperties := regexmod.DetermineProperties(pattern, "")
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
