package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func HandleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

type Entity struct {
	Technologies []struct {
		Name       string        `json:"name"`
		Versions   []interface{} `json:"versions"`
		Categories []struct {
			ID   int    `json:"id"`
			Slug string `json:"slug"`
			Name string `json:"name"`
		} `json:"categories"`
	} `json:"technologies"`
}

func main() {
	api := "https://whatcms.org/API/Tech?key=8291b3042eaf6d82b37f6af754422e479f2e454a7f094e0b8d8e978636806b79f54943&url="
	url := "https://www.geeksforgeeks.org/how-to-print-string-with-double-quotes-in-golang/"
	apiWithLink := api + url
	//req, err := http.Get(apiWithLink)
	//HandleErr(err)
	//defer req.Body.Close()
	//
	//body, err := io.ReadAll(req.Body)
	//HandleErr(err)

	//url := "https://api.larger.io/v1/search/key/G0KM2YWNWWAFMW6DZJ7Z66XPCE6ADOJE?domain=telegram.org"
	req, err := http.Get(apiWithLink)
	HandleErr(err)
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	HandleErr(err)

	var myInterface map[string]interface{}
	err = json.Unmarshal(body, &myInterface)
	HandleErr(err)

	//fmt.Println(myInterface)

	var results []interface{}

	for key, value := range myInterface {
		if key == "results" {
			temp := value.([]interface{})
			results = temp
		}
	}

	fmt.Println(results)


	for _, el := range results {
		var categories []string
		var name string
		var version string
		for key, value := range el.(map[string]interface{}) {
			if key == "categories" {
				//var tempString string
				for _, v := range value.([]interface{}) {
					categories = append(categories, v.(string))
				}
			}

			if key == "name" {
				name = value.(string)
			}

			if key == "version" {
				version = value.(string)
			}
		}

		fmt.Println(categories)
		if version != "" {
		fmt.Println("version: ", version)
		} else {
			fmt.Println("no versions are available")
		}
		fmt.Println("name: ", name)
	}



	//var out bytes.Buffer
	//json.Indent(&out, body, "", "    ")
	//out.WriteTo(os.Stdout)

	//var example []map[string]interface{}
	//err = json.Unmarshal(body, &example)
	//HandleErr(err)
	////fmt.Println(example)
	//var name string
	//var versions string
	//var categories []string
	//for key, value := range example[0] {
	//
	//
	//	if key == "technologies" {
	//		for _, s := range value.([]interface{}) {
	//			for k, v := range s.(map[string]interface{}) {
	//				if k == "categories" {
	//					for _, temp := range v.([]interface{}) {
	//						for qwe, asd := range temp.(map[string]interface{}) {
	//							tempString := fmt.Sprintf("%v : %v", qwe, asd)
	//							categories = append(categories, tempString)
	//						}
	//					}
	//				}
	//				if k == "versions" {
	//					versions = fmt.Sprintf("%v", v)
	//					//if tempString != "[]" {
	//					//	fmt.Println(tempString)
	//					//}
	//				}
	//
	//				if k == "name" {
	//					name = fmt.Sprintf("name: %v ", v)
	//				}
	//
	//			}
	//			fmt.Println(name, "name -------------------------------------------------")
	//
	//			fmt.Println(versions, "versions -------------------------------------------------")
	//			fmt.Println(categories, "categories -------------------------------------------------")
	//
	//		}
	//
	//
	//	}
	//}

	//req, err := http.Get("https://api.similartech.com/v1/site/wtotem.com/technologies?userkey=9ebc7232876b470d8da18c473999a1e1&format=json")
	//HandleErr(err)
	//defer req.Body.Close()
	//
	//body, err := io.ReadAll(req.Body)
	//HandleErr(err)
	//
	//var myInterface map[string]interface{}
	//err = json.Unmarshal(body, &myInterface)
	//HandleErr(err)
	////fmt.Println(myInterface)
	//
	//var myTemp []interface{}
	//
	//for key, value := range myInterface {
	//	if key == "technologies" {
	//		temp := value.([]interface{})
	//		myTemp = temp
	//	}
	//}
	//
	////fmt.Println(myTemp)
	//for _, el := range myTemp {
	//	var categories []string
	//	for key, value := range el.(map[string]interface{}) {
	//		if key == "categories" {
	//			tempString := fmt.Sprintf("%v", value)
	//			categories = append(categories, tempString)
	//		}
	//	}
	//	fmt.Println(categories)
	//}

	//var myInterface map[string]interface{}
	//err = json.Unmarshal(body, &myInterface)
	//HandleErr(err)
	//
	////fmt.Println(myInterface)
	//
	//var results []interface{}
	//
	//for key, value := range myInterface {
	//	if key == "results" {
	//		temp := value.([]interface{})
	//		results = temp
	//	}
	//}

	//fmt.Println(results)


	//for _, el := range results {
	//	var categories []string
	//	var name string
	//	var version string
	//	for key, value := range el.(map[string]interface{}) {
	//		if key == "categories" {
	//			//var tempString string
	//			for _, v := range value.([]interface{}) {
	//				categories = append(categories, v.(string))
	//			}
	//		}
	//
	//		if key == "name" {
	//			name = value.(string)
	//		}
	//
	//		if key == "version" {
	//			version = value.(string)
	//		}
	//	}
	//
	//	fmt.Println(categories)
	//	if version != "" {
	//	fmt.Println("version: ", version)
	//	} else {
	//		fmt.Println("no versions are available")
	//	}
	//	fmt.Println("name: ", name)
	//}


}
