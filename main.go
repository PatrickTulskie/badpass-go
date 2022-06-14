package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var wordlist10kHashes []string

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the password checker!</h1>"))
}

func passwordCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Setup the default response
	response := map[string]interface{}{
		"common": false,
	}

	// Get the query param, hash it
	testWord := r.URL.Query().Get("password")
	hasher := sha1.New()
	hasher.Write([]byte(testWord))
	testHash := hex.EncodeToString(hasher.Sum(nil))

	// Look for it in the list of hashes, break if found
	for _, hash := range wordlist10kHashes {
		if hash == testHash {
			response["common"] = true
			break
		}
	}

	// Convert response to JSON and send it
	responseJson, _ := json.Marshal(response)
	w.Write([]byte(responseJson))
}

func loadHashes() {
	fmt.Println("Reading Hash List...")

	file, err := os.Open("data/10k-most-common-sha1.txt")
	if err != nil {
		log.Fatalf("couldn't open wordlist")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		wordlist10kHashes = append(wordlist10kHashes, scanner.Text())
	}
}

func main() {
	loadHashes()
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/password-check", passwordCheckHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
