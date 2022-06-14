package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func main() {
	precomputeHashes()
	fmt.Println("Done!")
}

func precomputeHashes() {
	fmt.Println("Reading Wordlist...")

	file, err := os.Open("data/10k-most-common.txt")
	if err != nil {
		log.Fatalf("couldn't open wordlist")
	}
	defer file.Close()

	outfile, err := os.Create("data/10k-most-common-sha1.txt")
	if err != nil {
		log.Fatalf("couldn't create new file")
	}
	defer outfile.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		hasher := sha1.New()
		hasher.Write([]byte(scanner.Text()))
		sha1Hash := hex.EncodeToString(hasher.Sum(nil)) + "\n"
		outfile.WriteString(sha1Hash)
	}

}
