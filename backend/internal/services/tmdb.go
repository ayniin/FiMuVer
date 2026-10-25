package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"fimuver/internal/config"
)

const tmdbImageBaseURL = "https://image.tmdb.org/t/p/w342"

type TMDBClient struct {
	token   string
	baseURL string
	client  *http.Client
}

// TMDBSearchResult hat dieselbe Form wie die TVDB-Ergebnisse, damit das Frontend beide gleich behandeln kann.
// TmdbID dient nur zum Nachladen der IMDb-ID (ImdbID), gespeichert wird am Item die IMDb-ID.
type TMDBSearchResult struct {
	TmdbID   string `json:"tmdb_id"`
	Name     string `json:"name"`
	Overview string `json:"overview"`
	Year     int    `json:"year"`
	ImageURL string `json:"image_url"`
}

type tmdbSearchResponse struct {
	Results []struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Overview    string `json:"overview"`
		ReleaseDate string `json:"release_date"` // "1999-03-30" oder leer
		PosterPath  string `json:"poster_path"`  // "/abc.jpg" oder leer
	} `json:"results"`
}

func NewTMDBClient(cfg *config.TMDBConfig) *TMDBClient {
	return &TMDBClient{
		token:   cfg.Token,
		baseURL: cfg.BaseURL,
		client:  &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *TMDBClient) get(path string, out any) error {
	req, err := http.NewRequest(http.MethodGet, c.baseURL+path, nil)
	if err != nil {
		return fmt.Errorf("error creating tmdb request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("error while tmdb request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("tmdb request failed with status: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("error decoding tmdb response: %v", err)
	}
	return nil
}

// ImdbID holt die IMDb-ID zu einer TMDB-Film-ID. Leerer String, wenn TMDB keine kennt.
func (c *TMDBClient) ImdbID(tmdbID string) (string, error) {
	var raw struct {
		ImdbID string `json:"imdb_id"`
	}
	if err := c.get("/movie/"+url.PathEscape(tmdbID)+"/external_ids", &raw); err != nil {
		return "", err
	}
	return raw.ImdbID, nil
}

func (c *TMDBClient) SearchMovies(query string) ([]TMDBSearchResult, error) {
	var raw tmdbSearchResponse
	if err := c.get(fmt.Sprintf("/search/movie?query=%s&language=de-DE", url.QueryEscape(query)), &raw); err != nil {
		return nil, err
	}

	results := make([]TMDBSearchResult, 0, len(raw.Results))
	for _, r := range raw.Results {
		res := TMDBSearchResult{
			TmdbID:   strconv.Itoa(r.ID),
			Name:     r.Title,
			Overview: r.Overview,
		}
		if len(r.ReleaseDate) >= 4 {
			res.Year, _ = strconv.Atoi(r.ReleaseDate[:4])
		}
		if r.PosterPath != "" {
			res.ImageURL = tmdbImageBaseURL + r.PosterPath
		}
		results = append(results, res)
	}
	return results, nil
}
