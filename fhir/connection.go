package fhir

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Connection struct {
	BaseURL string
	client  *http.Client
}

func (c *Connection) Get(q string) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%v%v", c.BaseURL, q), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, nil
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
