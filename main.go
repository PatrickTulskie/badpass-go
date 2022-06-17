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

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("<h1>Welcome to the password checker!</h1>"))
// }

func passwordCheckHandler(response http.ResponseWriter, request *http.Request) {
	responseMap := map[string]interface{}{}
	if request.Method == "POST" {
		request.ParseForm()
		testHash := request.Form.Get("hash")
		if len(testHash) == 0 {
			testWord := request.Form.Get("password")
			fmt.Println("Word: " + testWord)
			if len(testWord) == 0 {
				response.WriteHeader(400)
				return
			}
			testHash = hashPassword(testWord)
		}
		fmt.Println("Hash: " + testHash)
		responseMap["common"] = findHash(testHash)

	} else {
		responseMap["error"] = "Bad Request"
		response.WriteHeader(404)
	}

	// Convert response to JSON and send it
	responseJson, _ := json.Marshal(responseMap)
	response.Write([]byte(responseJson))
}
func hashPassword(testWord string) string {
	hasher := sha1.New()
	hasher.Write([]byte(testWord))

	return hex.EncodeToString(hasher.Sum(nil))
}

func findHash(testHash string) bool {
	// Look for it in the list of hashes
	for _, hash := range wordlist10kHashes {
		if hash == testHash {
			return true
		}
	}

	return false
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
	mux.Handle("/", http.FileServer(http.Dir("static")))
	mux.HandleFunc("/password-check", passwordCheckHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
