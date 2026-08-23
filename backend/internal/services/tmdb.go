package services

import (
	"net/http"
)

type TMDBClient struct {
	apiKey       string
	baseUrl      string
	imageBaseUrl string
	client       *http.Client
}

type tmdbSearchResultStruct struct {+
	

}