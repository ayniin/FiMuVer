package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"fimuver/internal/config"
)

type TVDBClient struct {
	apiKey  string
	baseURL string
	client  *http.Client

	mu    sync.Mutex
	token string
}

type tvdbLoginResponse struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

type tvdbSearchResultStruct struct {
	TvdbId   string `json:"tvdb_id"`
	Name     string `json:"name"`
	Overview string `json:"overview"`
	Year     int    `json:"year"`
	ImageURL string `json:"image_url"`
	Type     string `json:"type"` //series / movie
}

// UnmarshalJSON überschreibt das Standard-Unmarshalling, weil TVDB "year"
// als String (z.B. "2020") statt als Zahl liefert.
func (r *tvdbSearchResultStruct) UnmarshalJSON(data []byte) error {
	type alias tvdbSearchResultStruct
	aux := &struct {
		Year string `json:"year"`
		*alias
	}{
		alias: (*alias)(r),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	if aux.Year != "" {
		year, err := strconv.Atoi(aux.Year)
		if err != nil {
			return fmt.Errorf("error parsing year %q: %v", aux.Year, err)
		}
		r.Year = year
	}

	return nil
}

type tvdbSearchResponse struct {
	Data []tvdbSearchResultStruct `json:"data"`
}

func NewTVDBClient(cfg *config.TVDBConfig) *TVDBClient {
	return &TVDBClient{
		apiKey:  cfg.TvdbApiKey,
		baseURL: cfg.BaseURL,
		client:  &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *TVDBClient) login() (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.token != "" {
		return c.token, nil
	}

	body, err := json.Marshal(map[string]string{"apikey": c.apiKey})
	if err != nil {
		return "", fmt.Errorf("error marshalling login request: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, c.baseURL+"/login", bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("error creating login request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error while login request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("login request failed with status: %s", resp.Status)
	}

	var loginResp tvdbLoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		return "", fmt.Errorf("error decoding login response: %v", err)
	}

	c.token = loginResp.Data.Token
	return c.token, nil

}

func (c *TVDBClient) doRequest(path string, out interface{}) error {
	token, err := c.login()
	if err != nil {
		return err
	}

	resp, err := c.get(path, token)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		c.mu.Lock()
		c.token = ""
		c.mu.Unlock()

		token, err = c.login()
		if err != nil {
			return err
		}

		resp, err = c.get(path, token)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error while tvdb-request: status %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("error decoding tvdb response: %v", err)
	}

	return nil
}

func (c *TVDBClient) get(path, token string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, c.baseURL+path, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating GET request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")

	return c.client.Do(req)
}

func (c *TVDBClient) SearchSeries(query string) ([]tvdbSearchResultStruct, error) {
	var resp tvdbSearchResponse
	path := fmt.Sprintf("/search?query=%s&type=series", url.QueryEscape(query))
	if err := c.doRequest(path, &resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (c *TVDBClient) SearchMovies(query string) ([]tvdbSearchResultStruct, error) {
	var resp tvdbSearchResponse
	path := fmt.Sprintf("/search?query=%s&type=movie", url.QueryEscape(query))
	if err := c.doRequest(path, &resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}
