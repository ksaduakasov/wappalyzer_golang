package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Entity struct {
	Technologies []Technologies
}

type Technologies struct {
	Slug string
	Name string
	Versions []string
	TrafficRank int
	ConfirmedAt int
	Categories []Categories
}

type Categories struct {
	Id int
	Slug string
	Name string
}

func main() {
	process("https://api.wappalyzer.com/lookup/v2/?urls=https://telegram.org/")
}

func process(url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("X-Api-Key", "lNhY8pKX65167nFKv5VyJ3O6LIJbyWlX88mBBoft")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out, body, "", "    ")
	out.WriteTo(os.Stdout)

	//sb := string(body)
	//var entity[] Entity
	//err = json.Unmarshal([]byte(sb), &entity)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(entity)

}
