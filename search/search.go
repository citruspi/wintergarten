package search

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/citruspi/wintergarten/configuration"
	"github.com/citruspi/wintergarten/films"
)

type TMDbResponse struct {
	Results      []Film `json:"results"`
	Page         int    `json:"page"`
	TotalPages   int    `json:"total_pages"`
	TotalResults int    `json:"total_results"`
}
type Film struct {
	Adult         bool    `json:"adult"`
	BackdropPath  string  `json:"backdrop_path"`
	ID            int     `json:"id"`
	OriginalTitle string  `json:"original_title"`
	Popularity    float64 `json:"popularity"`
	PosterPath    string  `json:"poster_path"`
	ReleaseDate   string  `json:"release_date"`
	Title         string  `json:"title"`
	Video         bool    `json:"video"`
	VoteAverage   float64 `json:"vote_average"`
	VoteCount     int     `json:"vote_count"`
}

var (
	conf configuration.Configuration
)

func init() {
	conf = configuration.Init()
}

func Films(query string) ([]Film, error) {
	var collection []Film

	response, err := queryTMDb(query)

	if err != nil {
		return collection, err
	}

	collection = response.Results

	for _, film := range collection {
		filmID := strconv.Itoa(film.ID)

		go films.Prepare(filmID)
	}

	return collection, nil
}

func queryTMDb(query string) (TMDbResponse, error) {
	var URL *url.URL
	parameters := url.Values{}
	var response TMDbResponse

	URL, err := url.Parse("http://api.themoviedb.org/3/search/movie")

	if err != nil {
		return response, err
	}

	parameters.Add("api_key", conf.TMDb.APIKEY)
	parameters.Add("query", query)

	URL.RawQuery = parameters.Encode()

	resp, err := http.Get(URL.String())

	if err != nil {
		return response, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)

	if err != nil {
		return response, err
	}

	return response, nil
}
