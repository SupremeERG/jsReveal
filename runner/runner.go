package runner

import (
	"log"
)

func Run() Options {
	opts := ParseOptions()

	if opts.Verbose {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	} else {
		log.SetFlags(0)
	}

	log.SetPrefix("jsreveal")

	return opts
	/*
		switch opts.Source {
		default:
			fmt.Println("./jsReveal -u <url to JS file>")
		case 1: // single file
			parser.ParseJS(opts.jsFilePath)
		case 2:
			fmt.Println("multiple urls")
		case 3:
			fmt.Println("single url")
		}*/
}
