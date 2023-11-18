package regexmod

import (
	"fmt"

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
		validPattern = fmt.Sprintf("%s.*(?:\n|$)", validPattern)
	}
	if regexProperties.CaseInsensitive {
		flags |= regexp2.IgnoreCase
	}

	return regexp2.Compile(validPattern, flags)
}

func DetermineProperties(pattern string, file string) RegexProperties {
	// file is the file of regular expressions
	// pattern is the actual regular expression
	var category string
	switch file {
	default:
		category = "Unindentified"
	case "api_key_regex.txt":
		category = "Endpoint (API)"
	}
	properties := RegexProperties{MatchLine: false, CaseInsensitive: false, Confidence: "high", Type: category}

	// add code for certain detections (endpoints = high confidence, credentials = medium confidence, )
	return properties
}
