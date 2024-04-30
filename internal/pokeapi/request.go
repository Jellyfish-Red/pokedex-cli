package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

func (client *APIClient) requestArea(pageURL *string) (LocationSummary, error) {
	// Prepare request URL
	url := fmt.Sprint(baseURL, "/location-area")
	if pageURL != nil {
		url = *pageURL
	}

	// Send GET request
	res, err := client.Client.Get(url)
	if err != nil {
		return LocationSummary{}, err
	}

	// Process data
	data, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and body: %s\n", res.StatusCode, data)
	}
	if err != nil {
		log.Fatal(err)
	}

	summary := LocationSummary{}
	unmarshalErr := json.Unmarshal(data, &summary)
	if unmarshalErr != nil {
		fmt.Println(err)
	}

	return summary, nil
}