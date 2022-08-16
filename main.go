package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type PASSWORDS struct {
	url      string
	username string
	password string
}

func main() {
	file, err := os.Open("old_password.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var lins []string
	passwords := []PASSWORDS{}
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		lins = append(lins, scanner.Text())
	}
	var URL string
	var USERNAME string
	var PASSWORD string
	for line := range lins {

		for i := 0; i < 10; i++ {
			if URL == "" {
				URL = strings.Join(strings.Split(lins[line], "Website URL: ")[1:], "Website URL: ")
			}
			if USERNAME == "" {
				USERNAME = strings.Join(strings.Split(lins[line], "Login: ")[1:], "Login: ")
			}
			if PASSWORD == "" {
				PASSWORD = strings.Join(strings.Split(lins[line], "Password: ")[1:], "Password: ")
			}
			if PASSWORD != "" {
				break
			}
		}
		if URL != "" && USERNAME != "" && PASSWORD != "" {
			passwords = append(passwords, PASSWORDS{URL, USERNAME, PASSWORD})
			URL = ""
			USERNAME = ""
			PASSWORD = ""
		}

	}

	f, err := os.Create("./PASSWORDS.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	w := csv.NewWriter(f)
	var header []string
	header = append(header, "url")
	header = append(header, "username")
	header = append(header, "password")
	w.Write(header)

	for _, obj := range passwords {
		var record []string
		record = append(record, obj.url, obj.username, obj.password)
		w.Write(record)
		record = nil
	}
	w.Flush()
}
