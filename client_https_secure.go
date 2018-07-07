package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://httpbin.org/get"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "99fc2220-a8c5-454b-b010-8b8c3bb969b1")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
