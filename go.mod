module github.com/SupremeERG/jsReveal

go 1.21.3

toolchain go1.21.4

// replace github.com/SupremeERG/jsReveal/pkg/regexmod => ./pkg/regexmod

// replace github.com/SupremeERG/jsReveal/pkg/fetchcode => ./pkg/fetchcode

// replace github.com/SupremeERG/jsReveal/internal/parser => ./internal/parser

// replace github.com/SupremeERG/jsReveal/pkg/misc => ./pkg/misc

// replace github.com/SUpremeERG/jsReveal/runner => ./runner

require (
	github.com/SupremeERG/jsReveal/internal/parser v0.0.0-20240112200251-42e506d4e376
	github.com/SupremeERG/jsReveal/pkg/fetchcode v0.0.0-20240112200251-42e506d4e376
	github.com/SupremeERG/jsReveal/pkg/misc v0.0.0-20240112200251-42e506d4e376
)

require (
	github.com/SupremeERG/colorPrint v1.0.0 // indirect
	github.com/SupremeERG/jsReveal/pkg/regexmod v0.0.0-20240112200251-42e506d4e376 // indirect
	github.com/SupremeERG/opshins v1.0.0 // indirect
	github.com/dlclark/regexp2 v1.10.0 // indirect
	golang.org/x/exp v0.0.0-20240112132812-db7319d0e0e3 // indirect
)
