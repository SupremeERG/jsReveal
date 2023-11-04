package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)


type RegexProperties struct {
	MatchLine       bool   `json:"match_line"`
	CaseInsensitive bool   `json:"case_insensitive"`
	Type            string `json:"type"`
}

func compilePattern(pattern string, regexProperties RegexProperties) (*regexp.Regexp, error) {
	//flags := 0
	validPattern := pattern
	if regexProperties.MatchLine {
		validPattern = fmt.Sprintf("%s.*(?:\n|$)", validPattern)
		// flags |= regexp.DOTALL
	}
	if regexProperties.CaseInsensitive {
		validPattern = fmt.Sprintf("(?i)%s", validPattern)
	}

	return regexp.Compile(validPattern)
}

func parseJS() {
	// Grab the JS file
	jsCode, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading JS file:", err)
		return
	}

	// Grab the regex patterns from JSON
	regexJSON, err := ioutil.ReadFile("regex.json")
	if err != nil {
		fmt.Println("Error reading regex file:", err)
		return
	}

	var categories map[string]RegexProperties
	err = json.Unmarshal(regexJSON, &categories)
	if err != nil {
		fmt.Println("Error parsing regex JSON:", err)
		return
	}

	for pattern, regexProperties := range categories {
		regexpPattern, err := compilePattern(pattern, regexProperties)
		if err != nil {
			fmt.Println("Error compiling regular expression:", err)
			return
		}

		matches := regexpPattern.FindAllString(string(jsCode), -1)
		if matches != nil {
			for _, match := range matches {
				if len(match) > 1000 {
					match = match[:250] // Prevents humungous blocks of minified code from being outputted
				}
				fmt.Printf("Category: %s\nString: %s\n", regexProperties.Type, match)
			}
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file.js>")
		return
	}

	parseJS()
}
