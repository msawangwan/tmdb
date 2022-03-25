package tmdb_test

import (
	"errors"
	"os"

	"fmt"
	"testing"

	"encoding/json"
	"strings"

	"github.com/msawangwan/tmdb"
)

const VERBOSE = true

func pretty(t *testing.T, o interface{}) {
	raw, err := json.MarshalIndent(o, "", " ")
	if err != nil {
		t.Fatal(err)
	}

	if VERBOSE {
		t.Logf(string(raw))
	}
}

func createClient() (tmdb.APIClientContract, error) {
	var (
		secret string
	)

	if v, found := os.LookupEnv("TMDB_API_KEY"); found {
		secret = v
	} else {
		return nil, nil
	}

	conf := strings.NewReader(fmt.Sprintf(`{
                "api": {
                        "key": "%s",
                        "version": "/3",
                        "baseurl": "https://api.themoviedb.org"
                },
                "account": {
                        "username": "foo",
                        "password": "bar"
                }
	}`, secret))

	return tmdb.New(conf, 10)
}

func TestCreateClient(t *testing.T) {
	client, err := createClient()
	if err != nil {
		t.Fatal(err)
	}

	pretty(t, client)
}

func TestEndpointv3GetMovieDetails(t *testing.T) {
	client, err := createClient()
	if err != nil {
		t.Fatal(err)
	}

	if client == nil {
		t.Skip("nil client: see README.md for help on configuring the client for test ..")
	}

	var testcases = []struct {
		label string
		req   tmdb.EndpointParameters
	}{
		{"hook", tmdb.NewGetMovieDetails("tt0102057")},
		{"gone_with_the_wind", tmdb.NewGetMovieDetails("tt0031381")},
		{"cheeseballs", tmdb.NewGetMovieDetails("tt3097934")},
	}

	for _, testcase := range testcases {
		t.Run(testcase.label, func(tt *testing.T) {
			if VERBOSE {
				tt.Logf("%s", testcase.req)
				tt.Logf("%s", client.BuildURL(testcase.req))
			}

			if client == nil {
				tt.Skip("no valid api key found, set one with 'TMDB_API_KEY'")
			}

			res, err := client.GET(testcase.req)
			if err != nil {
				if errors.Is(err, tmdb.ErrBadRequest) {
					return
				}
				t.Error(err)
			}

			o := map[string]interface{}{}

			if err := json.Unmarshal(res, &o); err != nil {
				tt.Error(err)
			}

			if VERBOSE {
				pretty(tt, o)
			}
		})
	}
}

func TestEndpointv3GetMovieKeywords(t *testing.T) {
	client, err := createClient()
	if err != nil {
		t.Fatal(err)
	}

	if client == nil {
		t.Skip("nil client: see README.md for help on configuring the client for test ..")
	}

	var testcases = []struct {
		label string
		req   tmdb.EndpointParameters
	}{
		{"hook", tmdb.NewGetMovieKeywords("tt0102057")},
		{"gone_with_the_wind", tmdb.NewGetMovieKeywords("tt0031381")},
		{"cheeseballs", tmdb.NewGetMovieKeywords("tt3097934")},
	}

	for _, testcase := range testcases {
		t.Run(testcase.label, func(tt *testing.T) {
			if VERBOSE {
				tt.Logf("%s", testcase.req)
				tt.Logf("%s", client.BuildURL(testcase.req))
			}

			if client == nil {
				tt.Skip("no valid api key found, set one with 'TMDB_API_KEY'")
			}

			res, err := client.GET(testcase.req)
			if err != nil {
				if errors.Is(err, tmdb.ErrBadRequest) {
					return
				}
				t.Error(err)
			}

			o := map[string]interface{}{}

			if err := json.Unmarshal(res, &o); err != nil {
				tt.Error(err)
			}

			if VERBOSE {
				pretty(tt, o)
			}
		})
	}
}

func TestEndpointv3GetMovieWatchProviders(t *testing.T) {
	client, err := createClient()
	if err != nil {
		t.Fatal(err)
	}

	if client == nil {
		t.Skip("nil client: see README.md for help on configuring the client for test ..")
	}

	var testcases = []struct {
		label string
		req   tmdb.EndpointParameters
	}{
		{"hook", tmdb.NewGetMovieWatchProviders("tt0102057")},
		{"gone_with_the_wind", tmdb.NewGetMovieWatchProviders("tt0031381")},
		{"cheeseballs", tmdb.NewGetMovieWatchProviders("tt3097934")},
	}

	for _, testcase := range testcases {
		t.Run(testcase.label, func(tt *testing.T) {
			if VERBOSE {
				tt.Logf("%s", testcase.req)
				tt.Logf("%s", client.BuildURL(testcase.req))
			}

			if client == nil {
				tt.Skip("no valid api key found, set one with 'TMDB_API_KEY'")
			}

			res, err := client.GET(testcase.req)
			if err != nil {
				if errors.Is(err, tmdb.ErrBadRequest) {
					return
				}
				t.Error(err)
			}

			o := map[string]interface{}{}

			if err := json.Unmarshal(res, &o); err != nil {
				tt.Error(err)
			}

			if VERBOSE {
				pretty(tt, o)
			}
		})
	}
}

func TestEndpointV3GetMovieCredits(t *testing.T) {
	client, err := createClient()
	if err != nil {
		t.Fatal(err)
	}

	if client == nil {
		t.Skip("nil client: see README.md for help on configuring the client for test ..")
	}

	var testcases = []struct {
		label string
		req   tmdb.EndpointParameters
	}{
		{"hook", tmdb.NewGetMovieCredits("tt0102057")},
		{"gone_with_the_wind", tmdb.NewGetMovieCredits("tt0031381")},
		{"cheeseballs", tmdb.NewGetMovieCredits("tt3097934")},
	}

	for _, testcase := range testcases {
		t.Run(testcase.label, func(tt *testing.T) {
			if VERBOSE {
				tt.Logf("%s", testcase.req)
				tt.Logf("%s", client.BuildURL(testcase.req))
			}

			if client == nil {
				tt.Skip("no valid api key found, set one with 'TMDB_API_KEY'")
			}

			res, err := client.GET(testcase.req)
			if err != nil {
				if errors.Is(err, tmdb.ErrBadRequest) {
					return
				}
				t.Error(err)
			}

			o := map[string]interface{}{}

			if err := json.Unmarshal(res, &o); err != nil {
				tt.Error(err)
			}

			if VERBOSE {
				pretty(tt, o)
			}
		})
	}
}

type getTMDbMovieSearchResults struct {
	page         int    // 1
	query        string // "star+wars"
	language     string // "en-US"
	includeAdult string // "false"
}

func newGetTMDbMovieSearchResults(page int, q string) getTMDbMovieSearchResults {
	return getTMDbMovieSearchResults{
		page:         page,
		query:        q,
		language:     "en-US",
		includeAdult: "false",
	}
}

func (m getTMDbMovieSearchResults) String() string {
	return fmt.Sprintf(
		"/search/movie?query=%s&language=%s&include_adult=%s&page=%d",
		m.query,
		m.language,
		m.includeAdult,
		m.page,
	)
}

func TestEndpointV3WithQueryParams(t *testing.T) {
	client, err := createClient()
	if err != nil {
		t.Fatal(err)
	}

	if client == nil {
		t.Skip("nil client: see README.md for help on configuring the client for test ..")
	}

	var testcases = []struct {
		label string
		req   tmdb.EndpointParameters
	}{
		{"first_page", newGetTMDbMovieSearchResults(1, "star+wars")},
		{"second_page", newGetTMDbMovieSearchResults(2, "star+wars")},
	}

	for _, testcase := range testcases {
		t.Run(testcase.label, func(tt *testing.T) {
			if VERBOSE {
				tt.Logf("%s", testcase.req)
				tt.Logf("%s", client.BuildURL(testcase.req))
			}

			if client == nil {
				tt.Skip("no valid api key found, set one with 'TMDB_API_KEY'")
			}

			res, err := client.GET(testcase.req)
			if err != nil {
				if errors.Is(err, tmdb.ErrBadRequest) {
					return
				}
				t.Error(err)
			}

			o := map[string]interface{}{}

			if err := json.Unmarshal(res, &o); err != nil {
				tt.Error(err)
			}

			if VERBOSE {
				pretty(tt, o)
			}
		})
	}
}
