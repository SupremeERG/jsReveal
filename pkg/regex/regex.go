package regexmod

import (
	"fmt"

	"github.com/dlclark/regexp2"
)

const Version = "v1.0.0"

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
