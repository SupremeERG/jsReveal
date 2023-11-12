package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	// "strings"
	"github.com/SupremeERG/jsReveal/pkg/fetchcode"
	"github.com/SupremeERG/jsReveal/pkg/parse"
)

// function to read JS links from a file
func readJSLinks(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var links []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		links = append(links, scanner.Text())
	}

	return links, scanner.Err()
}

func main() {
	jsFilePath := flag.String("f", "", "Path to the .js file") // planning to add all the options to a module called "runner"
	jsLinksPath := flag.String("l", "", "Path to the file with JS links")
	jsURL := flag.String("u", "", "URL to a JS file")
	verbose := flag.Bool("v", false, "Enable verbose output")
	flag.Parse()

	if *verbose {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	} else {
		log.SetFlags(0)
	}

	log.SetPrefix("jsreveal: ")

	// Check if the JS links file path is provided
	if *jsLinksPath != "" {
		if *verbose {
			log.Println("Processing JS links from file:", *jsLinksPath)
		}

		links, err := readJSLinks(*jsLinksPath)
		if err != nil {
			log.Fatal("Error reading JS links: ", err)
		}

		ch := make(chan string)
		for _, link := range links {
			go fetchcode.FetchJSFromURL(link, ch)

			go parse.ParseJSFromCode(<-ch) // Assuming parseJSFromCode accepts a string of JS code
		}
		return
	} else if *jsURL != "" {
		ch := make(chan string)
		fetchcode.FetchJSFromURL(*jsURL, ch)

		parse.ParseJSFromCode(<-ch)
		return
	}

	if *jsFilePath == "" {
		fmt.Println("Usage: go run main.go -f <input_file.js> [-v for verbose]\n-f - file with js code\n-l - file with links to js code\n-u - link to js code\n-v - for verbose")
		return
	}

	parse.ParseJS(*jsFilePath)
}
