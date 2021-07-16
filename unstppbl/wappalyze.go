package main

import (
	"encoding/json"
	"github.com/unstppbl/gowap"
	"log"
)

func main() {
	url := "https://wtotem.com"
	wapp, err := gowap.Init("./unstppbl/apps.json", true)
	if err != nil {
		log.Fatal(err)
	}
	res, err := wapp.Analyze(url)
	if err != nil {
		log.Fatal(err)
	}
	prettyJSON, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[*] Result for %s:\n%s", url, string(prettyJSON))
}
