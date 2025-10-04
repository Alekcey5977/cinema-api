package cinema

import (
	"sync"
)

type List struct {
	Movies map[string]Movie
	mtx    sync.RWMutex
}

func NewList() *List {
	return &List{
		Movies: make(map[string]Movie),
	}
}

func (l *List) AddMovie(movie Movie) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if _, ok := l.Movies[movie.Title]; ok {
		return ErrMovieAlreadyExists
	}

	l.Movies[movie.Title] = movie

	return nil
}

func (l *List) GetAdultMovie() map[string]Movie {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	movies := make(map[string]Movie, len(l.Movies))

	for k, v := range l.Movies {
		if v.Adult {
			movies[k] = v
		}
	}

	return movies

}

func (l *List) GetNotAdultMovie() map[string]Movie {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	movies := make(map[string]Movie, len(l.Movies))

	for k, v := range l.Movies {
		if !v.Adult {
			movies[k] = v
		}
	}

	return movies
}

func (l *List) GetMovie(title string) (Movie, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	getMovie, ok := l.Movies[title]
	if !ok {
		return Movie{}, ErrMovieNotFound
	}

	return getMovie, nil
}

func (l *List) ListMovies() map[string]Movie {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	movies := make(map[string]Movie, len(l.Movies))

	for k, v := range l.Movies {
		movies[k] = v
	}

	return movies
}

func (l *List) ChangeRatingMovie(title string, rating float64) (Movie, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	getMovie, ok := l.Movies[title]
	if ok {
		return Movie{}, ErrMovieNotFound
	}

	if err := getMovie.ChangeRating(rating); err != nil {
		return Movie{}, ErrIncorrectRating
	}

	l.Movies[title] = getMovie

	return l.Movies[title], nil
}

func (l *List) DeleteMovie(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	_, ok := l.Movies[title]
	if ok {
		return ErrMovieNotFound
	}

	delete(l.Movies, title)

	return nil
}
