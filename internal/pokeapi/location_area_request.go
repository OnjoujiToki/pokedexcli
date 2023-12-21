package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}
	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("Using cached data")
		locationAreaResp := LocationAreasResponse{}
		err := json.Unmarshal(data, &locationAreaResp)
		if err != nil {
			return LocationAreasResponse{}, err
		}
		return locationAreaResp, nil
	}
	request, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	response, err := c.httpClient.Do(request)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)
	if response.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad response status code: %d", response.StatusCode)
	}
	data, err = io.ReadAll(response.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	// save the data to the cache
	c.cache.Add(fullURL, data)

	locationAreaResp := LocationAreasResponse{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	return locationAreaResp, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("Using cached data")
		locationArea := LocationArea{}
		err := json.Unmarshal(data, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}
	request, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}
	response, err := c.httpClient.Do(request)
	if err != nil {
		return LocationArea{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)
	if response.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("Bad response status code: %d", response.StatusCode)
	}
	data, err = io.ReadAll(response.Body)
	if err != nil {
		return LocationArea{}, err
	}
	// save the data to the cache
	c.cache.Add(fullURL, data)

	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	return locationArea, nil
}

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("Using cached data")
		pokemon := Pokemon{}
		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}
	request, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}
	response, err := c.httpClient.Do(request)
	if err != nil {
		return Pokemon{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	if response.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad response status code: %d", response.StatusCode)
	}
	data, err = io.ReadAll(response.Body)
	if err != nil {
		return Pokemon{}, err
	}
	// save the data to the cache
	c.cache.Add(fullURL, data)

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}
