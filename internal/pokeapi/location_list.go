package pokeapi

import (
	"encoding/json"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreasResponse, error) {
	url := BaseUrl + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := LocationAreasResponse{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationAreasResponse{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer res.Body.Close()

	var locationsResponse LocationAreasResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationsResponse); err != nil {
		return LocationAreasResponse{}, err
	}

	toCache, _ := json.Marshal(locationsResponse)
	c.cache.Add(url, toCache)
	return locationsResponse, nil
}
