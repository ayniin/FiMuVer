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

type TMDBClient struct {
	apiKey       string
	baseUrl      string
	imageBaseUrl string
	client       *http.Client
}

type tmdbSearchResultStruct struct {
	TmdbId   int    `json:"id"`
	Name     string `json:"name"`
	Overview string `json:"overview"`
	Year     int    `json:"year"`
	ImageURL string `json:"image_url"`
	Type     string `json:"type"` //series / movie
}

// UnmarshalJSON überschreibt das Standard-Unmarshalling, weil TMDB Titel/Name
// und Erscheinungsdatum je nach Endpoint unterschiedlich benennt
// (title/release_date bei Filmen, name/first_air_date bei Serien) und das
// Jahr aus dem Datum extrahiert werden muss statt es als eigenes Feld zu liefern.
func (r *tmdbSearchResultStruct) UnmarshalJSON(data []byte) error {
	var aux struct {
		TmdbId       int    `json:"id"`
		Title        string `json:"title"`
		Name         string `json:"name"`
		Overview     string `json:"overview"`
		ReleaseDate  string `json:"release_date"`
		FirstAirDate string `json:"first_air_date"`
		PosterPath   string `json:"poster_path"`
		MediaType    string `json:"media_type"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	r.TmdbId = aux.TmdbId
	r.Overview = aux.Overview
	r.ImageURL = aux.PosterPath

	if aux.Title != "" {
		r.Name = aux.Title
	} else {
		r.Name = aux.Name
	}

	date := aux.ReleaseDate
	if date == "" {
		date = aux.FirstAirDate
	}
	if len(date) >= 4 {
		fmt.Sscanf(date[:4], "%d", &r.Year)
	}

	r.Type = aux.MediaType

	return nil
}

type tmdbSearchResponse struct {
	Results []tmdbSearchResultStruct `json:"results"`
}

func NewTMDBClient(cfg *config.TMBDConfig) *TMDBClient {
	return &TMDBClient{
		apiKey:       cfg.TmdbApiKey,
		baseUrl:      cfg.BaseURL,
		imageBaseUrl: cfg.ImageBaseURL,
		client:       &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *TMDBClient) doRequest(path string, out interface{}) error {
	resp, err := c.get(path)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error while tmdb-request: status %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("error decoding tmdb response: %v", err)
	}

	return nil
}

func (c *TMDBClient) get(path string) (*http.Response, error) {
	separator := "?"
	if strings.Contains(path, "?") {
		separator = "&"
	}

	req, err := http.NewRequest(http.MethodGet, c.baseUrl+path+separator+"api_key="+url.QueryEscape(c.apiKey), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating GET request: %v", err)
	}

	req.Header.Set("Accept", "application/json")

	return c.client.Do(req)
}

func (c *TMDBClient) resolveImageURLs(results []tmdbSearchResultStruct) []tmdbSearchResultStruct {
	for i := range results {
		if results[i].ImageURL != "" {
			results[i].ImageURL = c.imageBaseUrl + results[i].ImageURL
		}
	}
	return results
}

func (c *TMDBClient) SearchMovies(query string) ([]tmdbSearchResultStruct, error) {
	var resp tmdbSearchResponse
	path := fmt.Sprintf("/search/movie?query=%s", url.QueryEscape(query))
	if err := c.doRequest(path, &resp); err != nil {
		return nil, err
	}

	for i := range resp.Results {
		resp.Results[i].Type = "movie"
	}

	return c.resolveImageURLs(resp.Results), nil
}

func (c *TMDBClient) SearchSeries(query string) ([]tmdbSearchResultStruct, error) {
	var resp tmdbSearchResponse
	path := fmt.Sprintf("/search/tv?query=%s", url.QueryEscape(query))
	if err := c.doRequest(path, &resp); err != nil {
		return nil, err
	}

	for i := range resp.Results {
		resp.Results[i].Type = "series"
	}

	return c.resolveImageURLs(resp.Results), nil
}
