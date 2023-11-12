package fetchcode

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Function to fetch JS code from a URL
func FetchJSFromURL(url string, channel chan string) { //(string, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	channel <- string(body) //return string(body), nil
}

func FetchPatterns() ([]string, error) {
	file, err := os.Open("regex.txt") // Replace with your file path
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var patterns []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		patterns = append(patterns, scanner.Text())
	}

	return patterns, scanner.Err()

}
