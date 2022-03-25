package tmdb

import "fmt"

// normalizedConfig is used internally.
type normalizedConfig struct {
	baseurl string `json:"-"`
	key     string `json:"-"`
	version string `json:"-"`
}

func (n *normalizedConfig) bearerToken() string {
	return fmt.Sprintf("Bearer %s", n.key)
}

// APIClientConfig exposes fields that are mapped to a JSON configuration file.
type APIClientConfig struct {
	API struct {
		Baseurl string `json:"baseurl"`
		Version string `json:"version"`
		Key     string `json:"key"`
	} `json:"api"`

	Account struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"account"`
}
