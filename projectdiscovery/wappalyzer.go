package main

import (
	"fmt"
	"github.com/projectdiscovery/wappalyzergo"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.DefaultClient.Get("https://www.geeksforgeeks.org/how-to-print-string-with-double-quotes-in-golang/")
	if err != nil {
		log.Fatal(err)
	}
	data, _ := ioutil.ReadAll(resp.Body) // Ignoring error for example

	wappalyzerClient, err := wappalyzer.New()
	fingerprints := wappalyzerClient.Fingerprint(resp.Header, data)
	fmt.Printf("%v\n", fingerprints)

	// Output: map[Acquia Cloud Platform:{} Amazon EC2:{} Apache:{} Cloudflare:{} Drupal:{} PHP:{} Percona:{} React:{} Varnish:{}]

}