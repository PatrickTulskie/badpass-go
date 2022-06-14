# badpass-go

Toy web app to learn how to read from text files, scan arrays, and generally tell you if a password is a common password.

To use, download the 10k most common passwords first:

    wget https://raw.githubusercontent.com/danielmiessler/SecLists/master/Passwords/Common-Credentials/10k-most-common.txt -O data/10k-most-common.txt

Precompute Hashes With:

    go run precompute_hashes.go

Then run the server with

    go run main.go

Finally, to query it

    curl "localhost:8080/password-check?password=butthead"

Enjoy
