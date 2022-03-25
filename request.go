package tmdb

import "fmt"

// Exported errors.
var (
	ErrBadRequest = newBadRequestError("", 404)
)

// BadRequestError is returned when an error is returned from the tmDb API.
type BadRequestError struct {
	Resource   string `json:"resource"`
	StatusCode int    `json:"statusCode"`
}

func newBadRequestError(Resource string, statusCode int) *BadRequestError {
	return &BadRequestError{Resource: Resource, StatusCode: statusCode}
}

// Is implements the Error interface.
func (e *BadRequestError) Is(target error) bool {
	return e.Error() == target.Error()
}

// As implements the Error interface.
func (e *BadRequestError) As(target any) bool {
	if err, ok := target.(error); ok {
		return e.Is(err)
	}
	return false
}

// Error implements the Error interface.
func (e *BadRequestError) Error() string {
	return fmt.Sprintf("tmdb: bad request: %d", e.StatusCode)
}
