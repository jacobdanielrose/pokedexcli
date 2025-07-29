package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (PokemonResponse, error) {
	url := BaseUrl + "/pokemon/" + name

	if val, ok := c.cache.Get(url); ok {
		pokemonResponse := PokemonResponse{}
		err := json.Unmarshal(val, &pokemonResponse)
		if err != nil {
			return PokemonResponse{}, err
		}

		return pokemonResponse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonResponse{}, nil
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, nil
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)

	c.cache.Add(url, dat)

	pokemonResponse := PokemonResponse{}
	json.Unmarshal(dat, &pokemonResponse)

	return PokemonResponse{}, nil
}
