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

type Entity struct {
	Results[] Results
}

type Results struct {

}

func main() {
	process("https://whatcms.org/API/Tech?key=8291b3042eaf6d82b37f6af754422e479f2e454a7f094e0b8d8e978636806b79f54943&url=https://www.geeksforgeeks.org/how-to-print-string-with-double-quotes-in-golang/")
}

func process(url string) {
	req, err := http.Get(url)
	HandleErr(err)
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	HandleErr(err)
	var out bytes.Buffer
	json.Indent(&out, body, "", "    ")
	out.WriteTo(os.Stdout)

	//sb := string(body)
	//var entity[] Entity
	//err = json.Unmarshal([]byte(sb), &entity)
	//HandleErr(err)
	//fmt.Println(entity)

}

func HandleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
