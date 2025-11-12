package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) LocationArea(loc string) (LocationResp, error) {
	endpoint := "/location-area/"
	fullURL := baseURL + endpoint + loc
	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("data gotten from cache")
	}
	if !ok {
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return LocationResp{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationResp{}, err
		}
		defer resp.Body.Close()
		if resp.StatusCode > 399 {
			return LocationResp{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
		}
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationResp{}, err
		}
	}
	locationResp := LocationResp{}
	err := json.Unmarshal(data, &locationResp)
	if err != nil {
		return LocationResp{}, err
	}
	if !ok {
		c.cache.Add(fullURL, data)
	}
	return locationResp, nil
}
