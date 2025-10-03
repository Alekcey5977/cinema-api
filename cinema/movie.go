package cinema

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

type Movie struct {
	Title          string
	Description    string
	Genres         []string
	YearsOfRelease int
	Rating         float64 //  0 <= Rating <= 10
	Country        string
	Adult          bool
}

// Создать новый фильм
func NewMovie(title string,
	description string,
	genres []string,
	yearsOfRelease int,
	rating float64,
	country string,
	adult bool) *Movie {
	return &Movie{
		Title:          title,
		Description:    description,
		Genres:         genres,
		YearsOfRelease: yearsOfRelease,
		Rating:         rating,
		Country:        country,
		Adult:          adult,
	}
}

func (m Movie) ValidateMovieData() error {
	var errorLits []string

	// title validation
	if strings.TrimSpace(m.Title) == "" {
		errorLits = append(errorLits, "title is empty")
	}
	if utf8.RuneCountInString(m.Title) > 100 {
		errorLits = append(errorLits, "title of the movie cannot exceed 1000 characters")
	}

	// description validation
	if strings.TrimSpace(m.Description) == "" {
		errorLits = append(errorLits, "description is empty")
	}
	if utf8.RuneCountInString(m.Description) > 1000 {
		errorLits = append(errorLits, "description of the movie cannot exceed 1000 characters")
	}

	// genre validation
	if len(m.Genres) == 0 {
		errorLits = append(errorLits, "movie must have at least one genre")
	}
	if len(m.Genres) > 10 {
		errorLits = append(errorLits, "movie cannot have more than 10 genres")
	}
	for i, genre := range m.Genres {
		if strings.TrimSpace(genre) == "" {
			errorLits = append(errorLits, "genre is empty")
		}
		if utf8.RuneCountInString(genre) > 50 {
			errorLits = append(errorLits, "genre of the movie cannot exceed 1000 characters")
		}

		for j := i + 1; j < len(m.Genres); j++ {
			if strings.EqualFold(genre, m.Genres[j]) {
				errorLits = append(errorLits, "duplicate genres found:"+genre)
			}
		}
	}

	// yearsOfRelease validation
	currentYear := time.Now().Year()
	if m.YearsOfRelease < 1888 {
		errorLits = append(errorLits, "year pf production cannot be earlier than 1888")
	}
	if m.YearsOfRelease > currentYear+1 {
		errorLits = append(errorLits, "release data may not exceed the curent year by more than 1 year")
	}

	// rating validation
	if m.Rating < 0 || m.Rating > 10 {
		errorLits = append(errorLits, "rating should bi in the range from 0 to 10")
	}

	// country validation
	if strings.TrimSpace(m.Country) == "" {
		errorLits = append(errorLits, "country is not empty")
	}
	if utf8.RuneCountInString(m.Country) > 100 {
		errorLits = append(errorLits, "country of the movie cannot exceed 1000 characters")
	}

	// resalt validation
	if len(errorLits) > 0 {
		return errors.New("error validate the movie: " + strings.Join(errorLits, ";"))
	}

	return nil
}

func (m *Movie) ChangeRating(newRating float64) error {
	if newRating < 0 || newRating > 10 {
		return ErrIncorrectRating
	}

	m.Rating = newRating
	return nil
}
