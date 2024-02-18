module github.com/SupremeERG/jsReveal

go 1.21.3

toolchain go1.21.4

//replace github.com/SupremeERG/jsReveal/pkg/regexmod => ./pkg/regexmod

//replace github.com/SupremeERG/jsReveal/pkg/fetchcode => ./pkg/fetchcode

//replace github.com/SupremeERG/jsReveal/internal/parser => ./internal/parser

//replace github.com/SupremeERG/jsReveal/pkg/misc => ./pkg/misc

//replace github.com/SupremeERG/jsReveal/runner => ./runner

require (
	github.com/SupremeERG/jsReveal/pkg/fetchcode v0.0.0-20240114003256-88945ec39555
	github.com/SupremeERG/jsReveal/pkg/misc v0.0.0-20240114003256-88945ec39555
	github.com/SupremeERG/jsReveal/pkg/regexmod v0.0.0-20240114003000-22e6dd072f0c
	github.com/SupremeERG/opshins v0.1.0-beta.1
	github.com/dlclark/regexp2 v1.10.0
)

require (
	github.com/SupremeERG/colorPrint v1.0.0 // indirect
	golang.org/x/exp v0.0.0-20230728194245-b0cb94b80691 // indirect
)
