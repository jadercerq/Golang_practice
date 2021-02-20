package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type repo struct {
	name        string
	description string
}

func main() {
	fmt.Println("Calling API...")

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.github.com/orgs/jadercerq/repos", nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, _ := client.Do(req)

	println(resp.Body)

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var responseObject repo
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("API Response as struct %+v\n", responseObject)
}
