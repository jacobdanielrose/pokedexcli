package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := BaseUrl + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		pokemonResponse := Pokemon{}
		err := json.Unmarshal(val, &pokemonResponse)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResponse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, nil
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, nil
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonResponse := Pokemon{}
	json.Unmarshal(dat, &pokemonResponse)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, dat)

	return pokemonResponse, nil
}
