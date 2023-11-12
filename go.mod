module github.com/SupremeERG/jsReveal

go 1.21.3

require (
	github.com/SupremeERG/jsReveal/pkg/fetchcode v0.0.0-00010101000000-000000000000
	github.com/SupremeERG/jsReveal/pkg/parse v0.0.0-00010101000000-000000000000
	github.com/SupremeERG/jsReveal/pkg/regexmod v0.0.0-00010101000000-000000000000
	github.com/dlclark/regexp2 v1.10.0
)

replace github.com/SupremeERG/jsReveal/pkg/regexmod => ./pkg/regexmod

replace github.com/SupremeERG/jsReveal/pkg/fetchcode => ./pkg/fetchcode

replace github.com/SupremeERG/jsReveal/pkg/parse => ./pkg/parse
