package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// "encoding/json"
// "io"
// "net/http"

func (c *Client) GetALocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		locationResp := Location{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return Location{}, err
		}

		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resq, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}

	defer resq.Body.Close()

	dat, err := io.ReadAll(resq.Body)
	if err != nil {
		return Location{}, err
	}

	location := Location{}
	err = json.Unmarshal(dat, &location)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, dat)

	return location, nil
}
