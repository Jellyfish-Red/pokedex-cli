package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (client *APIClient) RequestArea(pageURL *string) (LocationSummary, error) {
	// Prepare request URL
	url := fmt.Sprint(baseURL, "/location-area")
	if pageURL != nil {
		url = *pageURL
	}
	
	//  Check if data already exists in cache. If no, send a GET request and store the data.
	data, dataFound := client.LocationCache.Get(url)
	if !dataFound {
		// Send GET request
		res, err := client.Client.Get(url)
		if err != nil {
			return LocationSummary{}, err
		}
	
		// Process data
		data, err = io.ReadAll(res.Body)
		defer res.Body.Close()
		if res.StatusCode > 299 {
			return LocationSummary{}, err
		}
		if err != nil {
			return LocationSummary{}, err
		}

		// Insert data into cache
		client.LocationCache.Add(url, data)
	} else {
		fmt.Println("DEBUG: Location Data found in cache; using cache data instead of performing GET request.")
	}

	// Package byte array into struct and send out
	summary := LocationSummary{}
	unmarshalErr := json.Unmarshal(data, &summary)
	if unmarshalErr != nil {
		return LocationSummary{}, unmarshalErr
	}

	return summary, nil
}