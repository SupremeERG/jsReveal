package fetchcode

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/SupremeERG/jsReveal/pkg/regexmod"
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

func FetchPatterns(regexFile string) ([]string, string, error) {
	file, err := os.Open(regexFile) // Replace with your file path
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	var patterns []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		patterns = append(patterns, scanner.Text())
	}

	return patterns, regexFile, scanner.Err()

}

func FetchPatternsFromJSON() (map[string]regexmod.RegexProperties, error) {
	regexJSON, err := fs.ReadFile(os.DirFS("."), "regex.json")
	if err != nil {
		log.Fatal("Error reading regex file: ", err)
		//fmt.Println("Error reading regex file:", err)
	}

	var categories map[string]regexmod.RegexProperties
	err = json.Unmarshal(regexJSON, &categories)
	if err != nil {
		fmt.Println("Error parsing regex JSON:", err)
		return nil, err
	}

	return categories, nil
}
