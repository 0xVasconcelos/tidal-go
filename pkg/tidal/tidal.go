package tidal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (c *Client) Search(search string) (*Search, error) {

	var result *Search

	resp, err := c.get(c.getEndpoint("/search"),
		url.Values{
			"query":       {search},
			"countryCode": {c.User.CountryCode},
		},
		url.Values{
			"X-Tidal-Token": {c.Token},
		},
	)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) GetStream(track *Track) (*Stream, error) {

	var stream *Stream

	resp, err := c.get(c.getEndpoint("/tracks/"+strconv.Itoa(track.ID)+"/streamUrl"),
		url.Values{
			"soundQuality": {c.SoundQuality},
			"countryCode":  {c.User.CountryCode},
			"sessionId":    {c.User.SessionID},
		},
		url.Values{
			"X-Tidal-Token": {c.Token},
		},
	)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &stream)

	if err != nil {
		return nil, err
	}

	return stream, nil
}

func (c *Client) GetMedia(stream *Stream) (Media, error) {

	var media Media

	resp, err := http.Get(stream.URL)

	if err != nil {
		return media, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return media, err
	}

	media.Media = body
	media.EncryptionKey = []byte(stream.EncryptionKey)

	return media, nil
}

func (c *Client) post(path string, form url.Values, headers map[string][]string) ([]byte, error) {

	client := http.Client{}

	req, err := http.NewRequest("POST", c.getEndpoint(path), strings.NewReader(form.Encode()))

	if err != nil {
		return nil, err
	}

	req.PostForm = form

	req.Header = headers

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
	}

	return body, err
}

func (c *Client) get(url string, query url.Values, headers map[string][]string) ([]byte, error) {

	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	req.Header = headers

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
	}

	return body, err
}

func (c *Client) getEndpoint(path string) string {
	return "https://api.tidalhifi.com/v1" + path
}
