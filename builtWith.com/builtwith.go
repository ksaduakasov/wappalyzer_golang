package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	process("https://api.builtwith.com/v19/api.json?KEY=4d94c975-7c4c-4a57-b8d9-9091acfdc4ac&LOOKUP=wtotem.com")
}

func process(url string) {
	req, err := http.Get(url)
	HandleErr(err)
	defer req.Body.Close()

	html, err := io.ReadAll(req.Body)
	HandleErr(err)

	fmt.Println(string(html))
}

func HandleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}


