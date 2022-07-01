package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Meta struct {
	Status int `json:"status"`
}

type Request struct {
	Meta Meta `json:"meta"`
	Data []Data `json:"data"`
}

type Data struct {
	Slug     string   `json:"slug"`
	IsCommon bool     `json:"is_common"`
	Tags     []string `json:"tags"`
	Japanese []Japanese
}

type Japanese struct {
	Word    string `json:"word"`
	Reading string `json:"reading"`
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("please provide a word to define") 
	}
	word := os.Args[1] 
	results, err := search(word)
	if err != nil {
		log.Fatalf("failed to get information for %s\n", word) 
	}
	if len(results.Data) == 0 || len(results.Data[0].Japanese) == 0 {
		log.Fatalf("unable to find a definition for %s", word) 
	}
	jap := results.Data[0].Japanese[0]

	if jap.Word == "" {

	fmt.Printf("%s\n", jap.Reading)
	} else {

	fmt.Printf("%s - %s\n", jap.Word, jap.Reading)
	}
}

func search(word string) (*Request, error) {
	uri := "https://jisho.org/api/v1/search/words"
	resp, err := http.Get(fmt.Sprintf("%s?keyword=%s", uri, word))
	if err != nil {
		return nil, fmt.Errorf("failed to GET %q: %w", word, err)
}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
}

	data := &Request{}
	json.Unmarshal(body, data)
	
return data, nil
}
