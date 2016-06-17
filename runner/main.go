package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hokaccha/go-prettyjson"
)

// Link to CSE docs: https://developers.google.com/custom-search/json-api/v1/reference/cse/list#request

const apiKey string = "AIzaSyBSX7Xr-juL_3oTB3ZQaj5Wmcv1__kfR0w"

func main() {
	// illmenu.TestSearch()
	// Other important keys: safe and searchType
	searchUrl := "https://www.googleapis.com/customsearch/v1?q=dog&cx=013062920850834713833:hsd8bzw_nki&key=AIzaSyBSX7Xr-juL_3oTB3ZQaj5Wmcv1__kfR0w"

	res, err := http.Get(searchUrl)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	s, err := getStations([]byte(body))
	if err != nil {
		panic(err.Error())
	}

	p, _ := prettyjson.Marshal(s)
	fmt.Println(string(p))
}

func getStations(body []byte) (map[string]interface{}, error) {
	s := map[string]interface{}{}

	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}
