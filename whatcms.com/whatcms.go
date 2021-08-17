package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

//8291b3042eaf6d82b37f6af754422e479f2e454a7f094e0b8d8e978636806b79f54943

func main() {
	process("https://whatcms.org/API/Tech?key=8291b3042eaf6d82b37f6af754422e479f2e454a7f094e0b8d8e978636806b79f54943&url=wtotem.com")
}

func process(url string) {
	req, err := http.Get(url)
	HandleErr(err)
	//req.Header.Add("x-api-key", "uJYNOLkDae4BkNGmLIybU8sijWH83g3B8HL5bOy8")
	defer req.Body.Close()

	html, err := io.ReadAll(req.Body)
	HandleErr(err)
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
