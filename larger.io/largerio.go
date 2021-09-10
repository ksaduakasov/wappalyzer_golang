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
	process("https://api.larger.io/v1/search/key/G0KM2YWNWWAFMW6DZJ7Z66XPCE6ADOJE?domain=wtotem.com")
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


