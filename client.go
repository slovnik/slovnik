package slovnik

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"time"
)

// Client for accessing slovnik web server
type Client struct {
	client  *http.Client
	baseURL *url.URL
}

// NewClient creates new client for accessing slovnik web server
func NewClient(baseURL string) (*Client, error) {
	client := &Client{}
	client.client = &http.Client{Timeout: 10 * time.Second}

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	client.baseURL = u
	return client, nil
}

// Translate word
func (c *Client) Translate(word string) ([]Word, error) {
	const methodURL = "/api/translate/"
	u := *c.baseURL
	u.Path = path.Join(u.Path, methodURL, word)
	r, err := c.client.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Got bad status (%d) from server", r.StatusCode)
	}

	w := []Word{}
	json.NewDecoder(r.Body).Decode(&w)
	return w, nil
}
