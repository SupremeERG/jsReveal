package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/SupremeERG/golang-pkg-pcre/src/pkg/pcre"
)

type RegexProperties struct {
	MatchLine       bool   `json:"match_line"`
	CaseInsensitive bool   `json:"case_insensitive"`
	Type            string `json:"type"`
}

func compilePattern(pattern string, regexProperties RegexProperties) (pcre.Regexp, *pcre.CompileError) {
	var flags int
	validPattern := pattern
	if regexProperties.MatchLine {
		validPattern = fmt.Sprintf("%s.*(?:\n|$)", validPattern)
	}
	if regexProperties.CaseInsensitive {
		//validPattern = fmt.Sprintf("%s", validPattern)
		flags |= pcre.CASELESS
	}

	return pcre.Compile(validPattern, flags)
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
	jsCode, err := fs.ReadFile(os.DirFS("."), os.Args[1])
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
			log.Fatal(fmt.Sprintf("Error compiling regular expression '%s': ", pattern), strings.ReplaceAll(err.Message, "\n", "\\n"))
			return
		}

		matches := regexpPattern.MatcherString(string(jsCode), 0)
		fmt.Println(matches.GroupString(0)) // im having issues with this regex package, the defualt one has limited options so i am trying a custom pkg with no documentation that uses perl syntax
		/*
			if matches != nil {
				for _, match := range matches {
					if len(match) > 1000 {
						match = match[:250] + "\n" // Prevents humungous blocks of minified code from being outputted
					}
					fmt.Printf("Category: %s\nString: %s\n", regexProperties.Type, match)
				}
			}*/
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
