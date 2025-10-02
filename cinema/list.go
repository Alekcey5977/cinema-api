package cinema

import (
	"sync"
)


type List struct {
	Movies map[string]movie
	mtx sync.RWMutex
}

func NewList() *List {
	return &List{
		Movies: make(map[string]movie),
	}
}

func (l *List) AddMovie(movie movie) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if _, ok := l.Movies[movie.Title]; ok {
		return ErrMovieAlreadyExists
	}

	l.Movies[movie.Title] = movie

	return nil
}

func (l *List) GetMovie(title string) (movie, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	getMovie, ok := l.Movies[title]
	if !ok {
		return movie{}, ErrMovieNotFound
	}

	return getMovie, nil
}

func (l *List) ListMovies() map[string]movie {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	movies := make(map[string]movie, len(l.Movies))

	for k,  v := range l.Movies {
		movies[k] = v
	}

	return movies
}

func (l *List) ChangeRatingMovie(title string, rating float64) (movie, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	getMovie, ok := l.Movies[title]
	if ok {
		return movie{}, ErrMovieNotFound
	}

	getMovie.ChangeRating(rating)

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