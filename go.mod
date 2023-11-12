module github.com/SupremeERG/jsReveal

go 1.21.3

replace github.com/SupremeERG/jsReveal/pkg/regexmod => ./pkg/regexmod

replace github.com/SupremeERG/jsReveal/pkg/fetchcode => ./pkg/fetchcode

replace github.com/SupremeERG/jsReveal/internal/parser => ./internal/parser

replace github.com/SUpremeERG/jsReveal/runner => ./runner

require github.com/SupremeERG/jsReveal/internal/parser v0.0.0-00010101000000-000000000000

require (
	github.com/SupremeERG/jsReveal/pkg/fetchcode v0.0.0-00010101000000-000000000000 // indirect
	github.com/SupremeERG/jsReveal/pkg/regexmod v0.0.0-00010101000000-000000000000 // indirect
	github.com/dlclark/regexp2 v1.10.0 // indirect
)
