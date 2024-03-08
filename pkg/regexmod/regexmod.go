package regexmod

import (
	"fmt"
	"strings"

	"github.com/dlclark/regexp2"
)

type RegexProperties struct {
	MatchLine       bool   `json:"match_line"`
	CaseInsensitive bool   `json:"case_insensitive"`
	Type            string `json:"type"`
	Confidence      string `json:"confidence"`
}

func CompilePattern(pattern string, regexProperties RegexProperties) (*regexp2.Regexp, error) {
	var flags regexp2.RegexOptions
	validPattern := pattern
	if regexProperties.MatchLine {
		validPattern = fmt.Sprintf("%s.*(?!$)", validPattern)
	}
	if regexProperties.CaseInsensitive {
		flags |= regexp2.IgnoreCase
	}

	return regexp2.Compile(validPattern, flags)
}

func DetermineProperties(pattern string, file string, matchLine bool) RegexProperties {
	// file is the file of regular expressions
	// pattern is the actual regular expression
	split := strings.Split(file, "/")
	file = split[len(split)-1]

	var category string
	switch file {
	default:
		category = "Unindentified"
	case "endpoints.txt":
		category = "Endpoint (Files, Directories, API Endpoint)"
	}
	properties := RegexProperties{MatchLine: matchLine, CaseInsensitive: false, Confidence: "high", Type: category}

	// add code for certain detections (endpoints = high confidence, credentials = medium confidence, )
	return properties
}
