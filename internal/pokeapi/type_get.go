package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetType(pokemonType string) (Type, error) {
	url := baseURL + "/type/" + pokemonType

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Type{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Type{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Type{}, err
	}

	typeResp := Type{}
	err = json.Unmarshal(dat, &typeResp)
	if err != nil {
		return Type{}, err
	}

	return typeResp, nil
}
