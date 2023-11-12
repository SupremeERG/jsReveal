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
		log.Fatal("Error fetching JS from URL: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		log.Println("Error fetching JS from URL: ", "404 File not Found")
		return
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Error reading JS: ", err)
		return
	}

	channel <- string(body) // string(body) is the JS code
	return
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
