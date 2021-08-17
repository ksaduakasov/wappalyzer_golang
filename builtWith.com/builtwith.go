package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	process("https://api.builtwith.com/v19/api.json?KEY=c4f0023c-b4d4-44b7-b5bd-4b0e1d4a8d2b&LOOKUP=wtotem.com")
}

func process(url string) {
	req, err := http.Get(url)
	HandleErr(err)
	defer req.Body.Close()

	html, err := io.ReadAll(req.Body)
	HandleErr(err)
	//fmt.Println(string(html))
	var out bytes.Buffer
	json.Indent(&out, html, "", "    ")
	out.WriteTo(os.Stdout)

	//fmt.Println(string(html))
}

func HandleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}


