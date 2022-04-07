package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

var wordlist10k []string

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the password checker!</h1>"))
}

func passwordCheckHandler(w http.ResponseWriter, r *http.Request) {
	testWord := r.URL.Query().Get("password")

	for _, word := range wordlist10k {
		if word == testWord {
			w.Write([]byte("{ common: true }"))
			return
		}
	}

	w.Write([]byte("{ common: false }"))
}

func setup10kWordList() {
	fmt.Println("Reading Wordlist...")

	file, err := os.Open("10k-most-common.txt")
	if err != nil {
		log.Fatalf("couldn't open wordlist")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		wordlist10k = append(wordlist10k, scanner.Text())
	}

	file.Close()
}

func main() {
	setup10kWordList()
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/password-check", passwordCheckHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
