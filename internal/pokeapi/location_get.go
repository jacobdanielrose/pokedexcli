package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(location string) (LocationAreaResponse, error) {
	url := BaseUrl + "/location-area/" + location

	if val, ok := c.cache.Get(url); ok {
		locationResp := LocationAreaResponse{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)

	c.cache.Add(url, dat)

	locationAreaResponse := LocationAreaResponse{}
	json.Unmarshal(dat, &locationAreaResponse)

	return locationAreaResponse, nil
}
