package cinema

import (
	"errors"
)

var ErrIncorrectRating = errors.New("incorrect rating")
var ErrMovieAlreadyExists = errors.New("movie already exist")
var ErrMovieNotFound = errors.New("movie not found")