package parse

import (
	"fmt"
	"log"
	"os"

	"github.com/SupremeERG/jsReveal/pkg/fetchcode"
	"github.com/SupremeERG/jsReveal/pkg/regexmod"
	"github.com/dlclark/regexp2"
)

func ParseJS(jsFilePath string) {
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

	regexProperties := regexmod.RegexProperties{MatchLine: false, CaseInsensitive: true} // Example properties

	for _, pattern := range patterns {
		regexpPattern, err := regexmod.CompilePattern(pattern, regexProperties)
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
func ParseJSFromCode(jsCode string) {
	var matchTest *regexp2.Match
	var matches = []string{}

	// Grab the regex patterns from the file
	patterns, err := fetchcode.FetchPatterns()
	if err != nil {
		log.Fatal("Error reading regex patterns: ", err)
	}

	regexProperties := regexmod.RegexProperties{MatchLine: false, CaseInsensitive: true} // Example properties

	for _, pattern := range patterns {
		/* JSON RARSE
		regexJSON := fetchPatterns()
		var categories map[string]RegexProperties
		err = json.Unmarshal(regexJSON, &categories)
		if err != nil {
			fmt.Println("Error parsing regex JSON:", err)
			return
		}*/
		regexpPattern, err := regexmod.CompilePattern(pattern, regexProperties)
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
