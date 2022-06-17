# badpass-go

Toy web app to learn how to read from text files, scan arrays, and generally tell you if a password is a common password.

To use, download the 10k most common passwords first:

    wget https://raw.githubusercontent.com/danielmiessler/SecLists/master/Passwords/Common-Credentials/10k-most-common.txt -O data/10k-most-common.txt

Precompute Hashes With:

    go run precomputeHashes.go

Then run the server with

    go run main.go

Finally, to query it, just use the index page or cURL

    curl -v -XPOST localhost:8080/password-check -d 'hash=583adc8aebb04a62cc76e71314b46474113be146'
    curl -v -XPOST localhost:8080/password-check -d 'password=butthead'

Enjoy
