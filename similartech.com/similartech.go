package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

//59c73eb9ec2841e4b22396c36f27e6d2

func main() {
	process("https://api.similartech.com/v1/site/wtotem.com/technologies?userkey=59c73eb9ec2841e4b22396c36f27e6d2&format=json")
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
