package cinema

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

var ErrIncorrectRating = errors.New("incorrect rating")
var ErrMovieAlreadyExists = errors.New("movie already exist")
var ErrMovieNotFound = errors.New("movie not found")


type ErrorDTO struct {
	Message string
	Time time.Time
}

func NewErrorDTO(message string, time time.Time) *ErrorDTO {
	return &ErrorDTO{
		Message: message,
		Time: time,
	}
}

func (e ErrorDTO) ToString() (string, error) {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal ErrorDTO: %w", err)
	}

	return string(b), nil
}