package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func fetchLocationAreaName(url string) (locationsNames []string, previousUrl string, nextUrl string, err error) {

	res, err := http.Get(url)
	if err != nil {
		return []string{}, "", "", fmt.Errorf("API response error")
	}

	dataJson, err := io.ReadAll(res.Body)
	//var data map[string]any
	var data struct {
		Count    float64             `json:"count"`
		Next     string              `json:"next"`
		Previous string              `json:"previous"`
		Results  []map[string]string `json:"results"`
		rest     json.RawMessage
	}

	err = json.Unmarshal(dataJson, &data)
	if err != nil {
		return []string{}, "", "", fmt.Errorf("Error while parsing json")
	}

	for _, e := range data.Results {
		for k, v := range e {
			if k == "name" {
				locationsNames = append(locationsNames, v)
			}
		}
	}

	return locationsNames, data.Previous, data.Next, nil
}
