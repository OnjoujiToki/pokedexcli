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
		return LocationAreasResponse{}, fmt.Errorf("Bad response status code: %d", response.StatusCode)
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
