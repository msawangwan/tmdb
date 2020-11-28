package tmdb

import "fmt"

// EndpointParameters defines an interface for a request object and its parameters.
type EndpointParameters interface {
	fmt.Stringer
}

// GetMovieDetails contains the parameters to an API request.
type GetMovieDetails struct {
	ResourceURI string `json:"resourceUri,omitempty"`
	MovieID     string `json:"movieId,omitempty"`
}

// NewGetMovieDetails returns the details of a given movie.
func NewGetMovieDetails(movieID string) GetMovieDetails {
	return GetMovieDetails{
		ResourceURI: "/movie/%s",
		MovieID:     movieID,
	}
}

func (m GetMovieDetails) String() string {
	return fmt.Sprintf(m.ResourceURI, m.MovieID)
}

// GetMovieKeywords returns a list of keywords used to describe a given movie.
type GetMovieKeywords struct {
	ResourceURI string `json:"resourceUri,omitempty"`
	MovieID     string `json:"movieId,omitempty"`
}

// NewGetMovieKeywords returns a request parameters object.
func NewGetMovieKeywords(movieID string) GetMovieKeywords {
	return GetMovieKeywords{
		ResourceURI: "/movie/%s/keywords",
		MovieID:     movieID,
	}
}

func (m GetMovieKeywords) String() string {
	return fmt.Sprintf(m.ResourceURI, m.MovieID)
}

// GetMovieWatchProviders returns a list of streaming services where a given movie
// can be watched.
type GetMovieWatchProviders struct {
	ResourceURI string `json:"resourceUri,omitempty"`
	MovieID     string `json:"movieId,omitempty"`
}

// NewGetMovieWatchProviders returns a request parameters object.
func NewGetMovieWatchProviders(movieID string) GetMovieWatchProviders {
	return GetMovieWatchProviders{
		ResourceURI: "/movie/%s/watch/providers",
		MovieID:     movieID,
	}
}

func (m GetMovieWatchProviders) String() string {
	return fmt.Sprintf(m.ResourceURI, m.MovieID)
}

// GetMovieCredits contains the parameters to an API request.
type GetMovieCredits struct {
	ResourceURI string `json:"resourceUri,omitempty"`
	MovieID     string `json:"movieId,omitempty"`
}

// NewGetMovieCredits returns details regarding the cast and crew.
func NewGetMovieCredits(movieID string) GetMovieCredits {
	return GetMovieCredits{
		ResourceURI: "/movie/%s/credits",
		MovieID:     movieID,
	}
}

func (m GetMovieCredits) String() string {
	return fmt.Sprintf(m.ResourceURI, m.MovieID)
}
