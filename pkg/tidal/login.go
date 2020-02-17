package tidal

import (
	"encoding/json"
	"net/url"
)

func (c *Client) Login() error {

	var loginResponse *User

	resp, err := c.post("/login/username",
		url.Values{
			"username": {c.Username},
			"password": {c.Password},
		},
		url.Values{
			"X-Tidal-Token": {c.Token},
		},
	)

	if err != nil {
		return err
	}

	err = json.Unmarshal(resp, &loginResponse)

	if err != nil {
		return err
	}

	c.User = loginResponse

	return nil
}
