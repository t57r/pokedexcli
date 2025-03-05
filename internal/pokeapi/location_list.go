package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := locationAreaURL
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}

// ExploreLocation -
func (c *Client) ExploreLocation(idOrName *string) (RespExploredLocation, error) {
	url := locationAreaURL + fmt.Sprintf("/%s", *idOrName)

	if val, ok := c.cache.Get(url); ok {
		locationResp := RespExploredLocation{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespExploredLocation{}, err
		}

		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespExploredLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespExploredLocation{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespExploredLocation{}, err
	}

	exploredLocationResp := RespExploredLocation{}
	err = json.Unmarshal(dat, &exploredLocationResp)
	if err != nil {
		return RespExploredLocation{}, err
	}

	c.cache.Add(url, dat)
	return exploredLocationResp, nil
}
