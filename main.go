package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"

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
		//validPattern = fmt.Sprintf("%s", validPattern)
		flags |= regexp2.IgnoreCase
	}

	return regexp2.Compile(validPattern, flags)
}

func fetchPatterns() []byte {
	regexJSON, err := fs.ReadFile(os.DirFS("."), "regex.json")
	if err != nil {
		log.Fatal("Error reading regex file: ", err)
		//fmt.Println("Error reading regex file:", err)
	}

	return regexJSON
}

func parseJS() {
	// Grab the JS file
	var matchTest *regexp2.Match
	var matches = []string{}

	jsCode, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading JS file:", err)
		return
	}

	// Grab the regex patterns from JSON
	regexJSON := fetchPatterns()

	var categories map[string]RegexProperties
	err = json.Unmarshal(regexJSON, &categories)
	if err != nil {
		fmt.Println("Error parsing regex JSON:", err)
		return
	}

	// start testing the file(s) for appealing code
	for pattern, regexProperties := range categories {
		regexpPattern, err := compilePattern(pattern, regexProperties)
		if err != nil {
			log.Fatal(fmt.Sprintf("Error compiling regular expression '%s': ", pattern), strings.ReplaceAll(err.Error(), "\n", "\\n"))
		}
		matchTest, _ = regexpPattern.FindStringMatch(string(jsCode))

		for matchTest != nil {
			matches = append(matches, matchTest.String())
			match := matchTest.String()
			if len(match) > 1000 {
				match = match[:250] + "\n" // Prevents humungous blocks of minified code from being outputted
			}
			fmt.Printf("Category: %s\nString: %s\n\n", regexProperties.Type, match) ///
			matchTest, err = regexpPattern.FindNextMatch(matchTest)
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
