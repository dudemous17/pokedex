package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// Global configuration to maintain the pagination state
var config = struct {
	NextLocationURL     *string
	PreviousLocationURL *string
}{
	NextLocationURL:     nil,
	PreviousLocationURL: nil,
}

func init() {
	// Initial URL for the first call to location areas
	*config.NextLocationURL = "https://pokeapi.co/api/v2/location-area/"
	*config.PreviousLocationURL = "https://pokeapi.co/api/v2/location-area/"
}

func commandMap() error {
	if config.NextLocationURL == nil || *config.NextLocationURL == "" {
		fmt.Println("You have reached the end of the location list.")
		return nil
	}

	//fmt.Printf("Fetching locations from: %s\n", *nextURL)
	response, err := http.Get(*config.NextLocationURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("API request failed with status: %s\n", response.Status)
		return nil
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var locationResponse LocationAreasResponse
	err = json.Unmarshal(body, &locationResponse)
	if err != nil {
		return err
	}

	// Display the names of the locations
	for _, location := range locationResponse.Results {
		fmt.Printf("- %s\n", location.Name)
	}

	// Update the global variable with the URL for the next page
	config.NextLocationURL = locationResponse.Next
	return nil
}
