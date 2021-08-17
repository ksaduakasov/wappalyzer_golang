package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

//59c73eb9ec2841e4b22396c36f27e6d2

func main() {
	process("https://api.similartech.com/v1/site/wtotem.com/technologies?userkey=3b3b8cdb2b9b43dbad4a6138ceab60ea&format=json")
}

func process(url string) {
	req, err := http.Get(url)
	HandleErr(err)
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	HandleErr(err)

	//fmt.Println(string(html))
	var out bytes.Buffer
	json.Indent(&out, body, "", "    ")
	out.WriteTo(os.Stdout)

	//fmt.Println(string(html))
}

func HandleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
