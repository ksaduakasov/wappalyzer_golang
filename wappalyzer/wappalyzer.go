package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	process("https://api.wappalyzer.com/lookup/v2/?urls=https://www.wtotem.com")
}

func process(url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("X-Api-Key", "uJYNOLkDae4BkNGmLIybU8sijWH83g3B8HL5bOy8")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	sb := string(body)
	log.Printf(sb)
}
