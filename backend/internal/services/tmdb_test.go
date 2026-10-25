package services

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"fimuver/internal/config"
)

func TestTMDBSearchMoviesMapping(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer tok" {
			t.Errorf("missing bearer token, got %q", r.Header.Get("Authorization"))
		}
		if r.URL.Path == "/movie/603/external_ids" {
			w.Write([]byte(`{"id":603,"imdb_id":"tt0133093"}`))
			return
		}
		if r.URL.Query().Get("query") != "matrix" {
			t.Errorf("unexpected query %q", r.URL.Query().Get("query"))
		}
		w.Write([]byte(`{"results":[
			{"id":603,"title":"Matrix","overview":"Neo...","release_date":"1999-03-30","poster_path":"/m.jpg"},
			{"id":1,"title":"Ohne Datum","release_date":"","poster_path":null}
		]}`))
	}))
	defer srv.Close()

	client := NewTMDBClient(&config.TMDBConfig{Token: "tok", BaseURL: srv.URL})
	got, err := client.SearchMovies("matrix")
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Fatalf("want 2 results, got %d", len(got))
	}
	if got[0].TmdbID != "603" || got[0].Year != 1999 || got[0].ImageURL != tmdbImageBaseURL+"/m.jpg" {
		t.Errorf("bad mapping: %+v", got[0])
	}
	if got[1].Year != 0 || got[1].ImageURL != "" {
		t.Errorf("empty date/poster should stay empty: %+v", got[1])
	}

	imdb, err := client.ImdbID(got[0].TmdbID)
	if err != nil || imdb != "tt0133093" {
		t.Errorf("want tt0133093, got %q (err %v)", imdb, err)
	}
}
