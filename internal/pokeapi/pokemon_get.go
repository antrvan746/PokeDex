package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetAPokemon(pokemonName string) (Pokemon, error) {

	url := baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resq, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer resq.Body.Close()

	dat, err := io.ReadAll(resq.Body)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, dat)

	pokemon := Pokemon{}
	err = json.Unmarshal(dat, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}