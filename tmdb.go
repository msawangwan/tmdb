package tmdb

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// BadRequestError is returned when an error is returned from the tmDb API.
type BadRequestError struct {
	Resource   string `json:"resource,omitempty"`
	StatusCode int    `json:"statusCode,omitempty"`
}

func newBadRequestError(Resource string, statusCode int) *BadRequestError {
	return &BadRequestError{Resource: Resource, StatusCode: statusCode}
}

// Is implements the Error interface.
func (e *BadRequestError) Is(target error) bool {
	return e.Error() == target.Error()
}

// As implements the Error interface.
func (e *BadRequestError) As(target error) bool {
	return e.Error() == target.Error()
}

// Error implements the Error interface.
func (e *BadRequestError) Error() string {
	return fmt.Sprintf("tmdb: bad request: %d", e.StatusCode)
}

// Exported errors.
var (
	ErrBadRequest = newBadRequestError("", 404)
)

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

// APIClient wraps the standard library http.Client and maintains any state required
// to interact with the tmDb API.
type APIClient struct {
	*APIClientConfig `json:"conf"`
	http.Client      `json:"-"`
	config           *normalizedConfig `json:"-"`
}

// New creates a new client, initialized with customized standard library http.Client.
func New(config io.Reader, timeoutSeconds int) (*APIClient, error) {
	raw, err := ioutil.ReadAll(config)
	if err != nil {
		return nil, err
	}

	var (
		ac *APIClientConfig = &APIClientConfig{}
	)

	if err := json.Unmarshal(raw, ac); err != nil {
		return nil, err
	}

	return &APIClient{
		APIClientConfig: ac,
		Client: http.Client{
			Timeout: time.Second * time.Duration(timeoutSeconds),
		},
		config: &normalizedConfig{
			key:     ac.API.Key,
			baseurl: strings.Trim(ac.API.Baseurl, "/") + "/" + strings.Trim(ac.API.Version, "/"),
			version: strings.Trim(ac.API.Version, "/"),
		},
	}, nil
}

// BuildURL formats the parameters for a given resource request, returning a url.
func (client *APIClient) BuildURL(params EndpointParameters) string {
	p := strings.TrimLeft(params.String(), "/")

	if client.config.version == "3" {
		q := fmt.Sprintf("api_key=%s", client.config.key)

		parts := strings.Split(p, "?")

		if len(parts) > 1 {
			q = fmt.Sprintf("%s&%s", q, parts[1])
		}

		p = parts[0]

		return fmt.Sprintf("%s/%s?%s", client.config.baseurl, p, q) // return fmt.Sprintf("%s/%s?api_key=%s", client.config.baseurl, p, client.config.key)
	}

	return fmt.Sprintf("%s/%s", client.config.baseurl, p)
}

// GET prepares and submites
func (client *APIClient) GET(params EndpointParameters) ([]byte, error) {
	var endpoint = client.BuildURL(params)

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	if client.config.version == "4" {
		req.Header.Add("Authorization", client.config.bearerToken())
	}

	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, newBadRequestError(params.String(), res.StatusCode)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	res.Body.Close()

	return data, nil
}
