package runner

import (
	"log"
)

func Run() Options {
	opts := ParseOptions()

	if opts.Verbose {
		log.SetFlags(log.Ltime | log.Lshortfile)
	} else {
		log.SetFlags(0)
	}

	log.SetPrefix("jsReveal --- ")

	return opts
}
