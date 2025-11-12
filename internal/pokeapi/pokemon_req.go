package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Pokemon(pokemon string) (PokemonResp, error) {
	endpoint := "/pokemon/"
	fullURL := baseURL + endpoint + pokemon
	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("data gotten from cache")
	}
	if !ok {
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return PokemonResp{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return PokemonResp{}, err
		}
		defer resp.Body.Close()
		if resp.StatusCode > 399 {
			return PokemonResp{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
		}
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return PokemonResp{}, err
		}
	}
	pokemonResp := PokemonResp{}
	err := json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return PokemonResp{}, err
	}
	if !ok {
		c.cache.Add(fullURL, data)
	}
	return pokemonResp, nil
}
