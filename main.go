package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func HandleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	req, err := http.Get("https://api.builtwith.com/free1/api.json?KEY=4d94c975-7c4c-4a57-b8d9-9091acfdc4ac&LOOKUP=wtotem.com")
	HandleErr(err)
	//req.Header.Add("x-api-key", "uJYNOLkDae4BkNGmLIybU8sijWH83g3B8HL5bOy8")
	defer req.Body.Close()

	html, err := io.ReadAll(req.Body)
	HandleErr(err)

	fmt.Print(string(html))
}