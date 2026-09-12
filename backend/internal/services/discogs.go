package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"fimuver/internal/config"
)

type DiscogsClient struct {
	token     string
	baseUrl   string
	userAgent string
	client    *http.Client
}

type discogsSearchResultStruct struct {
	DiscogsId int      `json:"id"`
	Title     string   `json:"title"`
	Year      string   `json:"year"`
	ImageURL  string   `json:"cover_image"`
	Type      string   `json:"type"` // release / master / artist / label
	Genres    []string `json:"genre"`
	Styles    []string `json:"style"`
	Country   string   `json:"country"`
	Format    []string `json:"format"`
}

type discogsSearchResponse struct {
	Results []discogsSearchResultStruct `json:"results"`
}

func NewDiscogsClient(cfg *config.DiscogsConfig) *DiscogsClient {
	return &DiscogsClient{
		token:     cfg.Token,
		baseUrl:   cfg.BaseURL,
		userAgent: cfg.UserAgent,
		client:    &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *DiscogsClient) doRequest(path string, out interface{}) error {
	resp, err := c.get(path)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error while discogs-request: status %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("error decoding discogs response: %v", err)
	}

	return nil
}

func (c *DiscogsClient) get(path string) (*http.Response, error) {
	separator := "?"
	if strings.Contains(path, "?") {
		separator = "&"
	}

	req, err := http.NewRequest(http.MethodGet, c.baseUrl+path+separator+"token="+url.QueryEscape(c.token), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating GET request: %v", err)
	}

	req.Header.Set("Accept", "application/json")
	// Discogs verlangt einen aussagekräftigen User-Agent, sonst wird die Anfrage abgelehnt
	req.Header.Set("User-Agent", c.userAgent)

	return c.client.Do(req)
}

func (c *DiscogsClient) SearchReleases(query string) ([]discogsSearchResultStruct, error) {
	var resp discogsSearchResponse
	path := fmt.Sprintf("/database/search?q=%s&type=release", url.QueryEscape(query))
	if err := c.doRequest(path, &resp); err != nil {
		return nil, err
	}

	return resp.Results, nil
}

func (c *DiscogsClient) SearchMasters(query string) ([]discogsSearchResultStruct, error) {
	var resp discogsSearchResponse
	path := fmt.Sprintf("/database/search?q=%s&type=master", url.QueryEscape(query))
	if err := c.doRequest(path, &resp); err != nil {
		return nil, err
	}

	return resp.Results, nil
}

func (c *DiscogsClient) SearchArtists(query string) ([]discogsSearchResultStruct, error) {
	var resp discogsSearchResponse
	path := fmt.Sprintf("/database/search?q=%s&type=artist", url.QueryEscape(query))
	if err := c.doRequest(path, &resp); err != nil {
		return nil, err
	}

	return resp.Results, nil
}
