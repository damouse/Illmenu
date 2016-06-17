package illmenu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const baseUrlString string = "https://www.googleapis.com/customsearch/v1"

var params map[string]string = map[string]string{
	"key":        "AIzaSyBSX7Xr-juL_3oTB3ZQaj5Wmcv1__kfR0w",
	"cx":         "013062920850834713833:hsd8bzw_nki",
	"safe":       "high",
	"searchType": "image",
}

// Searches google images with the given query, and returns a list of image urls
// Link to CSE docs: https://developers.google.com/custom-search/json-api/v1/reference/cse/list#request
func ImageSearch(query string) ([]string, error) {
	base, e := url.Parse(baseUrlString)
	if e != nil {
		return nil, e
	}

	// Build the query
	q := base.Query()
	q.Add("q", query)

	for k, v := range params {
		q.Add(k, v)
	}

	base.RawQuery = q.Encode()

	// Make the request
	fullJson := map[string]interface{}{}

	if res, e := http.Get(base.String()); e != nil {
		return nil, e
	} else if body, e := ioutil.ReadAll(res.Body); e != nil {
		return nil, e
	} else if e := json.Unmarshal(body, &fullJson); e != nil {
		return nil, e
	} else if res.Status != "200 OK" {
		return nil, fmt.Errorf("Google search failed: %v", fullJson)
	}

	// Extract the values from the json
	links := make([]string, 0)

	if results, ok := fullJson["items"].([]interface{}); !ok {
		fmt.Println("Cast failed! heres the json: ", fullJson)
		return nil, e
	} else {
		for _, result := range results {
			if j, ok := result.(map[string]interface{}); !ok {
				return nil, fmt.Errorf("Cast failed! Bad result type")
			} else if link, ok := j["link"]; !ok {
				return nil, e
			} else if s, ok := link.(string); !ok {
				return nil, e
			} else {
				links = append(links, s)
			}
		}
	}

	return links, nil
}
